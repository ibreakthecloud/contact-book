package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Simple ping-pong router to check sanity of server
func addPingRoutes(r *gin.Engine) {
	ping := r.Group("ping")

	ping.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

// addMetricsRoute that supports to ingest logs
func addContactRoute(r *gin.Engine) {
	metrics := r.Group(Contact)

	metrics.POST("/", contactController.Add)
}
