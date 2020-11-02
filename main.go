package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/contact-book/router"
	"github.com/ibreakthecloud/contact-book/store"
	"github.com/ibreakthecloud/contact-book/store/sqlite"
	"log"
	"os"
)

var (
	port = "80"
)

func init() {

	// init the store(database)
	store.NewStore = sqlite.New("")
}

func main() {

	// check if port is provided in environment
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// init the router
	r := router.New()

	r.Use(gin.Logger())

	// Listen constantly on given port
	log.Print("LISTENING ON PORT: ", port)
	log.Fatal(r.Run(":" + port))
}
