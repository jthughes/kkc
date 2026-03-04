-- name: CreatePlayerComplaint :one
INSERT INTO complaints (action_id, target_id)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetPlayerComplaints :one
SELECT * FROM complaints
WHERE action_id=$1;

-- name: DeletePlayerComplaint :exec
DELETE FROM complaints
WHERE action_id=$1;
