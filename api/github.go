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
			id, err := service.CreateCheckRun(checkSuiteEvent.CheckSuite.HeadSha)
			if err != nil {
				utils.SugarLogger.Errorf("Error creating check run: %v", err)
				return
			}
			utils.SugarLogger.Infof("Check run created with ID: %d", id)

			run := model.Run{
				ID:               uuid.New().String(),
				Commit:           checkSuiteEvent.CheckSuite.HeadSha,
				Status:           "queued",
				Name:             "mpbench / unit",
				Service:          "gr25",
				GithubCheckRunID: id,
			}
			service.CreateRun(run)
			runner.Queue.Add(run)
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
			id, err := service.CreateCheckRun(checkRunEvent.CheckRun.HeadSha)
			if err != nil {
				utils.SugarLogger.Errorf("Error creating check run: %v", err)
				return
			}
			utils.SugarLogger.Infof("Check run created with ID: %d", id)

			run := model.Run{
				ID:               uuid.New().String(),
				Commit:           checkRunEvent.CheckRun.HeadSha,
				Status:           "queued",
				Name:             "mpbench / unit",
				Service:          "gr25",
				GithubCheckRunID: id,
			}
			service.CreateRun(run)
			runner.Queue.Add(run)
		}()
	}

	c.JSON(http.StatusOK, gin.H{"message": "Github event received"})
}
