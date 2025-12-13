-- name: CreateMessage :one
INSERT INTO messages (id, created_at, updated_at, conversation_id, user_id, role, content)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetMessagesByConversation :many
SELECT
  m.id,
  m.created_at,
  m.updated_at,
  m.conversation_id,
  m.user_id,
  m.role,
  m.content,
  u.id as joined_user_id,
  u.name as joined_user_name,
  u.created_at as joined_user_created_at,
  u.updated_at as joined_user_updated_at
FROM messages m
LEFT JOIN users u ON m.user_id = u.id
WHERE m.conversation_id = $1
ORDER BY m.created_at ASC;
