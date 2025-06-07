package api

import (
	"mpbench/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "mpbench v" + config.Version + " is online!"})
}
