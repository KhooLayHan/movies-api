package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/KhooLayHan/movies-api/internal/repository/postgres"
	"github.com/KhooLayHan/movies-api/internal/validator"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

// createMovieHandler handles the "POST /v1/movies" endpoint.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Basic validation
	v := validator.New()
	v.Check(input.Title != "", "title", "must be provided")
	v.Check(input.Year != 0, "year", "must be provided")
	// v.Check(input.Runtime != 0, "runtime", "must be provided")
	// v.Check(input.Genres != nil, "genres", "must be provided")

	if !v.Valid() {
		http.Error(w, "Validation Error", http.StatusUnprocessableEntity)
		return
	}

	params := postgres.CreateMovieParams{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	movie, err := app.models.Movies.Create(r.Context(), params)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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

	movie, err := app.models.Movies.Get(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Uses the new helper to write the JSON response\
	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
