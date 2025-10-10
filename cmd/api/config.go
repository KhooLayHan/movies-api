package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config represents the configuration for the application.
type Config struct {
	Port int
	Env  string // To store the application environment (e.g. "development", "staging", and "production")
	DB   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  time.Duration
	}
}

// LoadConfig loads the application configuration from environment variables.
func LoadConfig() Config {
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
		return Config{}
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
		return Config{}
	}

	// Server settings
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development, staging, production)")

	// Database settings
	flag.StringVar(&cfg.DB.DSN, "db-dsn", cfg.DB.DSN, "PostgreSQL DSN")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", time.Minute, "PostgreSQL max idle time")

	return cfg
}
