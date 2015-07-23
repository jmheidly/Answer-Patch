package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServePostByIDWithInvalidID(t *testing.T) {

	//Creates a request with an invalid ID
	r, err := http.NewRequest("GET", "api/posts/5975ea52-2a91-483f-b52f-2b0257886773", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	ServePostByID(new(MockPostStore))(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected a status code of 400, because the MockPostStore's FindPostByID method always returns nil as a result, but recieved an http status code of %d", w.Code)
	} else if w.Body.String() != "No question exists with the provided id\n" {
		t.Errorf("Expected the content of the responsewriter to be \"No question exists with the provided id\", but instead the responsewriter contains %s", w.Body.String())
	}
}

func TestServeQuestionsByAuthorWithInvalidAuthor(t *testing.T) {

	//Creates a request with an invalid Author
	r, err := http.NewRequest("GET", "api/questions/posted-by/NonExistent", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	ServeQuestionsByAuthor(new(MockPostStore))(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected a status code of 400, because the MockPostStore's FindQuestionsByAuthor method always returns nil as a result, but recieved an http status code of %d", w.Code)
	} else if w.Body.String() != "No question(s) found with the provided query\n" {
		t.Errorf("Expected the content of the responsewriter to be \"No question(s) found with the provided query\", but instead the responsewriter contains %s", w.Body.String())
	}
}

func TestServeQuestionsByFilter(t *testing.T) {

	//Creates a request with filters that no questions satisfy
	r, err := http.NewRequest("GET", "api/questions/upvotes/desc/10", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	ServeQuestionsByFilter(new(MockPostStore))(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected a status code of 400, because the MockPostStore's FindQuestionsByFilter method always returns nil as a result, but recieved an http status code of %d", w.Code)
	} else if w.Body.String() != "No questions match the specifications in the url\n" {
		t.Errorf("Expected the content of the responsewriter to be \"No questions match the specifications in the url\", but instead the responsewriter contains %s", w.Body.String())
	}
}
