package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/services"
)

func ServeFindUser(store datastore.UserStoreServices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := store.FindUser(mux.Vars(r)["filter"], mux.Vars(r)["searchVal"])
		if user == nil {
			http.Error(w, "No user exists with the provided information", http.StatusBadRequest)
		} else {
			services.PrintJSON(w, user)
		}
	}
}

func ServeRegisterUser(store datastore.UserStoreServices) middleware.ContextRequiredHandlerFunc {
	return func(c *models.Context, w http.ResponseWriter, r *http.Request) {

		newUser := c.ParsedModel.(*models.UnauthUser)

		if store.IsUsernameRegistered(newUser.Username) {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}

		store.StoreUser(newUser.Username, newUser.HashPassword())

	}
}

func ServeLogin(store datastore.UserStoreServices) middleware.ContextRequiredHandlerFunc {
	return func(c *models.Context, w http.ResponseWriter, r *http.Request) {
		unauthUser := c.ParsedModel.(*models.UnauthUser)
		retrievedUser := store.FindUser("username", unauthUser.Username)
		if retrievedUser == nil {
			http.Error(w, "Credentials are incorrect", http.StatusUnauthorized)
			return
		}

		c.UserID = retrievedUser.ID
		token := c.Login(unauthUser.Password, retrievedUser.HashedPassword)
		if token == nil {
			http.Error(w, "Credentials are incorrect", http.StatusUnauthorized)
			return
		}

		services.PrintJSON(w, token)
	}
}
