-- name: CreateAsset :one
INSERT INTO assets (id, user_id, name, type, current_value, purchased_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetAssetsByUserID :many
SELECT * FROM assets WHERE user_id = $1 LIMIT $2 OFFSET $3;

