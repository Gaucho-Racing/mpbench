package runner

import (
	"mpbench/gr25"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func StartGR25Tests(mqttClient *mqtt.Client, db *gorm.DB) {
	gr25.RunECUTests(mqttClient, db)
}
