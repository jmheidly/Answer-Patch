package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/router"
)

func Handler(store *datastore.PostStore) *mux.Router {
	m := router.NewAPIRouter()

	m.Get(router.ViewPost).Handler(servePostByID(store))
	m.Get(router.ViewQuestionsByAuthor).Handler(serveQuestionsByAuthor(store))
	m.Get(router.ViewFilteredQuestions).Handler(serveQuestionsByFilter(store))

	return m
}

func errorHandler(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	fmt.Fprint(w, err)
}

func printJSON(w http.ResponseWriter, content interface{}) {

	w.Header().Set("Content-Type", "application/json")

	postJSON, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "")
		return
	}

	_, err = w.Write(postJSON)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "")
	}
}
