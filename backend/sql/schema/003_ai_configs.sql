-- +goose Up
CREATE TABLE ai_configs (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  last_used_at TIMESTAMP,
  name TEXT NOT NULL UNIQUE,
  model_id UUID NOT NULL REFERENCES models(id) ON DELETE CASCADE,
  system_prompt TEXT
);

-- +goose Down
DROP TABLE ai_configs;
