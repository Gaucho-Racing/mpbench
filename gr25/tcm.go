package gr25

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func RunTCMTests(mqttClient *mqtt.Client, db *gorm.DB) {
	SendTCMStatus(mqttClient, db)
	SendTCMResourceUtilization(mqttClient, db)
}

func SendTCMStatus(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x029,
		Name: "TCM Status Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"tcm_connection_status":   0,
			"tcm_mqtt_status":         0,
			"tcm_epic_shelter_status": 0,
			"tcm_camera_status":       0,
			"reserved_1":              0,
			"tcm_ping":                0,
			"tcm_cache_size":          0,
			"reserved_2":              0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendTCMResourceUtilization(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x02A,
		Name: "TCM Resource Utilization Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"tcm_cpu_usage":         0,
			"tcm_gpu_usage":         0,
			"tcm_memory_usage":      0,
			"tcm_storage_usage":     0,
			"tcm_power_consumption": 0,
			"tcm_cpu_temp":          0,
			"tcm_gpu_temp":          0,
		},
	}
	test1.Run(mqttClient, db)
}
