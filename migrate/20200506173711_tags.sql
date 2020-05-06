-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE tags (
  id         VARCHAR(64)  PRIMARY KEY,
  name       VARCHAR(255) NOT NULL UNIQUE,
  color      VARCHAR(64)  NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX tag_name_index ON tags(name);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE tags;
