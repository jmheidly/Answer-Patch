package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/app"
	"github.com/patelndipen/AP1/datastore"
)

func ServePostByID(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		question, answer := store.FindByID(mux.Vars(r)["id"])
		if question == nil {
			http.Error(w, "No question exists with the provided id", http.StatusBadRequest)
		} else {
			printJSON(w, question)
		}

		if answer != nil {
			printJSON(w, answer)
		}
	}
}

func ServeQuestionsByAuthor(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		questions := store.FindByAuthor(mux.Vars(r)["filterBy"], mux.Vars(r)["author"])
		if questions == nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		printJSON(w, questions)
	}

}

func ServeQuestionsByFilter(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions := store.FindByFilter(mux.Vars(r)["filter"], mux.Vars(r)["offset"])
		if questions == nil {
			http.Error(w, "No questions exist with the provided query", http.StatusBadRequest)
			return
		}
		printJSON(w, questions)
	}
}

func ServeSubmitQuestion(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		question := app.NewQuestionStruct()

		parseRequestBody(w, r.Body, question)

		val := store.CreateQuestion(question)
		if val != 0 {
			url := fmt.Sprintf("/api/posts/%d", val)
			http.Redirect(w, r, url, 303) // Status 303 - See Other
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}
