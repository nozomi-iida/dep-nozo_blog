package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/service"
)

type ArticleController struct {
	as *service.ArticleService
}

func NewArticleController(fileString string) (ArticleController, error)  {
	as, err := service.NewArticleService(
		service.WithSqliteArticleRepository(fileString),
	)
	if err != nil {
		return ArticleController{}, err
	}
	return ArticleController{as: as}, nil	
}

type ArticleRequest struct {
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	IsPublic bool `json:"isPublic" validate:"required"`
	Tags []string `json:"tags"`
}

func (ac *ArticleController) PostRequest(w http.ResponseWriter, r *http.Request)  {
	userId := r.Context().Value("userId").(uuid.UUID)
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)	
	var articleRequest ArticleRequest
	json.Unmarshal(body, &articleRequest)
	if !helpers.IsValid(w, articleRequest) {
		return
	}

	a, err := ac.as.Post(articleRequest.Title, articleRequest.Content, articleRequest.Tags ,articleRequest.IsPublic, userId)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(a, "", "\t")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
