-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE favorites (
  clip_id    VARCHAR(64) REFERENCES clips(id),
  user_id    VARCHAR(64) REFERENCES users(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT favorites_pk PRIMARY KEY (clip_id, user_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE favorites;
