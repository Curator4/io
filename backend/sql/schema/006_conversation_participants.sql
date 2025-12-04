-- +goose Up
CREATE TABLE conversation_participants (
  conversation_id UUID NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  joined_at TIMESTAMP NOT NULL DEFAULT now(),
  PRIMARY KEY (conversation_id, user_id)
);

-- +goose Down
DROP TABLE conversation_participants;
