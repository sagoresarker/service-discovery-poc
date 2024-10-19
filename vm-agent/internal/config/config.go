package config

import (
	"fmt"
	"strconv"
)

type Config struct {
	HostServiceURL string
	AgentPort      int
}

func LoadConfig() (*Config, error) {
	InitEnv()

	hostIP := GetHostIP()
	hostPort := GetHostPort()
	agentPort := GetAgentPort()

	agentPortInt, err := strconv.Atoi(agentPort)
	if err != nil {
		return nil, fmt.Errorf("invalid AGENT_PORT: %v", err)
	}

	return &Config{
		HostServiceURL: fmt.Sprintf("http://%s:%s", hostIP, hostPort),
		AgentPort:      agentPortInt,
	}, nil
}
