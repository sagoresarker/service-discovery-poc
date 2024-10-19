package config

type Config struct {
	Port int
}

func LoadConfig() (*Config, error) {
	// Load configuration from file or environment variables
	// For simplicity, we'll hardcode the configuration here
	return &Config{
		Port: 8092,
	}, nil
}
