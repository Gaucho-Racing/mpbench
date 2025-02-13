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

	runner.StartTest("gr24", "ff15acf247f3ade5036c9f2ed19fbdc70dd2b1ab")
}
