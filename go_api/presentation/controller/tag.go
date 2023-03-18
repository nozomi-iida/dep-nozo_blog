package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/tag"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/service"
)

type TagController struct {
	ts *service.TagService
}

func NewTagController(fileString string) (TagController, error) {
	ts, err := service.NewTagService(
		service.WithSqliteTagRepository(fileString),
	)

	if err != nil {
		return TagController{}, nil
	}

	return TagController{ts: ts}, nil
}

type tagResponse struct {
	TagId uuid.UUID `json:"tagId" validate:"required"`
	Name  string    `json:"name" validate:"required"`
}

type tagListResponse struct {
	Tags []tagResponse `json:"tags"`
}

func tagListToJson(tags []entity.Tag) tagListResponse {
	var trs = []tagResponse{}
	for _, t := range tags {
		trs = append(trs, tagResponse{TagId: t.TagID, Name: t.Name})
	}

	return tagListResponse{Tags: trs}
}

func (tc *TagController) ListRequest(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	query := tag.TagQuery{Keyword: keyword}
	tags, err := tc.ts.List(query)
	if err != nil {
		helpers.ErrorHandler(w, err)
		return
	}
	tgj := tagListToJson(tags)
	output, _ := json.MarshalIndent(tgj, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
