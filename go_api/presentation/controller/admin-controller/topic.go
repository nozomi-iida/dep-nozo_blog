package admincontroller

import (
	"encoding/json"
	"net/http"

	adminservice "github.com/nozomi-iida/nozo_blog/service/admin-service"
)

type TopicController struct {
	ts *adminservice.TopicService
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
