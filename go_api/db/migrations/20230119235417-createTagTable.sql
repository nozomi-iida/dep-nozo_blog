
-- +migrate Up
CREATE TABLE IF NOT EXISTS tags (
  tag_id CHAR(38) PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE tags;
