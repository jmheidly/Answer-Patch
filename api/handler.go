package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/router"
)

func Handler(store *datastore.PostStore) *mux.Router {
	m := router.NewAPIRouter()

	m.Get(router.ReadPost).Handler(ServePostByID(store))
	m.Get(router.ReadQuestionsByAuthor).Handler(ServeQuestionsByAuthor(store))
	m.Get(router.ReadFilteredQuestions).Handler(ServeQuestionsByFilter(store))
	m.Get(router.CreateQuestion).Handler(checkRequestBody(ServeSubmitQuestion(store)))

	return m
}

// Ensures that POST requests contain data in the request body
func checkRequestBody(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No data recieved through the request", http.StatusBadRequest)
		} else {
			fn(w, r)
		}
	}
}

func printJSON(w http.ResponseWriter, content interface{}) {

	w.Header().Set("Content-Type", "application/json")

	postJSON, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(postJSON)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func parseRequestBody(w http.ResponseWriter, requestBody io.ReadCloser, loc interface{}) {

	body, err := ioutil.ReadAll(io.LimitReader(requestBody, 1048576))
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	if err = requestBody.Close(); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, loc)
	if err != nil {
		http.Error(w, err.Error()+"\n", 422) //unprocessable entity
	}
}
