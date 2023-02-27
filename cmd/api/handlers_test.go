package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(r *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestAuthenticate(t *testing.T) {
	jsonToReturn := `
	{
		"error": false,
		"message": "some message"
	}
	`

	client := NewTestClient(func(r *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client
	postBody := map[string]interface{}{
		"email":    "me@here.com",
		"password": "verysecret",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest(http.MethodPost, "/authenticate", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("expected http.statusAccepted but got %d", rr.Code)
	}
}
