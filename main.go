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

	runner.StartTest("gr25", "d2eaec342f7e9673119d8bb62bd3c15d9db30269")
}
