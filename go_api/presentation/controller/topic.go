package controller

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog_go_api/domain/topic"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
	"github.com/nozomi-iida/nozo_blog_go_api/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog_go_api/service"
)

type TopicController struct {
	ts *service.TopicService
}

func NewTopicController(fileString string) (TopicController, error) {
	ts, err := service.NewTopicService(
		service.WithSqliteTopicRepository(fileString),
	)

	if err != nil {
		return TopicController{}, nil
	}

	return TopicController{ts: ts}, nil
}

type topicResponse struct {
	TopicId     uuid.UUID `json:"topicId" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description,omitempty"`
}

type topicListResponse struct {
	Topics []topicResponse `json:"topics"`
}

func topicListToJson(topics []entity.Topic) topicListResponse {
	var trs = []topicResponse{}
	for _, t := range topics {
		trs = append(trs, topicResponse{TopicId: t.TopicID, Name: t.Name, Description: t.Description})
	}

	return topicListResponse{Topics: trs}
}

func (tc *TopicController) ListRequest(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	query := topic.TopicQuery{Keyword: keyword}
	associatedWith := r.URL.Query().Get("associatedWith")
	if associatedWith != "" {
		aw, err := topic.NewAssociatedType(associatedWith)
		if err != nil {
			helpers.ErrorHandler(w, err)
			return
		}
		query.AssociatedWith = aw
	}

	topics, err := tc.ts.PublicList(query)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(topics, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (tc *TopicController) FindByNameRequest(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/topics")
	_, topicName := filepath.Split(sub)

	query := topic.PublicFindByNameQuery{}
	associatedWith := r.URL.Query().Get("associatedWith")
	if associatedWith != "" {
		na, err := topic.NewAssociatedType(associatedWith)
		if err != nil {
			helpers.ErrorHandler(w, err)
			return
		}
		query.AssociatedWith = na
	}

	topics, err := tc.ts.PublicFindByName(topicName, query)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(topics, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
