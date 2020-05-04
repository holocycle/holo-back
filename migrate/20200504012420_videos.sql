-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE videos (
  id                   VARCHAR(64) PRIMARY KEY,
  channel_id           VARCHAR(64) NOT NULL,
  title                TEXT NOT NULL,
  description          TEXT NOT NULL,
  duration             INTEGER NOT NULL,
  small_thumbnail_url  VARCHAR(255),
  medium_thumbnail_url VARCHAR(255),
  large_thumbnail_url  VARCHAR(255),
  published_at         TIMESTAMP WITH TIME ZONE NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE videos;
