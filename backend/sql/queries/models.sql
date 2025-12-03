-- name: CreateModel :one
INSERT INTO models (id, created_at, updated_at, provider_id, name, description)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetModelByName :one
SELECT * FROM models
WHERE name = $1;

-- name: GetModelByID :one
SELECT * FROM models
WHERE id = $1;

-- name: ListModels :many
SELECT * FROM models
ORDER BY name;

-- name: ListModelsByProvider :many
SELECT * FROM models
WHERE provider_id = $1
ORDER BY name;

-- name: DeleteModel :exec
DELETE FROM models
WHERE id = $1;
