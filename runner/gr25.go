package runner

import (
	"mpbench/gr25"
	"mpbench/model"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateGR25Run(commit string) (model.Run, error) {
	run := model.Run{
		ID:      uuid.New().String(),
		Commit:  commit,
		Status:  "queued",
		Name:    "mpbench / gr25",
		Service: "gr25",
	}
	return run, nil
}

func StartGR25Tests(run model.Run, mqttClient *mqtt.Client, db *gorm.DB) {
	gr25.RunECUTests(run, mqttClient, db)
}
