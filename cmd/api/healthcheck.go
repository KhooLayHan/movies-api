package main

import (
	"encoding/json"
	"net/http"
)

// healthCheckHandler writes a plain-text response indicating the status of the API
// with information about the application status, operating environment, and version.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map to hold our healthcheck data
	data := map[string]string{
		"status":      "available",
		"environment": "development",
		"version":     "1.0.0",
	}

	// Marshal the map into a JSON object.
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Adds a newline to the JSON for easier viewing in terminal clients.
	js = append(js, '\n')

	// Sets the `Content-Type` header to `application/json`.
	w.Header().Set("Content-Type", "application/json")

	// Writes the JSON response.
	if _, err := w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
