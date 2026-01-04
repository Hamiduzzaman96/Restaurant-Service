package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser, DBPassword, DBHost, DBName, DBPort string
	GRPCPort                                   string
	Timeout                                    time.Duration
	Retries                                    int
	Delay                                      time.Duration
}

func LoadConfig(service string) *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system env")
	}

	timeout, _ := time.ParseDuration(getEnv("DEFAULT_TIMEOUT", "3s"))
	delay, _ := time.ParseDuration(getEnv("DEFAULT_DELAY", "500ms"))

	return &Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "restaurant_db"),
		GRPCPort:   getEnv(service+"GRPC_PORT", ":50051"),
		Timeout:    timeout,
		Retries:    getEnvAsInt("DEFAULT_RETRIES", 3),
		Delay:      delay,
	}

}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		var v int
		_, err := fmt.Sscan(value, &v)
		if err == nil {
			return v
		}
	}
	return fallback
}
