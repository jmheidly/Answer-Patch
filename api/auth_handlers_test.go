package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

type MockPostStore struct {
	ExistingID string
}

func (store *MockPostStore) FindPostByID(id string) (*models.Question, *models.Answer) {
	return nil, nil
}

func (store *MockPostStore) FindQuestionsByFilter(filter, val string) []*models.Question {
	return nil
}

func (store *MockPostStore) SortQuestions(postComponent, filter, order, offset string) []*models.Question {
	return nil
}

func (store *MockPostStore) IsTitleUnique(title string) (bool, string) {
	return false, store.ExistingID
}

func (store *MockPostStore) StoreQuestion(user_id, title, content, category string) {
	return
}

func TestServeSubmitQuestionWithExistingQuestion(t *testing.T) {

	existingQuestion := &models.Question{AuthorID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Author: "Tester1", Title: "Where is the best sushi place?", Content: "I have cravings"}

	c := &m.Context{ParsedModel: existingQuestion}

	r, err := http.NewRequest("POST", "api/posts/question", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	ServeSubmitQuestion(&MockPostStore{ExistingID: "526c4576-0e49-4e90-b760-e6976c698574"})(c, w, r)

	if w.Code != http.StatusSeeOther {
		t.Errorf("Expected a status code of 303 due to the existence of a question with the same title as that of the question recieved in the request body, recieved a status code of %d", w.Code)

	}
}
