package main

import (
	"mpbench/config"
	"mpbench/runner"
	"mpbench/utils"
)

func main() {
	config.PrintStartupBanner()
	utils.InitializeLogger()
	defer utils.Logger.Sync()

	runner.StartTest()
}
