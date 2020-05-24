-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE clips (
  id          VARCHAR(64) PRIMARY KEY,
  user_id     VARCHAR(64) REFERENCES users(id),
  title       VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  video_id    VARCHAR(64) REFERENCES videos(id),
  begin_at    INTEGER NOT NULL DEFAULT 0,
  end_at      INTEGER NOT NULL DEFAULT 0,
  status      VARCHAR(64) NOT NULL,
  created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX clip_created_at_index ON clips(created_at);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE clips;
