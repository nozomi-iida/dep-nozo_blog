package admincontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	adminservice "github.com/nozomi-iida/nozo_blog/service/admin-service"
)

type ArticleController struct {
	as *adminservice.ArticleService
}

func NewArticleController(fileString string) (ArticleController, error)  {
	as, err := adminservice.NewArticleService(
		adminservice.WithSqliteArticleRepository(fileString),
	)
	if err != nil {
		return ArticleController{}, err
	}
	return ArticleController{as: as}, nil	
}

type ArticleResponse struct {
	ArticleID uuid.UUID `json:"articleId"`
	Title string `json:"title"`
	Content string `json:"content"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	Tags []entity.Tag `json:"tags,omitempty"`
	Topic *entity.Topic `json:"topic,omitempty"`
	Author entity.User `json:"author"`
}

type ArticleListResponse struct {
	Articles []ArticleResponse `json:"articles"`
}

func articleListToJson(articleDto article.ListArticleDto) ArticleListResponse {
	var ars = []ArticleResponse{}
	for _, a := range articleDto.Articles {
		ars = append(ars, ArticleResponse{
			ArticleID: a.ArticleID, 
			Title: a.Title, 
			Content: a.Content, 
			PublishedAt: a.PublishedAt, 
			Tags: a.Tags, 
			Topic: a.Topic, 
			Author: a.Author,
		})
	}

	return ArticleListResponse{Articles: ars}
}

func (ac *ArticleController) ListRequest(w http.ResponseWriter, r *http.Request)  {
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
