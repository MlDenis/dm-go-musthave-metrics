package environment

import (
	"log"
	"os"
	"time"
)

func GetEnvString(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Printf("env: %s is empty, default: %s", envName, defaultValue)
		return defaultValue
	}
	return value
}

func GetEnvDuration(envName string, defaultValue time.Duration) time.Duration {
	value, err := time.ParseDuration(os.Getenv(envName))
	if err != nil {
		log.Printf("env %s is empty: %s, default: %v", envName, err.Error(), defaultValue)
		return defaultValue
	}
	return value
}
