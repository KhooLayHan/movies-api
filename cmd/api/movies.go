package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KhooLayHan/movies-api/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// createMovieHandler handles the "POST /v1/movies" endpoint.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "create a new movie"); err != nil {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}
}

// getMovieHanlder handles the "GET /v1/movies/{id}" endpoint.
func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Creates a dummy movie for now
	movie := postgres.Movie{
		ID:        id,
		Title:     "Casablanca",
		Year:      1942,
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		CreatedAt: pgtype.Timestamp{},
		UpdatedAt: pgtype.Timestamp{},
	}

	// Uses the new helper to write the JSON response
	err = writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
