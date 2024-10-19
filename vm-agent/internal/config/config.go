package config

type Config struct {
	HostServiceURL string
	AgentPort      int
}

func LoadConfig() (*Config, error) {
	return &Config{
		HostServiceURL: "http://localhost:8092",
		AgentPort:      8093,
	}, nil
}
