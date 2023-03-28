package admincontroller_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/presentation"
	adminservice "github.com/nozomi-iida/nozo_blog_go_api/service/admin-service"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
)

func TestTopicController_ListRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	var r, _ = presentation.NewRouter(ts.Filename)
	testServer := httptest.NewServer(r)
	t.Cleanup(func() {
		testServer.Close()
	})
	us := test.CreateUser(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, testServer.URL+"/api/v1/admin/topics", nil)
	token, err := us.UserId.Encode()
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test          string
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "get all Article",
			expectedCount: 3,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			var got adminservice.TopicListDto
			err = json.NewDecoder(resp.Body).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.expectedCount != len(got.Topics) {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(got.Topics))
			}
		})
	}
}

func TestTopicController_PostRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	var r, _ = presentation.NewRouter(ts.Filename)
	testServer := httptest.NewServer(r)
	t.Cleanup(func() {
		testServer.Close()
	})
	us := test.CreateUser(t, ts.Filename)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, testServer.URL+"/api/v1/admin/topics", nil)
	token, err := us.UserId.Encode()
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test        string
		name        string
		description string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Success to create topic",
			name:        "test",
			description: "test",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestTopicController_PatchRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	var r, _ = presentation.NewRouter(ts.Filename)
	testServer := httptest.NewServer(r)
	t.Cleanup(func() {
		testServer.Close()
	})
	us := test.CreateUser(t, ts.Filename)
	targeted := factories.CreateTopic(t, ts.Filename)
	req, err := http.NewRequestWithContext(
		context.TODO(), 
		http.MethodPatch, 
		testServer.URL+fmt.Sprintf("/api/v1/admin/topics/%s", targeted.TopicID), 
		nil,
	)
	token, err := us.UserId.Encode()
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test        string
		name        string
		description string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Success to create topic",
			name:        "test",
			description: targeted.Description,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
