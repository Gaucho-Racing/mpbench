package gr25

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"mpbench/model"
	"mpbench/utils"
	"sort"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
	"github.com/gaucho-racing/mapache-go"
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
	timeout := time.After(20 * time.Second)
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

func GenerateBenchmarkResults(run model.Run, db *gorm.DB) {
	var signals []mapache.Signal
	db.Where("vehicle_id = ?", VehicleID).Find(&signals)

	var latencies []float64
	for _, signal := range signals {
		latency := signal.CreatedAt.Sub(signal.ProducedAt).Seconds() * 1000
		latencies = append(latencies, latency)
	}

	// Sort latencies for percentile calculations
	sort.Float64s(latencies)

	var sum float64
	for _, l := range latencies {
		sum += l
	}

	n := len(latencies)
	if n == 0 {
		utils.SugarLogger.Warn("No signals found for latency calculation")
		return
	}

	min := latencies[0]
	max := latencies[n-1]
	avg := sum / float64(n)
	p95 := latencies[int(float64(n)*0.95)]
	p99 := latencies[int(float64(n)*0.99)]

	metrics := []struct {
		name  string
		value float64
	}{
		{"Min Latency (ms)", min},
		{"Max Latency (ms)", max},
		{"Avg Latency (ms)", avg},
		{"95th Percentile Latency (ms)", p95},
		{"99th Percentile Latency (ms)", p99},
	}

	fmt.Printf("Latency Metrics: %v\n", metrics)
}