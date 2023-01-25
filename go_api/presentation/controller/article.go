package controller

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

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
	TopicID *uuid.UUID `json:"topicId"`
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

	a, err := ac.as.Post(articleRequest.Title, articleRequest.Content, articleRequest.Tags ,articleRequest.IsPublic, userId, articleRequest.TopicID)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(a, "", "\t")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (ac *ArticleController) PatchRequest(w http.ResponseWriter, r *http.Request)  {
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
		articleRequest.Tags,
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
