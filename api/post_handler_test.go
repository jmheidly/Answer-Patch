package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patelndipen/AP1/app"
)

type MockPostStore struct {
	existingID int
}

func (store *MockPostStore) FindByID(id string) (*app.Question, *app.Answer) {
	return nil, nil
}

func (store *MockPostStore) FindByAuthor(filter, author string) []*app.Question {
	return nil
}

func (store *MockPostStore) FindByFilter(filter, offset string) []*app.Question {
	return nil
}

func (store *MockPostStore) CreateQuestion(q *app.Question) int {
	return store.existingID
}

func TestServeSubmitQuestion(t *testing.T) {

	store := &MockPostStore{existingID: 1}
	handler := checkRequestBody(ServeSubmitQuestion(store))

	// Test whether empty request body reuslts in a status code of 400 - Bad Request
	req, err := http.NewRequest("POST", "api/posts/question", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Expected a status code of 400 due to the absence of a request body, recieved a status code of %d", w.Code)
	}

	// Test whether a request body containing an existing post results in a redirect to the existing question
	existingQuestion := &app.Question{Title: "Where is the best sushi place?", Author: "sisyphus", Content: "I have cravings"}

	body, err := json.Marshal(existingQuestion)
	if err != nil {
		t.Error(err)
	}

	req, err = http.NewRequest("POST", "api/posts/question", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != 303 {
		t.Errorf("Expected a status code of 303 due the existence of a question with the same title as that of the question recieved in the request body, recieved a status code of %d", w.Code)

	}
}
