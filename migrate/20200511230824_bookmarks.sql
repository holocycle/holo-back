-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE bookmarks (
  cliplist_id  VARCHAR(64) REFERENCES cliplists(id),
  user_id      VARCHAR(64) REFERENCES users(id),
  created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT bookmarks_pk PRIMARY KEY (cliplist_id, user_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE bookmarks;
