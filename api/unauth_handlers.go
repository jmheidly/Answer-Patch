package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/services"
)

func ServePostByID(store datastore.PostStoreServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		question, answer := store.FindPostByID(mux.Vars(r)["id"])
		if question == nil {
			http.Error(w, "No question exists with the provided id", http.StatusBadRequest)
			return
		} else if answer == nil {
			services.PrintJSON(w, question)
		} else {
			services.PrintJSON(w, answer)
		}
	}
}

func ServeQuestionsByAuthor(store datastore.PostStoreServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		questions := store.FindQuestionsByAuthor(mux.Vars(r)["filterBy"], mux.Vars(r)["author"])
		if questions == nil {
			http.Error(w, "No question(s) found with the provided query", http.StatusBadRequest)
			return
		}
		services.PrintJSON(w, questions)
	}

}

func ServeQuestionsByFilter(store datastore.PostStoreServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)
		questions := store.FindQuestionsByFilter(routeVars["postComponent"], routeVars["filter"], routeVars["order"], routeVars["offset"])
		if questions == nil {
			http.Error(w, "No questions match the specifications in the url", http.StatusBadRequest)
			return
		}
		services.PrintJSON(w, questions)
	}
}
