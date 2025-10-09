-- name: CreateMovie :one
INSERT INTO movies (title, year, runtime, genres)
VALUES ($1, $2, $3, $4)
RETURNING id, version, created_at;

-- name: GetMovie :one
SELECT id, title, year, runtime, genres, version, created_at, updated_at
FROM movies
WHERE id = $1;

-- name: UpdateMovie :one
UPDATE movies
SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1, updated_at = NOW()
WHERE id = $5
RETURNING version, updated_at;

-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1;
