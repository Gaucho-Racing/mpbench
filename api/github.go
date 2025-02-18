package api

import (
	"mpbench/config"
	"mpbench/model"
	"mpbench/service"
	"mpbench/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubEventHandler(c *gin.Context) {
	// // Read and print body
	// body, err := c.GetRawData()
	// if err != nil {
	// 	utils.SugarLogger.Errorf("Error reading body: %v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
	// 	return
	// }
	// utils.SugarLogger.Infof("Body: %s", string(body))

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

		go func() {
			id, err := service.CreateCheckRun(checkSuiteEvent.CheckSuite.HeadSha)
			if err != nil {
				utils.SugarLogger.Errorf("Error creating check run: %v", err)
				return
			}
			utils.SugarLogger.Infof("Check run created with ID: %s", id)
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
