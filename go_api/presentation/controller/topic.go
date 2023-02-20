package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
	topics, err := tc.ts.List()
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	output, _ := json.MarshalIndent(topics, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
