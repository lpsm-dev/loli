package config

import "os"

// GetEnv function - return the value of and environmant variable if exist and if not return a fallback
func GetEnv(enviroment, fallback string) string {
	if value, okay := os.LookupEnv(enviroment); okay {
		return value
	}
	return fallback
}
