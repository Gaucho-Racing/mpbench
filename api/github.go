package api

import (
	"mpbench/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubEventHandler(c *gin.Context) {
	// Print headers
	for key, values := range c.Request.Header {
		for _, value := range values {
			utils.SugarLogger.Infof("Header %s: %s", key, value)
		}
	}

	// Read and print body
	body, err := c.GetRawData()
	if err != nil {
		utils.SugarLogger.Errorf("Error reading body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	utils.SugarLogger.Infof("Body: %s", string(body))

	ghEventType := c.Request.Header.Get("X-GitHub-Event")
	if ghEventType == "check_suite" {
		utils.SugarLogger.Infof("Check suite event received")
		var checkSuiteEvent struct {
			Action string `json:"action"`
		}
		if err := c.ShouldBindJSON(&checkSuiteEvent); err != nil {
			utils.SugarLogger.Errorf("Error parsing check suite event: %v", err)
			return
		}

		if checkSuiteEvent.Action == "requested" || checkSuiteEvent.Action == "rerequested" {
			utils.SugarLogger.Infof("Check suite %s", checkSuiteEvent.Action)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Github event received"})
}
