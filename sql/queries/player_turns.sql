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

-- name: GetPlayerTurn :one
SELECT pt.* FROM player_turns pt
JOIN game_turns gt ON pt.turn_id = gt.id
WHERE pt.player_id = $1 AND gt.term = $2 AND gt.month = $3;


-- name: GetLastPlayerTurn :one
SELECT * FROM player_turns pt
    JOIN actions ON actions.id = pt.id
    JOIN elevation_points ep ON ep.action_id = pt.id
WHERE pt.player_id = $1
ORDER BY pt.created_at DESC LIMIT 1;
