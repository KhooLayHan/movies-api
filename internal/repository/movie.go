package repository

import (
	"context"

	"github.com/KhooLayHan/movies-api/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MovieRepository interface defines the methods for interacting with movie data.
type MovieRepository interface {
	Create(ctx context.Context, params postgres.CreateMovieParams) (postgres.CreateMovieRow, error)
	Get(ctx context.Context, id int64) (postgres.Movie, error)

	// UpdateMovie(ctx context.Context, id int, movie *Movie) error
	// DeleteMovie(ctx context.Context, id int) error
}

// PostgresMovieRepo is the concrete implementation for PostgreSQL.
type PostgresMovieRepo struct {
	DB *postgres.Queries
}

// NewMovieRepository creates a new movie repository with the given database connection.
func NewMovieRepository(db *pgxpool.Pool) *PostgresMovieRepo {
	return &PostgresMovieRepo{
		DB: postgres.New(db),
	}
}

// Create implements the MovieRepository interface.
func (r *PostgresMovieRepo) Create(ctx context.Context, params postgres.CreateMovieParams) (postgres.CreateMovieRow, error) {
	return r.DB.CreateMovie(ctx, params)
}

// Get implements the MovieRepository interface.
func (r *PostgresMovieRepo) Get(ctx context.Context, id int64) (postgres.Movie, error) {
	return r.DB.GetMovie(ctx, id)
}

// Update implements the MovieRepository interface.
func (r *PostgresMovieRepo) Update(ctx context.Context, params postgres.UpdateMovieParams) (postgres.UpdateMovieRow, error) {
	return r.DB.UpdateMovie(ctx, params)
}

// Delete implements the MovieRepository interface.
func (r *PostgresMovieRepo) Delete(ctx context.Context, id int64) error {
	return r.DB.DeleteMovie(ctx, id)
}
