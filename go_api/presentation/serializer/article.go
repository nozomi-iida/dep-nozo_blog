package serializer

import (
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
)

type ArticleJson struct {
	ArticleID   uuid.UUID  `json:"articleId"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	Tags        []TagJson  `json:"tags"`
	Topic       *TopicJson `json:"topic,omitempty"`
	Author      UserJson   `json:"author"`
}

func Article2Json(article article.ArticleDto) ArticleJson {
	tags := []TagJson{}
	if len(article.Tags) > 0 {
		for _, tag := range article.Tags {
			tags = append(tags, Tag2Json(tag))
		}
	}
	var topic TopicJson
	if article.Topic != nil {
		topic = Topic2Json(*article.Topic)
	}

	return ArticleJson{
		ArticleID:   article.ArticleID,
		Title:       article.Title,
		Content:     article.Content,
		PublishedAt: article.PublishedAt,
		Tags:        tags,
		Topic:       &topic,
		Author:      User2Json(article.Author),
	}
}
