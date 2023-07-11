package main

import (
	"log"

	config "github.com/LorezV/gophkeeper/config/server"
	server "github.com/LorezV/gophkeeper/internal/server/app"
)

// @title Gophkeeper Server
// @version 1.0.0
// @description Gophkeeper project
// @contact.name Derkach Dmitriy
// @contact.url https://github.com/LorezV
// @contact.email dima.derkach2004@gmail.com
// @host localhost:8080
// @BasePath /api/v1
// Main func.
func main() {
	log.Println("Server App")

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	server.Run(cfg)
}
