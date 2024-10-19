package config

import (
	"fmt"
	"strconv"
	"github.com/sagoresarker/svc-discovery-vm-poc/common"
)

type Config struct {
	HostServiceURL string
	AgentPort      int
}

func LoadConfig() (*Config, error) {
	common.InitEnv()

	hostIP := common.GetHostIP()
	hostPort := common.GetHostPort()
	agentPort := common.GetAgentPort()

	agentPortInt, err := strconv.Atoi(agentPort)
	if err != nil {
		return nil, fmt.Errorf("invalid AGENT_PORT: %v", err)
	}

	return &Config{
		HostServiceURL: fmt.Sprintf("http://%s:%s", hostIP, hostPort),
		AgentPort:      agentPortInt,
	}, nil
}
