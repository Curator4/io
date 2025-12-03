-- name: CreateProvider :one
INSERT INTO providers (id, created_at, updated_at, name)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1
)
RETURNING *;

-- name: GetProvider :one
SELECT * FROM providers
WHERE name = $1;

-- name: ListProviders :many
SELECT * FROM providers
ORDER BY name;

-- name: DeleteProvider :exec
DELETE FROM providers
WHERE name = $1;
