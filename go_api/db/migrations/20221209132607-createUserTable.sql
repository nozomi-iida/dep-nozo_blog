
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  user_id char(38) PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

-- +migrate Down
DROP TABLE users;
