-- +goose Up
CREATE TABLE messages (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  role TEXT NOT NULL,
  context TEXT,
  conversation_id UUID REFERENCES conversations(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE messages;
