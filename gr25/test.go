package gr25

import (
	"encoding/binary"
	"fmt"
	"mpbench/model"
	"mpbench/mqtt"
	"mpbench/service"
	"mpbench/utils"
	"sync"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
	"github.com/gaucho-racing/mapache-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RunTests(run model.Run, mqttClient *mq.Client, db *gorm.DB) {
	var tests = []MessageTest{}
	tests = append(tests, GenerateECUTests()...)

	wg := sync.WaitGroup{}
	wg.Add(len(tests))

	for _, test := range tests {
		go func(test MessageTest) {
			test.Run(run, mqttClient, db)
			wg.Done()
		}(test)
	}
	wg.Wait()
}

const VehicleID = "gr25-test"
const UploadKey = 10310

type MessageTest struct {
	// ID is the CAN ID of the message
	ID int
	// Node is the node that the message is sent from
	Node string
	// Name is a public-facing name for the test
	Name string
	// Data is the data contained in the message
	Data []byte
	// ExpectedValues is a map of the expected signal ids and their final values
	ExpectedValues map[string]interface{}
}

func (m MessageTest) Run(run model.Run, mqttClient *mq.Client, db *gorm.DB) bool {
	timestamp := time.Now().UnixMicro()
	// Create byte array to hold timestamp (8 bytes) + uploadKey (2 bytes) + data
	result := make([]byte, 10+len(m.Data))
	binary.BigEndian.PutUint64(result[0:8], uint64(timestamp))
	binary.BigEndian.PutUint16(result[8:10], uint16(UploadKey))
	copy(result[10:], m.Data)

	utils.SugarLogger.Infof("STARTING TEST: 0x%03x %s", m.ID, m.Name)
	run_test := model.RunTest{
		ID:     uuid.New().String(),
		RunID:  run.ID,
		Name:   fmt.Sprintf("0x%03x %s", m.ID, m.Name),
		Status: "in_progress",
		Data:   formatData(m.Data),
	}
	service.CreateRunTest(run_test)

	SendMqttMessage(mqttClient, fmt.Sprintf("gr25/%s/%s/%03x", VehicleID, m.Node, m.ID), result)
	WaitForSignals(len(m.ExpectedValues), timestamp, db)
	status := m.Verify(run_test, db, timestamp)
	if status == "passed" {
		utils.SugarLogger.Infof("✅ TEST PASSED: 0x%03x %s", m.ID, m.Name)
		run_test.Status = "passed"
	} else if status == "partial" {
		utils.SugarLogger.Infof("⚠️ TEST PARTIAL: 0x%03x %s", m.ID, m.Name)
		run_test.Status = "partial"
	} else {
		utils.SugarLogger.Infof("❌ TEST FAILED: 0x%03x %s", m.ID, m.Name)
		run_test.Status = "failed"
	}
	service.CreateRunTest(run_test)
	return status == "passed"
}

func WaitForSignals(numSignals int, timestamp int64, db *gorm.DB) {
	timeout := time.After(20 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			utils.SugarLogger.Warnf("Timeout waiting for signals (received %d/%d)", getSignalCount(timestamp, db), numSignals)
			return
		case <-ticker.C:
			count := getSignalCount(timestamp, db)
			if count >= int64(numSignals) {
				utils.SugarLogger.Infof("Received %d/%d signals", count, numSignals)
				return
			}
		}
	}
}

func getSignalCount(timestamp int64, db *gorm.DB) int64 {
	var count int64
	db.Model(&mapache.Signal{}).
		Where("timestamp = ?", timestamp).
		Where("vehicle_id = ?", VehicleID).
		Count(&count)
	return count
}

func (m MessageTest) Verify(run_test model.RunTest, db *gorm.DB, timestamp int64) string {
	failedSignals := []string{}
	for key, value := range m.ExpectedValues {
		valueFloat, ok := value.(float64)
		if !ok {
			valueFloat = float64(value.(int))
		}
		var signal mapache.Signal
		db.Where("timestamp = ?", timestamp).Where("vehicle_id = ?", VehicleID).Where("name = ?", key).First(&signal)
		if signal.Name == "" {
			utils.SugarLogger.Infof("%s: DNE != %v", key, value)
			failedSignals = append(failedSignals, key)
			run_test_result := model.RunTestResult{
				ID:         uuid.New().String(),
				RunTestID:  run_test.ID,
				SignalName: key,
				Status:     "failed",
				Value:      "DNE",
				Expected:   fmt.Sprintf("%v", value),
			}
			service.CreateRunTestResult(run_test_result)
		} else if !almostEqual(signal.Value, valueFloat, 1e-6) {
			utils.SugarLogger.Infof("%s: %f scaled (%d raw) != %v", key, signal.Value, signal.RawValue, value)
			failedSignals = append(failedSignals, key)
			run_test_result := model.RunTestResult{
				ID:         uuid.New().String(),
				RunTestID:  run_test.ID,
				SignalName: key,
				Status:     "failed",
				Value:      fmt.Sprintf("%f scaled (%d raw)", signal.Value, signal.RawValue),
				Expected:   fmt.Sprintf("%v", value),
			}
			service.CreateRunTestResult(run_test_result)
		} else {
			run_test_result := model.RunTestResult{
				ID:         uuid.New().String(),
				RunTestID:  run_test.ID,
				SignalName: key,
				Status:     "passed",
				Value:      fmt.Sprintf("%f scaled (%d raw)", signal.Value, signal.RawValue),
				Expected:   fmt.Sprintf("%v", value),
			}
			service.CreateRunTestResult(run_test_result)
		}
	}
	utils.SugarLogger.Infof("Correctly ingested %d/%d signals", len(m.ExpectedValues)-len(failedSignals), len(m.ExpectedValues))
	if len(failedSignals) == 0 {
		return "passed"
	} else if len(failedSignals) < len(m.ExpectedValues) {
		return "partial"
	} else {
		return "failed"
	}
}

func SendMqttMessage(mqttClient *mq.Client, topic string, message []byte) {
	err := mqtt.PublishMessage(mqttClient, topic, message)
	if err != nil {
		utils.SugarLogger.Error("Failed to publish MQTT message", err)
	} else {
		utils.SugarLogger.Infof("Published MQTT message to %s", topic)
	}
}

func formatData(data []byte) string {
	if len(data) == 0 {
		return "{}"
	}

	result := "{"
	for i, b := range data {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("0x%02x", b)
	}
	result += "}"
	return result
}

func almostEqual(a, b, epsilon float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}
