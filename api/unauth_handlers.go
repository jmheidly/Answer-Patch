package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	m "github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/services"
)

func ServePostByID(store datastore.PostStoreServices) m.HandlerFunc {
	return func(c *m.Context, w http.ResponseWriter, r *http.Request) {
		question, answer := store.FindPostByID(mux.Vars(r)["id"])
		if question == nil {
			http.Error(w, "No question exists with the provided id", http.StatusBadRequest)
			return
		}

		services.PrintJSON(w, question)
		if answer != nil {
			services.PrintJSON(w, answer)
		}
	}
}

func ServeQuestionsByFilter(store datastore.PostStoreServices) m.HandlerFunc {
	return func(c *m.Context, w http.ResponseWriter, r *http.Request) {

		questions := store.FindQuestionsByFilter(mux.Vars(r)["filter"], mux.Vars(r)["val"])
		if questions == nil {
			http.Error(w, "No question(s) found with the provided query", http.StatusBadRequest)
			return
		}
		services.PrintJSON(w, questions)
	}

}

func ServeSortedQuestions(store datastore.PostStoreServices) m.HandlerFunc {
	return func(c *m.Context, w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)
		questions := store.SortQuestions(routeVars["postComponent"], routeVars["filter"], routeVars["order"], routeVars["offset"])
		if questions == nil {
			http.Error(w, "No questions match the specifications in the url", http.StatusBadRequest)
			return
		}
		services.PrintJSON(w, questions)
	}
}
