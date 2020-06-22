-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE cliplists (
  id             VARCHAR(64) PRIMARY KEY,
  user_id        VARCHAR(64) REFERENCES users(id),
  title          VARCHAR(255) NOT NULL,
  description    TEXT NOT NULL,
  status         VARCHAR(64) NOT NULL,
  created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE cliplists;
