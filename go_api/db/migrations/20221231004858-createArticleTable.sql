
-- +migrate Up
CREATE TABLE IF NOT EXISTS articles (
  article_id CHAR(38) PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  published_at DATETIME,
  -- external_url TEXT,
  created_at DATETIME NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  updated_at DATETIME NOT NULL DEFAULT (DATETIME('now', 'localtime')),
  author_id CHAR(38) NOT NULL,
  Foreign Key (author_id) REFERENCES users(user_id)
);

-- +migrate Down
DROP Table articles;
