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
	contact := r.Group(Contact, basicAuth())

	contact.GET("/", contactController.Get)
	contact.POST("/", contactController.Add)
	contact.PUT("/", contactController.Update)
	contact.DELETE("/", contactController.Delete)
}

// basicAuth adds basic authentication
func basicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		AuthUserName:    AuthPassword,
	})
}