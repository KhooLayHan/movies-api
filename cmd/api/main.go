package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// application struct to hold application-wide state
type application struct {
	config Config
	db     *pgxpool.Pool
}

func main() {
	// Load the configuration.
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database.
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Defers the database call so that the connection pool closes before exiting the main() function.
	defer db.Close()
	log.Println("Database connection pool established.")

	// Initialize the application struct.
	app := &application{
		config: cfg,
		db:     db,
	}

	// Initialize the chi router.
	router := chi.NewRouter()

	// Register the healthcheck handler for the "/v1/healthcheck" route.
	router.Get("/v1/healthcheck", app.healthCheckHandler)
	router.Post("/v1/healthcheck", app.createMovieHandler)
	router.Get("/v1/healthcheck/{id}", app.getMovieHandler)

	// Server port to listen on.
	port := 4040

	// Print a message indicating the server is starting
	log.Printf("Starting server on port %d", port)

	// Creates the HTTP server with chi router
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Start the HTTP server with our server mux.
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// openDB creates a new connection pool for a given configuration.
func openDB(cfg Config) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Creates a context with a 5-second timeout to test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the database to test the connection
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
