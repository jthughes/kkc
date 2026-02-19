-- name: CreateUser :one
INSERT INTO users (id, created_at, username)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username=$1 LIMIT 1;

-- name: ClearUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;
