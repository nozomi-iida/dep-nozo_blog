
-- +migrate Up
ALTER TABLE articles
ADD topic_id CHAR(38)
REFERENCES topic(id);

-- +migrate Down
ALTER TABLE articles
DROP topic_id;
-- DROP FOREIGN KEY FK_TopicArticle;
