-- name: CreatePlayerComplaint :one
INSERT INTO complaints (action_id, target_id)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetPlayerComplaints :one
SELECT * FROM elevation_points
WHERE action_id=$1;

-- name: DeletePlayerComplaint :exec
DELETE FROM elevation_points
WHERE id=$1;
