-- name: CreateUser :one
INSERT INTO users(id, name, createdAt, updatedAt)
VALUES ($1, $2, $3, $4)
RETURNING *;