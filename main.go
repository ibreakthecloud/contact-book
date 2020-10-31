package main

import (
	"github.com/ibreakthecloud/contact-book/router"
	"log"
	"os"
)

var (
	port = "9000"
)

func main() {

	// check if port is provided in environment
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// 1. init the store(database)

	// 2. init the router
	r := router.New()

	// 3. Listen constantly on any port
	log.Print("LISTENING ON PORT: ", port)
	log.Fatal(r.Run(":" + port))
}
