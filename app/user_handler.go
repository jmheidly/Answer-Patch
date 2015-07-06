package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

func ServeUserByID(store datastore.UserStoreServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := store.FindUserByID(mux.Vars(r)["id"])
		if user == nil {
			http.Error(w, "No user exists with the provided id", http.StatusBadRequest)
		} else {
			middleware.PrintJSON(w, user)
		}
	}
}

func ServePostUser(store datastore.UserStoreServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var missingComponent string

		user := models.NewUser()

		errMessage, statusCode := middleware.ParseRequestBody(r, user)
		if statusCode != 0 {
			http.Error(w, errMessage, statusCode)
			return
		}

		switch {
		case user.Username == "":
			missingComponent = "username"
		case user.HashedPassword == "":
			missingComponent = "password"
		}

		if missingComponent != "" {
			http.Error(w, "The user's "+missingComponent+" was not recieved in the payload", http.StatusBadRequest)
			return
		}

		ExistingID := store.CheckUserExistence(user.Username)

		if ExistingID != "" {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		store.StoreUser(user.ID, user.Username, models.HashPassword(user.HashedPassword))

	}
}
