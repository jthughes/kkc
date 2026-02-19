-- name: NewPlayerTurn :one
INSERT INTO player_turns (player_id, turn_id)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetPlayerTurnsByID :many
SELECT * FROM player_turns
WHERE player_id = $1;
