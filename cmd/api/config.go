package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the configuration for the application.
type Config struct {
	Port int
	DB   struct {
		DSN string
	}
}

// LoadConfig loads the application configuration from environment variables.
func LoadConfig() (Config, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config

	// Parse the server port
	portStr := os.Getenv("API_PORT")
	if portStr == "" {
		log.Fatal("API_PORT environment variable not set")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid API_PORT environment variable")
		return Config{}, err
	}
	cfg.Port = port

	// Parse the database DSN
	dbConnection := os.Getenv("DB_CONNECTION")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	cfg.DB.DSN = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbConnection, dbUser, dbPassword, dbHost, dbPort, dbName)
	if cfg.DB.DSN == "" {
		log.Fatal("DB_DSN environment variable not set")
		return Config{}, err
	}

	return cfg, nil
}
