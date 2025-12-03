-- name: CreateMessage :one
INSERT INTO messages (id, created_at, updated_at, conversation_id, role, content)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetMessagesByConversation :many
SELECT * FROM messages
WHERE conversation_id = $1
ORDER BY created_at ASC;
