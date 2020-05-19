-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE videos (
  id                   VARCHAR(64) PRIMARY KEY,
  channel_id           VARCHAR(64) NOT NULL,
  title                TEXT NOT NULL DEFAULT '',
  description          TEXT NOT NULL DEFAULT '',
  duration             INTEGER NOT NULL DEFAULT 0,
  small_thumbnail_url  VARCHAR(255) NOT NULL DEFAULT '',
  medium_thumbnail_url VARCHAR(255) NOT NULL DEFAULT '',
  large_thumbnail_url  VARCHAR(255) NOT NULL DEFAULT '',
  published_at         TIMESTAMP WITH TIME ZONE NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE videos;
