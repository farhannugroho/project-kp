package main

import (
	"log"
	"project-kp/config"
	"project-kp/router"
)

func main() {
	// Init DB
	config.InitDb()

	// init router
	r := router.InitRouter()

	// start server
	log.Fatal(r.Run())
}
