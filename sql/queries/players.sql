-- name: CreatePlayer :one
INSERT INTO players (user_id, game_id, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players
WHERE name=$1 LIMIT 1;

-- name: ClearPlayers :exec
DELETE FROM players;

-- name: GetPlayers :many
SELECT * FROM players;
