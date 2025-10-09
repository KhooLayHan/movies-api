package repository

import (
	"github.com/KhooLayHan/movies-api/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MovieRepository interface defines the methods for interacting with movie data.
type MovieRepository interface {
	// GetMovie(ctx context.Context, id int) (*Movie, error)
	// CreateMovie(ctx context.Context, movie *Movie) error
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
