-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO livers (id, name, channel_id, main_color, sub_color) VALUES
  ('sakuramiko', 'さくらみこ', 'UC-hM6YJuNYVAmUWxeIr9FeA', '#FEA5D1', '#F7002F'),
  ('houshoku-marine', '宝鐘マリン', 'UCCzUftO8KOVkV4wQG1vkUvg', '#F7002F', '#FF7608');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

TRUNCATE TABLE livers;
