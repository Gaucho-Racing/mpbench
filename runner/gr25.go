package runner

import (
	"mpbench/gr25"
	"mpbench/model"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func StartGR25Tests(run model.Run, mqttClient *mqtt.Client, db *gorm.DB) {
	gr25.RunECUTests(run, mqttClient, db)
}
