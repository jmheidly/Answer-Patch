package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

type MockUserStore struct {
}

func (store *MockUserStore) FindUserByID(id string) *models.User {
	return nil
}

func (store *MockUserStore) StoreUser(id, username, hashedpassword) {
}

func (store *MockUserStore) CheckUserExistence(string) string {
	return ""
}

func TestPostUser(t *testing.T) {

	store := &MockUserStore{}
	handler := middleware.CheckRequestBody(ServeSubmitUser(store))

	// Test whether a request body containing an existing username results in a http status of 400 - Bad Request
	existingUser := &models.User{ID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Username: "Where is the best sushi place?", HashedPassword: "password"}

	body, err := json.Marshal(existingUser)
	if err != nil {
		t.Error(err)
	}

	req, err = http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Expected a status code of 400 due to the existence of a user with the same username as that of the user recieved in the request body, recieved a status code of %d", w.Code)

	}
}
