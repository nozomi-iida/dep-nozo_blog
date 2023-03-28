package admincontroller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	adminservice "github.com/nozomi-iida/nozo_blog/service/admin-service"
)

type TopicController struct {
	ts *adminservice.TopicService
}

type topicRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

func (tr *topicRequest) FromJson(data []byte) error {
	return json.Unmarshal(data, tr)
}

func NewTopicController(fileString string) (TopicController, error) {
	ts, err := adminservice.NewTopicService(
		adminservice.WithSqliteTopicRepository(fileString),
	)
	if err != nil {
		return TopicController{}, err
	}
	return TopicController{ts: ts}, nil
}

func (tc *TopicController) ListRequest(w http.ResponseWriter, r *http.Request) {
	topics, err := tc.ts.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, _ := json.MarshalIndent(topics, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (tc *TopicController) PostRequest(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var topicRequest topicRequest
	if err := topicRequest.FromJson(body); err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	helpers.Validate(w, topicRequest)

	err := tc.ts.Create(topicRequest.Name, topicRequest.Description)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (tc *TopicController) PatchRequest(w http.ResponseWriter, r *http.Request) {
	topicId, err := uuid.Parse(r.URL.Query().Get("topic_id"))
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var topicRequest topicRequest
	if err := topicRequest.FromJson(body); err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	helpers.Validate(w, topicRequest)

	err = tc.ts.Update(topicId, topicRequest.Name, topicRequest.Description)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
