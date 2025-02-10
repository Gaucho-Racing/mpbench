package runner

import (
	"context"
	"mpbench/utils"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartTest() {
	InitializeContainers()
}

func InitializeContainers() {
	ctx := context.Background()
	utils.SugarLogger.Info("Starting MQTT Broker")
	req := testcontainers.ContainerRequest{
		Image:        "emqx/nanomq:latest",
		ExposedPorts: []string{"1883/tcp"},
		WaitingFor:   wait.ForLog("NanoMQ Broker is started successfully!"),
	}
	mqttContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		utils.SugarLogger.Error("Failed to start MQTT Broker", err)
	}
	defer mqttContainer.Terminate(ctx)

	mqttPort, err := mqttContainer.MappedPort(ctx, "1883")
	if err != nil {
		utils.SugarLogger.Error("Failed to get MQTT Broker port", err)
	}
	utils.SugarLogger.Info("MQTT Broker started successfully at ", mqttPort.Port())

	utils.SugarLogger.Info("Starting SingleStore")
	req = testcontainers.ContainerRequest{
		Image:        "ghcr.io/singlestore-labs/singlestoredb-dev:latest",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("NanoMQ Broker is started successfully!"),
		Env: map[string]string{
			"ROOT_PASSWORD": "password",
		},
	}
	singleStoreContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		utils.SugarLogger.Error("Failed to start SingleStore", err)
	}
	defer mqttContainer.Terminate(ctx)

	singleStorePort, err := singleStoreContainer.MappedPort(ctx, "3306")
	if err != nil {
		utils.SugarLogger.Error("Failed to get SingleStore port", err)
	}
	utils.SugarLogger.Info("SingleStore started successfully at ", singleStorePort.Port())
}
