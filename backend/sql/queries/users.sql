-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
