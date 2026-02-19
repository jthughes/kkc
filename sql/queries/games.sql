-- name: CreateGame :one
INSERT INTO games (id, created_at, game_master, name, type, type_number)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetGamer :one
SELECT * FROM games
WHERE id=$1 LIMIT 1;

-- name: ClearGames :exec
DELETE FROM games;

-- name: GetGames :many
SELECT * FROM games;
