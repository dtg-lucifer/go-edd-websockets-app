package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoute(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"status": "OK",
	})
}
