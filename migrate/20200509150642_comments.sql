-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE comments (
  id         VARCHAR(64) PRIMARY KEY,
  user_id    VARCHAR(64) REFERENCES users(id),
  clip_id    VARCHAR(64) REFERENCES clips(id),
  content    VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX comment_clip_id_index ON comments(clip_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE comments;
