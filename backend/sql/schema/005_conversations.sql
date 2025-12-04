-- +goose Up
CREATE TABLE conversations (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  last_used_at TIMESTAMP,
  name TEXT
);

-- +goose Down
DROP TABLE conversations;
