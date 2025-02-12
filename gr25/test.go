package gr25

import (
	"encoding/binary"
	"fmt"
	"time"
)

type MessageTest struct {
	ID             int
	Data           []byte
	ExpectedValues map[string]interface{}
}

func (m MessageTest) Run(mqttPort int, dbPort int) bool {
	timestamp := time.Now().UnixMicro()
	uploadKey := 103103
	// Create byte array to hold timestamp (8 bytes) + uploadKey (2 bytes) + data
	result := make([]byte, 10+len(m.Data))
	binary.BigEndian.PutUint64(result[0:8], uint64(timestamp))
	binary.BigEndian.PutUint16(result[8:10], uint16(uploadKey))
	copy(result[10:], m.Data)

	SendMqttMessage(mqttPort, fmt.Sprintf("gr25-test/%d", m.ID), result)
	time.Sleep(1 * time.Second)
	return m.Verify(dbPort)
}

// TODO: Make this actually execute db queries properly
func (m MessageTest) Verify(dbPort int) bool {
	for key, value := range m.ExpectedValues {
		fmt.Println(key, value)
	}
	return false
}

// TODO: Make this send mqtt messages in the correct format
func SendMqttMessage(port int, topic string, message []byte) {

}
