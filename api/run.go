package api

import (
	"mpbench/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRuns(c *gin.Context) {
	runs := service.GetAllRuns()
	c.JSON(http.StatusOK, runs)
}

func GetRunByID(c *gin.Context) {
	id := c.Param("id")
	run := service.GetRunByID(id)
	if run.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "No run found with id: " + id})
		return
	}
	c.JSON(http.StatusOK, run)
}
