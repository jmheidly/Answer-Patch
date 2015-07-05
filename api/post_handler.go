package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

func ServePostByID(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		question, answer := store.FindByID(mux.Vars(r)["id"])
		if question == nil {
			http.Error(w, "No question exists with the provided id", http.StatusBadRequest)
			return
		} else {
			middleware.PrintJSON(w, question)
		}

		if answer != nil {
			middleware.PrintJSON(w, answer)
		}
	}
}

func ServeQuestionsByUser(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		questions := store.FindByAuthor(mux.Vars(r)["filterBy"], mux.Vars(r)["user"])
		if questions == nil {
			http.Error(w, "", http.StatusBadRequest)
			return
			return
		}
		middleware.PrintJSON(w, questions)
	}

}

func ServeQuestionsByFilter(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)
		questions := store.FindByFilter(routeVars["postCompoent"], routeVars["filter"], routeVars["order"], routeVars["offset"])
		if questions == nil {
			http.Error(w, "No questions match the specifications in the url", http.StatusBadRequest)
			return
		}
		middleware.PrintJSON(w, questions)
	}
}

func ServeSubmitQuestion(store datastore.PostStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var missingComponent string

		question := models.NewQuestion()

		errMessage, statusCode := middleware.ParseRequestBody(r, question)
		if statusCode != 0 {
			http.Error(w, errMessage, statusCode)
			return
		}

		switch {
		case question.Title == "":
			missingComponent = "title"
		case question.Author == "":
			missingComponent = "author"
		case question.Content == "":
			missingComponent = "content"
		}

		if missingComponent != "" {
			http.Error(w, "The question "+missingComponent+" was not recieved in the payload", http.StatusBadRequest)
			return
		}

		existingID := store.CheckQuestionExistence(question.Title)

		if existingID != "" {
			url := fmt.Sprintf("/api/posts/%d", existingID)
			http.Redirect(w, r, url, 303) // Status 303 - See Other
		} else {
			question.ID = middleware.GenerateID("que")
			store.StoreQuestion(question.ID, question.Title, question.Author, question.Content)
			w.WriteHeader(http.StatusCreated)
		}
	}
}
