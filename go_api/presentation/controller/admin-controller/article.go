package admincontroller

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/presentation/serializer"
	adminservice "github.com/nozomi-iida/nozo_blog/service/admin-service"
)

type ArticleController struct {
	as *adminservice.ArticleService
}

func NewArticleController(fileString string) (ArticleController, error) {
	as, err := adminservice.NewArticleService(
		adminservice.WithSqliteArticleRepository(fileString),
	)
	if err != nil {
		return ArticleController{}, err
	}
	return ArticleController{as: as}, nil
}

type ArticleRequest struct {
	Title    string     `json:"title" validate:"required"`
	Content  string     `json:"content" validate:"required"`
	IsPublic bool       `json:"isPublic"`
	TagNames []string   `json:"tagNames"`
	TopicID  *uuid.UUID `json:"topicId"`
}

type ArticleResponse struct {
	ArticleID   uuid.UUID     `json:"articleId"`
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	PublishedAt *time.Time    `json:"publishedAt,omitempty"`
	Tags        []entity.Tag  `json:"tags"`
	Topic       *entity.Topic `json:"topic,omitempty"`
	Author      entity.User   `json:"author"`
}

type ArticleListResponse struct {
	Articles []ArticleResponse `json:"articles"`
}

func articleToJson(article article.ArticleDto) ArticleResponse {
	tags := article.Tags
	if len(article.Tags) <= 0 {
		tags = []entity.Tag{}
	}
	return ArticleResponse{
		ArticleID:   article.ArticleID,
		Title:       article.Title,
		Content:     article.Content,
		PublishedAt: article.PublishedAt,
		Tags:        tags,
		Topic:       article.Topic,
		Author:      article.Author,
	}
}

func articleListToJson(articleDto article.ListArticleDto) ArticleListResponse {
	var ars = []ArticleResponse{}
	for _, a := range articleDto.Articles {
		ars = append(ars, ArticleResponse{
			ArticleID:   a.ArticleID,
			Title:       a.Title,
			Content:     a.Content,
			PublishedAt: a.PublishedAt,
			Tags:        a.Tags,
			Topic:       a.Topic,
			Author:      a.Author,
		})
	}

	return ArticleListResponse{Articles: ars}
}

func (ac *ArticleController) ListRequest(w http.ResponseWriter, r *http.Request) {
	articles, err := ac.as.List()
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	aj := articleListToJson(articles)
	output, _ := json.MarshalIndent(aj, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (ac *ArticleController) FindByIdRequest(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, queryArticleID := filepath.Split(sub)
	articleID, err := uuid.Parse(queryArticleID)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	article, err := ac.as.FindById(articleID)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	aj := serializer.Article2Json(article)
	output, _ := json.MarshalIndent(aj, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (ac *ArticleController) PatchRequest(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, queryArticleID := filepath.Split(sub)
	articleID, err := uuid.Parse(queryArticleID)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var articleRequest ArticleRequest
	json.Unmarshal(body, &articleRequest)
	if !helpers.IsValid(w, articleRequest) {
		return
	}

	a, err := ac.as.Update(
		articleID,
		articleRequest.Title,
		articleRequest.Content,
		articleRequest.TagNames,
		articleRequest.TopicID,
		articleRequest.IsPublic,
	)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(a, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
