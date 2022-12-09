
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id char(38) NOT NULL DEFAULT(printf('{%s-%s-%s-%s-%s}', lower(hex(randomblob(4))), lower(hex(randomblob(2))), lower(hex(randomblob(2))), lower(hex(randomblob(2))), lower(hex(randomblob(6))))),
  username TEXT NOT NULL UNIQUE 
);

-- +migrate Down
DROP TABLE users;
