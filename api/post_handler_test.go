package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

type MockPostStore struct {
	existingID string
}

func (store *MockPostStore) FindPostByID(id string) (*models.Question, *models.Answer) {
	return nil, nil
}

func (store *MockPostStore) FindQuestionsByUser(filter, author string) []*models.Question {
	return nil
}

func (store *MockPostStore) FindQuestionsByFilter(postComponent, filter, order, offset string) []*models.Question {
	return nil
}

func (store *MockPostStore) CheckQuestionExistence(title string) string {
	return store.existingID
}

func (store *MockPostStore) StoreQuestion(user_id, title, content string) {
	return
}

func TestServeSubmitQuestion(t *testing.T) {

	store := &MockPostStore{existingID: "526c4576-0e49-4e90-b760-e6976c698574"}
	handler := middleware.CheckRequestBody(ServeSubmitQuestion(store))

	// Test whether empty request body results in a status code of 400 - Bad Request
	req, err := http.NewRequest("POST", "api/posts/question", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Expected a status code of 400 due to the absence of a request body, recieved a status code of %d", w.Code)
	}

	// Test whether a request body containing an existing question title results in a redirect to the existing question title
	existingQuestion := &models.Question{UserID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Title: "Where is the best sushi place?", Content: "I have cravings"}

	body, err := json.Marshal(existingQuestion)
	if err != nil {
		t.Error(err)
	}

	req, err = http.NewRequest("POST", "api/posts/question", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != 303 {
		t.Errorf("Expected a status code of 303 due to the existence of a question with the same title as that of the question recieved in the request body, recieved a status code of %d", w.Code)

	}
}
