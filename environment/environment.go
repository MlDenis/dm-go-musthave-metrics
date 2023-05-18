package environment

import (
	"log"
	"os"
	"strconv"
)

func GetEnvString(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Printf("env: %s is empty, default: %s", envName, defaultValue)
		return defaultValue
	}
	return value
}

func GetEnvDuration(envName string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		log.Printf("env %s is empty: %s, default: %v", envName, err.Error(), defaultValue)
		return defaultValue
	}
	return value
}
