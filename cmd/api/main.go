package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Register the healthcheck handler for the "/v1/healthcheck" route.
	mux.HandleFunc("/v1/healthcheck", healthCheckHandler)

	// Server port to listen on.
	port := 4040

	// Print a message indicating the server is starting
	log.Printf("Starting server on port %d", port)

	// Start the HTTP server with our server mux.
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
