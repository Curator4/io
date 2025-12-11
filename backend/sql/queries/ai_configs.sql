-- name: CreateAIConfig :one
INSERT INTO ai_configs (id, created_at, updated_at, name, model_id, system_prompt)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetAIConfigByName :one
SELECT * FROM ai_configs
WHERE name = $1;

-- name: GetAIConfigByID :one
SELECT
  ac.*,
  sqlc.embed(m)
FROM ai_configs ac
JOIN models m ON ac.model_id = m.id
WHERE ac.id = $1;

-- name: ListAIConfigs :many
SELECT * FROM ai_configs
ORDER BY name;

-- name: UpdateAIConfigModel :one
UPDATE ai_configs
SET
  model_id = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateAIConfigPrompt :one
UPDATE ai_configs
SET
  system_prompt = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateAIConfigLastUsed :exec
UPDATE ai_configs
SET last_used_at = NOW()
WHERE id = $1;

-- name: DeleteAIConfig :exec
DELETE FROM ai_configs
WHERE id = $1;
