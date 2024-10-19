package main

import (
	"log"

	"github.com/sagoresarker/svc-discovery-host-poc/internal/config"
	"github.com/sagoresarker/svc-discovery-host-poc/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	server := server.NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
