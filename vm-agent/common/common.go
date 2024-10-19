package common

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func InitEnv() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}
	currentDir := filepath.Dir(filename)

	for i := 0; i < 5; i++ {
		envPath := filepath.Join(currentDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			err := godotenv.Load(envPath)
			if err != nil {
				log.Fatalf("Error loading .env file from path %s: %v", envPath, err)
			}
			log.Printf("Loaded .env file from: %s", envPath)
			return
		}
		currentDir = filepath.Dir(currentDir)
	}

	log.Fatal("Could not find .env file in the project directory or its parents")
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environment variable %s is not set", key)
	}
	return value
}

func CheckEnvVars(vars []string) {
	var missingEnvVars []string

	for _, envVar := range vars {
		if os.Getenv(envVar) == "" {
			missingEnvVars = append(missingEnvVars, envVar)
		}
	}

	if len(missingEnvVars) > 0 {
		log.Fatalf("The following required environment variables are not set: %s",
			strings.Join(missingEnvVars, ", "))
	}
}

func GetHostIP() (hostIP string) {
	requiredEnvVars := []string{"HOST_IP"}
	CheckEnvVars(requiredEnvVars)

	hostIP = GetEnv("HOST_IP")

	return hostIP
}

func GetHostPort() (hostPort string) {
	requiredEnvVars := []string{"HOST_PORT"}
	CheckEnvVars(requiredEnvVars)

	hostPort = GetEnv("HOST_PORT")

	return hostPort
}

func GetAgentPort() (agentPort string) {
	requiredEnvVars := []string{"AGENT_PORT"}
	CheckEnvVars(requiredEnvVars)

	agentPort = GetEnv("AGENT_PORT")

	return agentPort
}
