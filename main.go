package main

import (
	"mpbench/api"
	"mpbench/config"
	"mpbench/database"
	"mpbench/runner"
	"mpbench/utils"
)

func main() {
	config.PrintStartupBanner()
	utils.InitializeLogger()
	utils.VerifyConfig()
	defer utils.Logger.Sync()

	database.InitializeDB()
	runner.InitializeScheduler(config.MaxWorkers)
	defer runner.Queue.Stop()

	router := api.SetupRouter()
	api.InitializeRoutes(router)
	err := router.Run(":" + config.Port)
	if err != nil {
		utils.SugarLogger.Fatalln(err)
	}
}
