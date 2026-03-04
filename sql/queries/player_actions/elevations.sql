-- name: CreatePlayerEPSubmission :one
INSERT INTO elevation_points (action_id)
VALUES (
    $1
)
RETURNING *;

-- name: GetPlayerEPByPlayerActionID :one
SELECT * FROM elevation_points
WHERE action_id=$1 LIMIT 1;

-- name: UpdatePlayerEPSubmission :exec
UPDATE elevation_points SET (ep_alchemy, ep_archives, ep_arithmetics, ep_artificery, ep_linguistics, ep_naming, ep_physicking, ep_rhetoric_and_logic, ep_sympathy)
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
) WHERE action_id=$1;
