package main

import (
	server "cachedapi/cmd"
	"cachedapi/pkg/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Unable to load env file: %v. Loading default values.", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load application configuration: %v", err)
	}

	server.Run(cfg)
}
