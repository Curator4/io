-- name: CreateConversation :one
INSERT INTO conversations (id, created_at, updated_at, name)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1
)
RETURNING *;

-- name: GetConversation :one
SELECT * FROM conversations
WHERE id = $1;

-- name: ListRecentConversations :many
SELECT * FROM conversations
ORDER BY last_used_at DESC NULLS LAST
LIMIT $1;

-- name: UpdateConversationName :one
UPDATE conversations
SET
  name = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateConversationLastUsed :exec
UPDATE conversations
SET last_used_at = NOW()
WHERE id = $1;

-- name: DeleteConversation :exec
DELETE FROM conversations
WHERE id = $1;
