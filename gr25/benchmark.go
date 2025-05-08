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

	run_test1 := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (10000 @ 10ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test1)

	run_test2 := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (10000 @ 50ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test2)

	run_test3 := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (10000 @ 100ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test3)

	run_test4 := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (10000 @ 250ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test4)

	run_test5 := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   "Fast (10000 @ 500ms)",
		Status: "in_progress",
	}
	service.CreateRunTest(run_test5)

	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (10000 @ 10ms)")
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
	status := VerifyBenchmarkResults(run_test1, endTime.Sub(startTime), results)
	run_test1.Status = status
	service.CreateRunTest(run_test1)

	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (10000 @ 50ms)")
	startTime = time.Now()
	messageMin = 10000
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
	status = VerifyBenchmarkResults(run_test2, endTime.Sub(startTime), results)
	run_test2.Status = status
	service.CreateRunTest(run_test2)

	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (10000 @ 100ms)")
	startTime = time.Now()
	messageMin = 10000
	numSent = 0
	for numSent < messageMin {
		for _, test := range tests {
			PublishMessageFuzz(mqttClient, test)
			time.Sleep(100 * time.Millisecond)
		}
		numSent += numSignals
	}
	WaitForBenchmark(numSent, db)
	endTime = time.Now()
	utils.SugarLogger.Infof("Benchmark completed in %s", endTime.Sub(startTime))
	results = GenerateBenchmarkResults(run, db)
	status = VerifyBenchmarkResults(run_test3, endTime.Sub(startTime), results)
	run_test3.Status = status
	service.CreateRunTest(run_test3)

	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (10000 @ 250ms)")
	startTime = time.Now()
	messageMin = 10000
	numSent = 0
	for numSent < messageMin {
		for _, test := range tests {
			PublishMessageFuzz(mqttClient, test)
			time.Sleep(250 * time.Millisecond)
		}
		numSent += numSignals
	}
	WaitForBenchmark(numSent, db)
	endTime = time.Now()
	utils.SugarLogger.Infof("Benchmark completed in %s", endTime.Sub(startTime))
	results = GenerateBenchmarkResults(run, db)
	status = VerifyBenchmarkResults(run_test4, endTime.Sub(startTime), results)
	run_test4.Status = status
	service.CreateRunTest(run_test4)

	utils.SugarLogger.Infof("STARTING BENCHMARK: FAST (10000 @ 500ms)")
	startTime = time.Now()
	messageMin = 10000
	numSent = 0
	for numSent < messageMin {
		for _, test := range tests {
			PublishMessageFuzz(mqttClient, test)
			time.Sleep(500 * time.Millisecond)
		}
		numSent += numSignals
	}
	WaitForBenchmark(numSent, db)
	endTime = time.Now()
	utils.SugarLogger.Infof("Benchmark completed in %s", endTime.Sub(startTime))
	results = GenerateBenchmarkResults(run, db)
	status = VerifyBenchmarkResults(run_test5, endTime.Sub(startTime), results)
	run_test5.Status = status
	service.CreateRunTest(run_test5)
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
	timeout := time.After(5 * time.Minute)
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

	// Delete all signals after collecting metrics
	db.Where("vehicle_id = ?", VehicleID).Delete(&mapache.Signal{})

	return map[string]float64{
		"min": min,
		"max": max,
		"avg": avg,
		"p95": p95,
		"p99": p99,
	}
}

func VerifyBenchmarkResults(run_test model.RunTest, totalTime time.Duration, results map[string]float64) string {
	if len(results) == 0 {
		return "failed"
	}

	// Maximum acceptable latency for each benchmark type
	// In order of total time, min, max, avg, p95, p99
	quotas := map[string]int64{
		"min": 10,
		"max": 5000,
		"avg": 600,
		"p95": 1000,
		"p99": 2000,
	}

	rts := []model.RunTestResult{}
	status := ""

	if totalTime.Milliseconds() <= 20*60*1000 {
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
		Expected:   formatResultMillis(5 * 60 * 1000),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["min"]) <= quotas["min"] {
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
		Expected:   formatResultMillis(quotas["min"]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["max"]) <= quotas["max"] {
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
		Expected:   formatResultMillis(quotas["max"]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["avg"]) <= quotas["avg"] {
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
		Expected:   formatResultMillis(quotas["avg"]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["p95"]) <= quotas["p95"] {
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
		Expected:   formatResultMillis(quotas["p95"]),
	}
	service.CreateRunTestResult(run_test_result)
	rts = append(rts, run_test_result)

	if int64(results["p99"]) <= quotas["p99"] {
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
		Expected:   formatResultMillis(quotas["p99"]),
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
