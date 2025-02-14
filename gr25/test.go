package gr25

import (
	"encoding/binary"
	"fmt"
	"mpbench/mqtt"
	"mpbench/utils"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
	"github.com/gaucho-racing/mapache-go"
	"gorm.io/gorm"
)

const VehicleID = "gr25-test"
const UploadKey = 10310

type MessageTest struct {
	// ID is the CAN ID of the message
	ID int
	// Name is a public-facing name for the test
	Name string
	// Data is the data contained in the message
	Data []byte
	// ExpectedValues is a map of the expected signal ids and their final values
	ExpectedValues map[string]interface{}
}

func (m MessageTest) Run(mqttClient *mq.Client, db *gorm.DB) bool {
	timestamp := time.Now().UnixMicro()
	// Create byte array to hold timestamp (8 bytes) + uploadKey (2 bytes) + data
	result := make([]byte, 10+len(m.Data))
	binary.BigEndian.PutUint64(result[0:8], uint64(timestamp))
	binary.BigEndian.PutUint16(result[8:10], uint16(UploadKey))
	copy(result[10:], m.Data)

	utils.SugarLogger.Infof("STARTING TEST: 0x%03x %s", m.ID, m.Name)

	SendMqttMessage(mqttClient, fmt.Sprintf("gr25/%s/%03x", VehicleID, m.ID), result)
	time.Sleep(1 * time.Second)
	status := m.Verify(db, timestamp)
	if !status {
		utils.SugarLogger.Infof("❌ TEST FAILED: 0x%03x %s", m.ID, m.Name)
		return false
	}
	utils.SugarLogger.Infof("✅ TEST PASSED: 0x%03x %s", m.ID, m.Name)
	return true
}

func (m MessageTest) Verify(db *gorm.DB, timestamp int64) bool {
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
		} else if signal.Value != valueFloat {
			utils.SugarLogger.Infof("%s: %f scaled (%d raw) != %v", key, signal.Value, signal.RawValue, value)
			failedSignals = append(failedSignals, key)
		}
	}
	utils.SugarLogger.Infof("Correctly ingested %d/%d signals", len(m.ExpectedValues)-len(failedSignals), len(m.ExpectedValues))
	return len(failedSignals) == 0
}

func SendMqttMessage(mqttClient *mq.Client, topic string, message []byte) {
	err := mqtt.PublishMessage(mqttClient, topic, message)
	if err != nil {
		utils.SugarLogger.Error("Failed to publish MQTT message", err)
	} else {
		utils.SugarLogger.Infof("Published MQTT message to %s", topic)
	}
}
