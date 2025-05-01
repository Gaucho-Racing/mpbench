package gr25

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"mpbench/model"
	"mpbench/service"
	"mpbench/utils"
	"sort"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
	"github.com/gaucho-racing/mapache-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RunBenchmark(run model.Run, mqttClient *mq.Client, db *gorm.DB) {
	var tests = []MessageTest{}
	tests = append(tests, GenerateECUTests()...)
	tests = append(tests, GenerateACUTests()...)

	numSignals := 0
	for _, test := range tests {
		numSignals += len(test.ExpectedValues)
	}

	// Fast test
	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (10000 @ 10ms)")
	run_test := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (10000 @ 10ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test)
	startTime := time.Now()
	messageMin := 10000
	numSent := 0
	for numSent < messageMin {
		for _, test := range tests {
			PublishMessageFuzz(mqttClient, test)
			time.Sleep(10 * time.Millisecond)
		}
		numSent += numSignals
	}
	WaitForBenchmark(numSent, db)
	endTime := time.Now()
	utils.SugarLogger.Infof("Benchmark completed in %s", endTime.Sub(startTime))
	results := GenerateBenchmarkResults(run, db)
	status := VerifyBenchmarkResults(run_test, "fast", endTime.Sub(startTime), results)
	run_test.Status = status
	service.CreateRunTest(run_test)

	// Fast test
	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (5000 @ 10ms)")
	run_test = model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (5000 @ 10ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test)
	startTime = time.Now()
	messageMin = 5000
	numSent = 0
	for numSent < messageMin {
		for _, test := range tests {
			PublishMessageFuzz(mqttClient, test)
			time.Sleep(10 * time.Millisecond)
		}
		numSent += numSignals
	}
	WaitForBenchmark(numSent, db)
	endTime = time.Now()
	utils.SugarLogger.Infof("Benchmark completed in %s", endTime.Sub(startTime))
	results = GenerateBenchmarkResults(run, db)
	status = VerifyBenchmarkResults(run_test, "fast", endTime.Sub(startTime), results)
	run_test.Status = status
	service.CreateRunTest(run_test)

	// Fast test
	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (5000 @ 50ms)")
	run_test = model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (5000 @ 50ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test)
	startTime = time.Now()
	messageMin = 5000
	numSent = 0
	for numSent < messageMin {
		for _, test := range tests {
			PublishMessageFuzz(mqttClient, test)
			time.Sleep(50 * time.Millisecond)
		}
		numSent += numSignals
	}
	WaitForBenchmark(numSent, db)
	endTime = time.Now()
	utils.SugarLogger.Infof("Benchmark completed in %s", endTime.Sub(startTime))
	results = GenerateBenchmarkResults(run, db)
	status = VerifyBenchmarkResults(run_test, "fast", endTime.Sub(startTime), results)
	run_test.Status = status
	service.CreateRunTest(run_test)
}

func PublishMessageFuzz(mqttClient *mq.Client, m MessageTest) {
	timestamp := time.Now().UnixMilli()
	// Create byte array to hold timestamp (8 bytes) + uploadKey (2 bytes) + data
	result := make([]byte, 10+len(m.Data))
	binary.BigEndian.PutUint64(result[0:8], uint64(timestamp))
	binary.BigEndian.PutUint16(result[8:10], uint16(UploadKey))
	// Generate random data of the same length as m.Data
	fuzz := make([]byte, len(m.Data))
	for i := range fuzz {
		fuzz[i] = byte(rand.Intn(256))
	}
	copy(result[10:], fuzz)

	SendMqttMessage(mqttClient, fmt.Sprintf("gr25/%s/%s/%03x", VehicleID, m.Node, m.ID), result)
}

func WaitForBenchmark(numSignals int, db *gorm.DB) {
	timeout := time.After(60 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			utils.SugarLogger.Warnf("Timeout waiting for signals (received %d/%d)", getTotalSignalCount(db), numSignals)
			return
		case <-ticker.C:
			count := getTotalSignalCount(db)
			if count >= int64(numSignals) {
				utils.SugarLogger.Infof("Received %d/%d signals", count, numSignals)
				return
			}
		}
	}
}

func getTotalSignalCount(db *gorm.DB) int64 {
	var count int64
	db.Model(&mapache.Signal{}).
		Where("vehicle_id = ?", VehicleID).
		Count(&count)
	return count
}

