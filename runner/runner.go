package runner

import (
	"context"
	"fmt"
	"mpbench/database"
	"mpbench/gr25"
	"mpbench/mqtt"
	"mpbench/service"
	"mpbench/utils"
	"strconv"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
	"github.com/gaucho-racing/mapache-go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

func RunTestSuite(serviceName string, mqttClient *mq.Client, db *gorm.DB) {
	if serviceName == "gr25" {
		gr25.RunECUTests(mqttClient, db)
	}
}

func StartTest(serviceName string, commit string) {
	utils.SugarLogger.Infof("Initializing tests for %s at commit %s", serviceName, commit)
	// Checkout repo and build docker image
	repoDir, err := service.CheckoutCommit(commit)
	if err != nil {
		utils.SugarLogger.Error("Failed to checkout repo", err)
		return
	}

	image, err := service.BuildDockerImage(commit, repoDir, serviceName)
	if err != nil {
		utils.SugarLogger.Error("Failed to build docker image", err)
		return
	}

	// Start dependent containers
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

	// Start Mapache service container
	mapacheContainer, err := InitializeMapacheContainer(image, mqttPort, dbPort)
	if err != nil {
		utils.SugarLogger.Errorf("Failed to initialize Mapache service %s", image, err)
		return
	}
	defer mapacheContainer.Terminate(ctx)

	time.Sleep(5 * time.Second)

	RunTestSuite(serviceName, mqttClient, db)

	utils.SugarLogger.Infof("Finished running tests for %s at commit %s", serviceName, commit)
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

// StdoutLogConsumer is a LogConsumer that prints the log to stdout
type StdoutLogConsumer struct{}

// Accept prints the log to stdout
func (lc *StdoutLogConsumer) Accept(l testcontainers.Log) {
	fmt.Print(string(l.Content))
}

func InitializeMapacheContainer(image string, mqttPort int, dbPort int) (testcontainers.Container, error) {
	ctx := context.Background()
	utils.SugarLogger.Infof("Starting %s", image)
	req := testcontainers.ContainerRequest{
		Image:        image,
		ExposedPorts: []string{"7000/tcp"},
		WaitingFor:   wait.ForLog("Connected to database"),
		Env: map[string]string{
			"ENV":               "PROD",
			"PORT":              "7000",
			"MQTT_HOST":         "host.docker.internal",
			"MQTT_PORT":         strconv.Itoa(mqttPort),
			"MQTT_USER":         "mpbench",
			"MQTT_PASSWORD":     "mpbench",
			"DATABASE_HOST":     "host.docker.internal",
			"DATABASE_PORT":     strconv.Itoa(dbPort),
			"DATABASE_NAME":     "mapache",
			"DATABASE_USER":     "root",
			"DATABASE_PASSWORD": "password",
		},
		LogConsumerCfg: &testcontainers.LogConsumerConfig{
			Opts: []testcontainers.LogProductionOption{
				testcontainers.WithLogProductionTimeout(10 * time.Second),
			},
			Consumers: []testcontainers.LogConsumer{
				&StdoutLogConsumer{},
			},
		},
	}
	mapacheContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		utils.SugarLogger.Errorf("Failed to start %s", image, err)
		return nil, err
	}

	mapachePort, err := mapacheContainer.MappedPort(ctx, "7000")
	if err != nil {
		utils.SugarLogger.Errorf("Failed to get %s port", image, err)
	}
	utils.SugarLogger.Infof("%s started successfully on port %s", image, mapachePort.Port())

	return mapacheContainer, nil
}
