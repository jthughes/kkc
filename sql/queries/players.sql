-- name: CreatePlayer :one
INSERT INTO players (user_id, game_id, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetPlayerByName :one
SELECT * FROM players
INNER JOIN users ON players.user_id=users.ID
WHERE name=$1 LIMIT 1;

-- name: GetPlayerByID :one
SELECT users.username, players.* FROM players
INNER JOIN users ON players.user_id=users.ID
WHERE players.id=$1 LIMIT 1;

-- name: ClearPlayers :exec
DELETE FROM players;

-- name: GetPlayers :many
SELECT * FROM players INNER JOIN users ON players.user_id=users.ID
WHERE game_id = $1;
