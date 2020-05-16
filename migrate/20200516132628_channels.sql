-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE channels (
  id VARCHAR(64)       PRIMARY KEY,
  title                TEXT NOT NULL,
  description          TEXT NOT NULL,
  small_thumbnail_url  VARCHAR(255) DEFAULT '',
  medium_thumbnail_url VARCHAR(255) DEFAULT '',
  large_thumbnail_url  VARCHAR(255) DEFAULT '',
  small_banner_url     VARCHAR(255) DEFAULT '',
  medium_banner_url    VARCHAR(255) DEFAULT '',
  large_banner_url     VARCHAR(255) DEFAULT '',
  view_count           BIGINT DEFAULT 0,
  comment_count        BIGINT DEFAULT 0,
  subscriber_count     BIGINT DEFAULT 0,
  video_count          BIGINT DEFAULT 0,
  published_at         TIMESTAMP WITH TIME ZONE NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE channels
