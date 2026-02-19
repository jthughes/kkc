-- name: NewPlayerStatus :one
INSERT INTO player_status (player_id, turn_id, sane, crockery, lodging, imre, university, medica, coin, ep_linguistics, ep_arithmetics, ep_rhetoric_and_logic, ep_archives, ep_sympathy, ep_physicking, ep_alchemy, ep_artificery, ep_naming)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18
)
RETURNING *;

-- name: GetAllPlayerStatus :many
SELECT * FROM player_status;

-- name: GetPlayerStatusByID :one
SELECT * FROM player_status
WHERE player_id = $1 AND turn_id = $2;
