package api

import (
	"mpbench/config"
	"mpbench/model"
	"mpbench/runner"
	"mpbench/service"
	"mpbench/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GithubEventHandler(c *gin.Context) {
	ghEventType := c.Request.Header.Get("X-GitHub-Event")
	if ghEventType == "check_suite" {
		utils.SugarLogger.Infof("Check suite event received")
		var checkSuiteEvent model.GithubCheckSuiteEvent
		if err := c.ShouldBindJSON(&checkSuiteEvent); err != nil {
			utils.SugarLogger.Errorf("Error parsing check suite event: %v", err)
			return
		}
		if checkSuiteEvent.CheckSuite.App.ClientID != config.GithubAppClientID {
			return
		}

		// Print headers
		for key, values := range c.Request.Header {
			for _, value := range values {
				utils.SugarLogger.Infof("Header %s: %s", key, value)
			}
		}

		utils.SugarLogger.Infof("Check suite event: %+v", checkSuiteEvent)
		if checkSuiteEvent.CheckSuite.Status == "completed" {
			return
		}

		go func() {
			id, err := service.CreateCheckRun(checkSuiteEvent.CheckSuite.HeadSha)
			if err != nil {
				utils.SugarLogger.Errorf("Error creating check run: %v", err)
				return
			}
			utils.SugarLogger.Infof("Check run created with ID: %s", id)

			run := model.Run{
				ID:               uuid.New().String(),
				Commit:           checkSuiteEvent.CheckSuite.HeadSha,
				Status:           "queued",
				Name:             "mpbench / unit",
				Service:          "gr25",
				GithubCheckRunID: id,
			}
			service.CreateRun(run)
			runner.StartRun(run)
		}()

	} else if ghEventType == "check_run" {
		utils.SugarLogger.Infof("Check run event received")
		var checkRunEvent model.GithubCheckRunEvent
		if err := c.ShouldBindJSON(&checkRunEvent); err != nil {
			utils.SugarLogger.Errorf("Error parsing check run event: %v", err)
			return
		}
		if checkRunEvent.CheckRun.App.ClientID != config.GithubAppClientID {
			return
		}

		// Print headers
		for key, values := range c.Request.Header {
			for _, value := range values {
				utils.SugarLogger.Infof("Header %s: %s", key, value)
			}
		}

		utils.SugarLogger.Infof("Check run event: %+v", checkRunEvent)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Github event received"})
}
