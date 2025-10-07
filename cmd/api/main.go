package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// application struct to hold application-wide state
type application struct {
	// Any dependencies required by the application
}

func main() {
	// Initialize the application struct
	app := &application{}

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
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
