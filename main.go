package main

import (
	"mpbench/config"
	"mpbench/service"
	"mpbench/utils"
)

func main() {
	config.PrintStartupBanner()
	utils.InitializeLogger()
	defer utils.Logger.Sync()

	repoDir := service.CheckoutCommit("ff15acf247f3ade5036c9f2ed19fbdc70dd2b1ab")
	service.BuildDockerImage("ff15acf247f3ade5036c9f2ed19fbdc70dd2b1ab", repoDir, "gr24")
	// runner.StartTest()
}
