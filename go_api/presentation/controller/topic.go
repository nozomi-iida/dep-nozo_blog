package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (tc *TopicController) CreteRequest(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var topicRequest topicRequest
	json.Unmarshal(body, &topicRequest)
	fmt.Printf("rq: %v \n", topicRequest.Name)
	helpers.Validate(w, topicRequest)

	t, err := tc.ts.Create(topicRequest.Name, entity.SetDescription(topicRequest.Description))
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(t, "", "\t")

	w.WriteHeader(http.StatusCreated)
	w.Write(output)
}
