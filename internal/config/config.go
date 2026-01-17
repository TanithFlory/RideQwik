package config

import "os"

type Config struct {
	Port             string
	Environment      string
	JWTSecret        string
	UberClientID     string
	UberClientSecret string
	UberRedirectURI  string
	UberAPIBaseURL   string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBSSLMode        string
}

func New() *Config {
	return &Config{
		Port:             getEnv("PORT", "8080"),
		Environment:      getEnv("ENVIRONMENT", "development"),
		JWTSecret:        getEnv("JWT_SECRET", ""),
		UberClientID:     getEnv("UBER_CLIENT_ID", ""),
		UberClientSecret: getEnv("UBER_CLIENT_SECRET", ""),
		UberRedirectURI:  getEnv("UBER_REDIRECT_URI", ""),
		UberAPIBaseURL:   getEnv("UBER_API_BASE_URL", "https://api.uber.com"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		DBName:           getEnv("DB_NAME", "rideqwik"),
		DBSSLMode:        getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
