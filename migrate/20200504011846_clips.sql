-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE clips (
  id          VARCHAR(64) PRIMARY KEY,
  user_id     VARCHAR(64) NOT NULL,
  title       VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  video_id    VARCHAR(64) NOT NULL,
  begin_at    INTEGER NOT NULL,
  end_at      INTEGER NOT NULL,
  created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE clips;
