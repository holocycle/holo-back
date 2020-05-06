-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE clip_tags (
  id         VARCHAR(64) PRIMARY KEY,
  user_id    VARCHAR(64) REFERENCES users(id),
  clip_id    VARCHAR(64) REFERENCES clips(id),
  tag_id     VARCHAR(64) REFERENCES tags(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX clip_tags_clip_id_index ON clip_tags(clip_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE clip_tags;
