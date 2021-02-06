package config

import "os"

// GetEnv function - return the value of and environment variable if exist and if not return a fallback.
func GetEnv(environment, fallback string) string {
	if value, okay := os.LookupEnv(environment); okay {
		return value
	}
	return fallback
}
