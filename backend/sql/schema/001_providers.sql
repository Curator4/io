-- +goose Up
CREATE TABLE providers (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE providers;
