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
  m.*,
  sqlc.embed(u)
FROM messages m
LEFT JOIN users u ON m.user_id = u.id
WHERE m.conversation_id = $1
ORDER BY m.created_at ASC;
