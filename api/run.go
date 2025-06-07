package api

import (
	"mpbench/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRuns(c *gin.Context) {
	limit := c.Query("limit")
	if limit == "" {
		runs := service.GetAllRuns()
		c.JSON(http.StatusOK, runs)
	} else {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		runs := service.GetMostRecentRuns(limitInt)
		c.JSON(http.StatusOK, runs)
	}
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
