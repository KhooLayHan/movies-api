package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

// envelope is a helper type for unwrapping JSON responses.
type envelope map[string]any

// writeJSON is a helper function for sending JSON responses.
// It takes the destination http.ResponseWriter, the HTTP status code, the data to encode,
// and a header map.
func writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(js); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Uses the http.MaxBytesReader to prevent huge request bodies from crashing the server.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dst); err != nil {
		return err
	}

	return nil
}
