package runner

import (
	"context"
	"mpbench/database"
	"mpbench/gr25"
	"mpbench/mqtt"
	"mpbench/utils"
	"strconv"

	"github.com/gaucho-racing/mapache-go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartTest() {
	mqttPort, dbPort, mqttContainer, singleStoreContainer, err := InitializeContainers()
	if err != nil {
		utils.SugarLogger.Error("Failed to initialize containers", err)
		return
	}

	ctx := context.Background()
	defer mqttContainer.Terminate(ctx)
	defer singleStoreContainer.Terminate(ctx)

	mqttClient, err := mqtt.ConnectMQTT("localhost", mqttPort, "mpbench")
	if err != nil {
		utils.SugarLogger.Error("Failed to connect to MQTT", err)
		return
	}

	db, err := database.ConnectDB("root", "password", "localhost", strconv.Itoa(dbPort), "information_schema")
	if err != nil {
		utils.SugarLogger.Error("Failed to connect to database", err)
		return
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS mapache")
	db.Exec("USE mapache")
	db.AutoMigrate(&mapache.Signal{})

	gr25.SendECUStatusOne(mqttClient, db)
}

func InitializeContainers() (int, int, testcontainers.Container, testcontainers.Container, error) {
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
		return 0, 0, nil, nil, err
	}

	mqttPort, err := mqttContainer.MappedPort(ctx, "1883")
	if err != nil {
		utils.SugarLogger.Error("Failed to get MQTT Broker port", err)
	}
	utils.SugarLogger.Info("MQTT Broker started successfully on port ", mqttPort.Port())

	utils.SugarLogger.Info("Starting SingleStore")
	req = testcontainers.ContainerRequest{
		Image:        "ghcr.io/singlestore-labs/singlestoredb-dev:latest",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("Listening on 0.0.0.0:8080"),
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

	singleStorePort, err := singleStoreContainer.MappedPort(ctx, "3306")
	if err != nil {
		utils.SugarLogger.Error("Failed to get SingleStore port", err)
		return 0, 0, nil, nil, err
	}
	utils.SugarLogger.Info("SingleStore started successfully on port ", singleStorePort.Port())

	return mqttPort.Int(), singleStorePort.Int(), mqttContainer, singleStoreContainer, nil
}
