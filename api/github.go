package api

import (
	"mpbench/config"
	"mpbench/model"
	"mpbench/runner"
	"mpbench/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubEventHandler(c *gin.Context) {
	ghEventType := c.Request.Header.Get("X-GitHub-Event")
	if ghEventType == "check_suite" {
		var checkSuiteEvent model.GithubCheckSuiteEvent
		if err := c.ShouldBindJSON(&checkSuiteEvent); err != nil {
			utils.SugarLogger.Errorf("Error parsing check suite event: %v", err)
			return
		}
		if checkSuiteEvent.CheckSuite.App.ClientID != config.GithubAppClientID {
			return
		}

		utils.SugarLogger.Infof("Check suite event: %s", checkSuiteEvent.Action)
		if checkSuiteEvent.Action != "requested" && checkSuiteEvent.Action != "rerequested" {
			return
		}

		go func() {
			runner.CreateGR25Runs(checkSuiteEvent.CheckSuite.HeadSha)
		}()

	} else if ghEventType == "check_run" {
		var checkRunEvent model.GithubCheckRunEvent
		if err := c.ShouldBindJSON(&checkRunEvent); err != nil {
			utils.SugarLogger.Errorf("Error parsing check run event: %v", err)
			return
		}
		if checkRunEvent.CheckRun.App.ClientID != config.GithubAppClientID {
			return
		}

		utils.SugarLogger.Infof("Check run event: %s", checkRunEvent.Action)
		if checkRunEvent.Action != "rerequested" {
			return
		}

		go func() {
			runner.CreateGR25Runs(checkRunEvent.CheckRun.HeadSha)
		}()
	}

	c.JSON(http.StatusOK, gin.H{"message": "Github event received"})
}
