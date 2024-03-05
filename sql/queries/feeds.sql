-- name: GetUserFeeds :many
SELECT * FROM feeds
WHERE user_id = $1
ORDER BY updated_at DESC;

-- name: CreateFeed :one
INSERT INTO feeds (created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
