package config

import "os"

type Config struct {
	Port        string
	Environment string
	// DBHost      string
	// DBPort      string
	// DBUser      string
	// DBPassword  string
	// DBName      string
	// JWTSecret   string
}

func New() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		// DBHost:      getEnv("DB_HOST", "localhost"),
		// DBPort:      getEnv("DB_PORT", "5432"),
		// DBUser:      getEnv("DB_USER", "postgres"),
		// DBPassword:  getEnv("DB_PASSWORD", ""),
		// DBName:      getEnv("DB_NAME", "rideqwik"),
		// JWTSecret:   getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
