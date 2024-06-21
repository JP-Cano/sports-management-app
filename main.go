package main

import (
	"fmt"
	"github.com/JP-Cano/sports-management-app/src/config"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/database"
	"log"
)

func main() {
	server, err := NewServer()
	defer database.Close(server.DB)
	database.Migrate(server.DB)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	port := config.Env().Port
	if port == "" {
		port = "8080"
	}
	server.Run(fmt.Sprintf(":%s", port))
}
