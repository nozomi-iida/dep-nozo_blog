package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/service"
)

type TopicController struct {
	ts *service.TopicService
}

func NewTopicController(fileString string) (TopicController, error)  {
	ts, err := service.NewTopicService(
		service.WithSqliteTopicRepository(fileString),
	)	

	if err != nil {
		return TopicController{}, nil
	}

	return TopicController{ts: ts}, nil
}

type topicRequest struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

type topicResponse struct {
	TopicId uuid.UUID `json:"topicId" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
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

func (tc *TopicController) CreteRequest(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var topicRequest topicRequest
	json.Unmarshal(body, &topicRequest)
	helpers.Validate(w, topicRequest)

	t, err := tc.ts.Create(topicRequest.Name, topicRequest.Description)
	tp := topicResponse{TopicId: t.TopicID, Name: t.Name, Description: t.Description}
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	output, _ := json.MarshalIndent(tp, "", "\t")

	w.WriteHeader(http.StatusCreated)
	w.Write(output)
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
