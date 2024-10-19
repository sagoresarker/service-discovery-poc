package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HostServiceURL string
	AgentPort      int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	hostIP := os.Getenv("HOST_IP")
	if hostIP == "" {
		return nil, fmt.Errorf("HOST_IP environment variable is not set")
	}

	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		hostPort = "8092" // Default port if not specified
	}

	agentPort := os.Getenv("AGENT_PORT")
	if agentPort == "" {
		agentPort = "8093" // Default port if not specified
	}

	return &Config{
		HostServiceURL: fmt.Sprintf("http://%s:%s", hostIP, hostPort),
		AgentPort:      8093,
	}, nil
}
