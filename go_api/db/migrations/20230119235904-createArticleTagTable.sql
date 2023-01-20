
-- +migrate Up
CREATE TABLE IF NOT EXISTS article_tags (
  article_id CHAR(38),
  tag_id CHAR(38),
  PRIMARY KEY(article_id, tag_id),
  FOREIGN KEY (article_id) REFERENCES articles(article_id)
  FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
);

-- +migrate Down
DROP TABLE article_tags;
