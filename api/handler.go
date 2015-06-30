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

		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			http.Error(w, "This api only accepts JSON payloads. Be sure to specify the \"Content-Type\" of the payload in the request header.", http.StatusBadRequest)
		} else if r.Body == nil {
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

func parseRequestBody(r *http.Request, loc interface{}) (string, int) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return "Error reading request body", http.StatusInternalServerError
	}

	if err = r.Body.Close(); err != nil {
		return "", http.StatusInternalServerError
	}

	err = json.Unmarshal(body, loc)
	if err != nil {
		return err.Error() + "\n", 422 //unprocessable entity
	}

	return "", 0

}
