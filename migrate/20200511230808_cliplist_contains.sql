-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE cliplist_contains (
  cliplist_id VARCHAR(64) REFERENCES cliplists(id),
  index       INTEGER NOT NULL DEFAULT 0,
  clip_id     VARCHAR(64) REFERENCES clips(id),
  created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT cliplist_contains_pk PRIMARY KEY (cliplist_id, index)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE cliplist_contains;
