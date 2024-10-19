package main

import (
	"log"

	"github.com/sagoresarker/svc-discovery-vm-poc/internal/agent"
	"github.com/sagoresarker/svc-discovery-vm-poc/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	agent := agent.NewAgent(cfg)
	if err := agent.Run(); err != nil {
		log.Fatalf("Agent failed: %v", err)
	}
}
