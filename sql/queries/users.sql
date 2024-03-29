-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserApiKeyByID :one
SELECT api_key FROM users
WHERE id = $1;

-- name: GetUserByApiKey :one
SELECT * FROM users
WHERE api_key = $1;
