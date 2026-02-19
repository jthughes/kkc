-- name: CreatePlayer :one
INSERT INTO players (id, created_at, user_id, game_id, name, skindancer, class)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    false,
    NULL
)
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players
WHERE name=$1 LIMIT 1;

-- name: ClearPlayers :exec
DELETE FROM players;

-- name: GetPlayers :many
SELECT * FROM players;
