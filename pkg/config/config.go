package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds database configuration values
type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	DBSSLMode  string
}

// LoadConfig reads configuration from environment variables or a .env file
func LoadConfig() (*Config, error) {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	cfg := &Config{
		DBUser:     getEnv("DB_USER", "root_user"),
		DBPassword: getEnv("DB_PASSWORD", "Dost0n1k"),
		DBName:     getEnv("DB_NAME", "parking_searcher"),
		DBHost:     getEnv("DB_HOST", "172.17.0.2"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	return cfg, nil
}

// getEnv retrieves environment variables with a fallback default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// DSN returns the PostgreSQL connection string
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode,
	)
}
