package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Simple handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	// Server port to listen on
	port := 4040

	// Print a message indicating the server is starting
	log.Printf("Starting server on port %d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
