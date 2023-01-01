
-- +migrate Up
CREATE TABLE IF NOT EXISTS article_tags (
  article_id CHAR(38),
  name TEXT,
  PRIMARY KEY(article_id, name),
  FOREIGN KEY (article_id) REFERENCES articles(article_id)
);

-- +migrate Down
DROP TABLE article_tags;
