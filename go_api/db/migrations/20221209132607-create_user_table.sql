
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  username TEXT NOT NULL UNIQUE 
);

-- +migrate Down
DROP TABLE users;
