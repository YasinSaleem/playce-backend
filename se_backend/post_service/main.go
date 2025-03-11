package main

import (
	"log"
	"post_service/config"
	"post_service/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()
	r := routes.SetupRoutes()

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
