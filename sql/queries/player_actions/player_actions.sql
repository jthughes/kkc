-- name: CreatePlayerAction :one
INSERT INTO actions (player_turn_id)
VALUES (
    $1
)
RETURNING *;

-- name: GetPlayerActionByPlayerTurnID :one
SELECT * FROM actions
WHERE player_turn_id=$1 LIMIT 1;

-- name: UpdatePlayerAction :one
UPDATE actions SET (lodging, visit_imre, attend_university)
= (
    $2,
    $3,
    $4
) WHERE player_turn_id=$1
RETURNING *;
