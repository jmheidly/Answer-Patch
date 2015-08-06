package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
	auth "github.com/patelndipen/AP1/services"
)

type MockUserStore struct {
	User         *models.User
	IsRegistered bool
}

func (store *MockUserStore) FindUser(filter, searchVal string) *models.User {
	return store.User
}

func (store *MockUserStore) StoreUser(username, hashedpassword string) {
}

func (store *MockUserStore) IsUsernameRegistered(username string) bool {
	return store.IsRegistered
}

func TestServeFindUserWithInvalidUser(t *testing.T) {

	r, err := http.NewRequest("GET", "username/NonExistent", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	ServeFindUser(&MockUserStore{})(m.NewContext(), w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected a status code of 400, because the MockUserStore's FindUser method always returns nil, but recieved an http status code of %d", w.Code)
	} else if w.Body.String() != "No user exists with the provided information\n" {
		t.Errorf("Expected the content of the responsewriter to be \"No user exists with the provided information\", but instead the responsewriter contains %s", w.Body.String())
	}
}

func TestServeRegisterUserWithRegisterUsername(t *testing.T) {

	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	registeredUser := &models.UnauthUser{Username: "RegisteredUsername"}
	c := &m.Context{ParsedModel: registeredUser}
	ServeRegisterUser(&MockUserStore{IsRegistered: true})(c, w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected a status code of 400, because the MockUserStore's IsUsernameRegistered returned true, but recieved an http status code of %d", w.Code)
	} else if w.Body.String() != "Username already exists\n" {
		t.Errorf("Expected the content of the responsewriter to be \"Username already exists\", but instead the responsewriter contains %s", w.Body.String())
	}
}

func TestServeLoginWithFailedLogin(t *testing.T) {

	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	ac := &auth.AuthContext{UserID: "ID"}
	unauthUser := &models.UnauthUser{Username: "Username", Password: "Wrong Password"}
	c := &m.Context{ac, nil, unauthUser}

	ServeLogin(&MockUserStore{User: &models.User{HashedPassword: "Hash"}})(c, w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected a status code of 401, because the MockAuthContext's Login method always returns nil, but recieved an http status code of %d", w.Code)
	} else if w.Body.String() != "Credentials are incorrect\n" {
		t.Errorf("Expected the content of the responsewriter to be \"Credentials are incorrect\", but instead the responsewriter contains %s", w.Body.String())
	}
}
