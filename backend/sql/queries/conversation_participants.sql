-- name: AddParticipant :exec
INSERT INTO conversation_participants (conversation_id, user_id)
VALUES ($1, $2);

-- name: RemoveParticipant :exec
DELETE FROM conversation_participants
WHERE conversation_id = $1 AND user_id = $2;

-- name: GetConversationParticipants :many
SELECT
	cp.conversation_id,
	cp.user_id,
	cp.joined_at
FROM conversation_participants cp
WHERE cp.conversation_id = $1;

-- name: GetUserConversations :many
SELECT c.* FROM conversations c
JOIN conversation_participants cp ON c.id = cp.conversation_id
WHERE cp.user_id = $1
ORDER BY c.updated_at DESC;
