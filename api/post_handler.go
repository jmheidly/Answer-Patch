package api

import (
	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"net/http"
)

func servePostByID(store *datastore.PostStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		question, answer := store.FindByID(mux.Vars(r)["id"])
		if question == nil {
			errorHandler(w, http.StatusBadRequest, "No question exists with the provided id")
			return
		} else {
			printJSON(w, question)
		}

		if answer != nil {
			printJSON(w, answer)
		}
	})
}

func serveQuestionsByAuthor(store *datastore.PostStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		questions := store.FindByAuthor(mux.Vars(r)["filterBy"], mux.Vars(r)["author"])
		if questions == nil {
			errorHandler(w, http.StatusBadRequest, "")
			return
		}
		printJSON(w, questions)
	})

}

func serveQuestionsByFilter(store *datastore.PostStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		questions := store.FindByFilter(mux.Vars(r)["filter"], mux.Vars(r)["offset"])
		if questions == nil {
			errorHandler(w, http.StatusBadRequest, "No questions exist with the provided query")
			return
		}
		printJSON(w, questions)
	})
}
