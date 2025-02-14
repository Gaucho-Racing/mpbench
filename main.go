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

	runner.StartTest("gr25", "a729129f9047374ab26265f436bb096b8fc9fe42")
}
