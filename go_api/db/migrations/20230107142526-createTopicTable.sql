
-- +migrate Up
CREATE TABLE IF NOT EXISTS topics (
  topic_id CHAR(38) PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  description TEXT NOT NULL DEFAULT "",
  created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);

-- +migrate Down
DROP TABLE topics;
