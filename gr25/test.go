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

func (m MessageTest) Run(mqttPort int, dbPort int) {
	timestamp := time.Now().UnixMicro()
	uploadKey := 103103
	// Create byte array to hold timestamp (8 bytes) + uploadKey (2 bytes) + data
	result := make([]byte, 10+len(m.Data))
	binary.BigEndian.PutUint64(result[0:8], uint64(timestamp))
	binary.BigEndian.PutUint16(result[8:10], uint16(uploadKey))
	copy(result[10:], m.Data)



func (m MessageTest) Verify(dbPort int) {
	for key, value := range m.ExpectedValues {
		query := fmt.Sprintf("SELECT %s FROM gr25_messages WHERE id = %d", key, m.ID)
		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}
	}
}

func SendMqttMessage(port int, topic string, message []byte) {
	
}