func GenerateBenchmarkResults(run model.Run, db *gorm.DB) map[string]float64 {
	var signals []mapache.Signal
	db.Where("vehicle_id = ?", VehicleID).Find(&signals)

	var latencies []float64
	for _, signal := range signals {
		latency := signal.CreatedAt.Sub(signal.ProducedAt).Seconds() * 1000
		latencies = append(latencies, latency)
	}

	// Sort latencies for percentile calculations
	sort.Float64s(latencies)

	n := len(latencies)
	if n < 20 {
		utils.SugarLogger.Warnf("Insufficient signals (%d) for meaningful statistics", n)
		return map[string]float64{}
	}

	var sum float64
	for _, l := range latencies {
		sum += l
	}

	min := latencies[0]
	max := latencies[n-1]
	avg := sum / float64(n)
	p95 := latencies[int(float64(n)*0.95)]
	p99 := latencies[int(float64(n)*0.99)]

	return map[string]float64{
		"min": min,
		"max": max,
		"avg": avg,
		"p95": p95,
		"p99": p99,
	}
}

func VerifyBenchmarkResults(run_test model.RunTest, benchmarkType string, totalTime time.Duration, results map[string]float64) string {
	if len(results) == 0 {
		return "failed"
	}

	// Maximum acceptable latency for each benchmark type
	// In order of total time, min, max, avg, p95, p99
	quotas := map[string][]int64{
		"fast": {20000, 100, 5000, 600, 4000, 4500},
	}

	rts := []model.RunTestResult{}
	status := ""

	if totalTime.Milliseconds() <= quotas[benchmarkType][0] {
		status = "passed"
	} else {
		status = "failed"
	}
	run_test_result := model.RunTestResult{
		ID:         uuid.New().String(),
		RunTestID:  run_test.ID,
		SignalName: "total_ingest_time",
		Status:     status,
		Value:      formatResultMillis(totalTime.Milliseconds()),
		Expected:   formatResultMillis(quotas[benchmarkType][0]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["min"]) <= quotas[benchmarkType][1] {
		status = "passed"
	} else {
		status = "failed"
	}
	run_test_result = model.RunTestResult{
		ID:         uuid.New().String(),
		RunTestID:  run_test.ID,
		SignalName: "min_latency",
		Status:     status,
		Value:      formatResultMillis(int64(results["min"])),
		Expected:   formatResultMillis(quotas[benchmarkType][1]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["max"]) <= quotas[benchmarkType][2] {
		status = "passed"
	} else {
		status = "failed"
	}
	run_test_result = model.RunTestResult{
		ID:         uuid.New().String(),
		RunTestID:  run_test.ID,
		SignalName: "max_latency",
		Status:     status,
		Value:      formatResultMillis(int64(results["max"])),
		Expected:   formatResultMillis(quotas[benchmarkType][2]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["avg"]) <= quotas[benchmarkType][3] {
		status = "passed"
	} else {
		status = "failed"
	}
	run_test_result = model.RunTestResult{
		ID:         uuid.New().String(),
		RunTestID:  run_test.ID,
		SignalName: "avg_latency",
		Status:     status,
		Value:      formatResultMillis(int64(results["avg"])),
		Expected:   formatResultMillis(quotas[benchmarkType][3]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["p95"]) <= quotas[benchmarkType][4] {
		status = "passed"
	} else {
		status = "failed"
	}
	run_test_result = model.RunTestResult{
		ID:         uuid.New().String(),
		RunTestID:  run_test.ID,
		SignalName: "p95_latency",
		Status:     status,
		Value:      formatResultMillis(int64(results["p95"])),
		Expected:   formatResultMillis(quotas[benchmarkType][4]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["p99"]) <= quotas[benchmarkType][5] {
		status = "passed"
	} else {
		status = "failed"
	}
	run_test_result = model.RunTestResult{
		ID:         uuid.New().String(),
		RunTestID:  run_test.ID,
		SignalName: "p99_latency",
		Status:     status,
		Value:      formatResultMillis(int64(results["p99"])),
		Expected:   formatResultMillis(quotas[benchmarkType][5]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	numFailed := 0
	for _, rt := range rts {
		if rt.Status == "failed" {
			numFailed++
		}
	}

	if numFailed == 6 {
		utils.SugarLogger.Infof("❌ BENCHMARK FAILED: %s", run_test.Name)
		return "failed"
	} else if numFailed > 0 {
		utils.SugarLogger.Infof("⚠️ BENCHMARK PARTIAL: %s", run_test.Name)
		return "partial"
	} else {
		utils.SugarLogger.Infof("✅ BENCHMARK PASSED: %s", run_test.Name)
		return "passed"
	}
}

func formatResultMillis(millis int64) string {
	if millis < 1000 {
		return fmt.Sprintf("%dms", millis)
	}
	seconds := float64(millis) / 1000.0
	if seconds < 60 {
		return fmt.Sprintf("%dms (%.3fs)", millis, seconds)
	}
	minutes := int(seconds / 60)
	remainingSeconds := seconds - float64(minutes*60)
	return fmt.Sprintf("%dms (%dm %.3fs)", millis, minutes, remainingSeconds)
}
