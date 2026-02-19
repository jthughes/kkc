-- name: NewGameTurn :one
INSERT INTO game_turns (game_id, name, term, month)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetGameTurns :many
SELECT * FROM game_turns
WHERE game_id = $1;
