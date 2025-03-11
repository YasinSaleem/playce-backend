package main

import (
	"log"
	"user_service/config"
	"user_service/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
