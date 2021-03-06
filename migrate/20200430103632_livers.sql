-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE livers (
  id         VARCHAR(64) PRIMARY KEY,
  name       VARCHAR(255) NOT NULL,
  channel_id VARCHAR(64) NOT NULL,
  main_color VARCHAR(64) NOT NULL,
  sub_color  VARCHAR(64) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE livers;
