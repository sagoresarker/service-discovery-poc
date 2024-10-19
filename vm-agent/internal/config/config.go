package config

type Config struct {
	HostServiceURL string
}

func LoadConfig() (*Config, error) {
	// Load configuration from file or environment variables
	// For simplicity, we'll hardcode the configuration here
	return &Config{
		HostServiceURL: "http://localhost:8092",
	}, nil
}
