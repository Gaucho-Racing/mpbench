package main

import (
	"mpbench/api"
	"mpbench/config"
	"mpbench/utils"
)

func main() {
	config.PrintStartupBanner()
	utils.InitializeLogger()
	utils.VerifyConfig()
	defer utils.Logger.Sync()

	// runner.StartTest("gr25", "ec4a64b247b4378a132f434a85aca770d6ce22a1")
	router := api.SetupRouter()
	api.InitializeRoutes(router)
	err := router.Run(":" + config.Port)
	if err != nil {
		utils.SugarLogger.Fatalln(err)
	}
}
