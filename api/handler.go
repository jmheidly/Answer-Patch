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

	m.Get(router.ViewPost).Handler(ServePostByID(store))
	m.Get(router.ViewFilteredPosts).Handler(ServeFilteredPosts(store))

	return m
}

func errorHandler(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	fmt.Fprint(w, err)
}

func printJSON(w http.ResponseWriter, posts interface{}) {

	postJSON, err := json.MarshalIndent(posts, "", " ")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(postJSON)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "")
	}
}
