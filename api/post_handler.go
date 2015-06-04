package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
)

func ServePostByID(store *datastore.PostStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		post, err := store.FindByID(mux.Vars(r)["id"])
		if err != nil {
			errorHandler(w, http.StatusBadRequest, err.Error())
			return
		}
		printJSON(w, post)
	})
}

func ServeFilteredPosts(store *datastore.PostStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		posts, err := store.FindByFilter(mux.Vars(r)["filterBy"], mux.Vars(r)["filterVal"])
		if err != nil {
			errorHandler(w, http.StatusBadRequest, err.Error())
			return
		}
		printJSON(w, posts)
	})

}
