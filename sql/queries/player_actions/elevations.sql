-- name: CreatePlayerEPSubmission :one
INSERT INTO elevation_points (action_id)
VALUES (
    $1
)
RETURNING *;

-- name: GetPlayerEPByPlayerActionID :one
SELECT * FROM elevation_points
WHERE action_id=$1 LIMIT 1;

-- name: UpdatePlayerEPSubmission :one
UPDATE elevation_points SET (ep_linguistics, ep_arithmetics, ep_rhetoric_and_logic, ep_archives, ep_sympathy, ep_physicking, ep_alchemy, ep_artificery, ep_naming)
= (
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
) WHERE action_id=$1
RETURNING *;
