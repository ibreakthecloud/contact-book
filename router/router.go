package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/contact-book/controller/contact"
	"os"
)

// Endpoints constants
const (
	Contact = "contact"
)

var (
	AuthUserName = os.Getenv("username")
	AuthPassword = os.Getenv("password")
)

var (
	contactController = contact.New()
)

// New instantiates a new gin router to handle API requests
func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	registerAllEndpoints(router)

	return router
}

// registerAllEndpoints registers all of the endpoints supported by the server
func registerAllEndpoints(r *gin.Engine) {

	// registers the below endpoints
	addPingRoutes(r)
	addContactRoute(r)
}
