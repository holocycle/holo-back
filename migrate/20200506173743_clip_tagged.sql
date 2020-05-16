-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE clip_tagged (
  clip_id    VARCHAR(64) REFERENCES clips(id),
  tag_id     VARCHAR(64) REFERENCES tags(id),
  user_id    VARCHAR(64) REFERENCES users(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT clip_tagged_pk PRIMARY KEY (clip_id, tag_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE clip_tagged;
