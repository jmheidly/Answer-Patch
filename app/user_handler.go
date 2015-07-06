package app

import (
	"net/http"

	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

func ServeUserByID(store models.UserStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := store.FindByID(mux.Vars(r)["id"])
		if user == nil {
			http.Error(w, "No user exists with the provided id", http.StatusBadRequest)
		} else {
			middleware.PrintJSON(w, user)
		}
	}
}

func ServePostUser(store models.UserStoreActions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
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

		stmt, err := store.DB.Prepare(`SELECT id FROM user WHERE username=$1`)
		if err != nil {
			http.Error(w, "", http.StatusInternalError)
		}

		ExistingID := datastore.CheckExistence(stmt, username)
		
		if ExistingID != 0 {
			http.Error(w, "Username exists", http.StatusBadRequest)
			return
		}

		user.HashedPassword, err = model.HashPassword(&user.HashedPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequiest)
			return
		}
		user.ID = middleware.GenerateID("usr")

		store.StoreUser(user.ID, user.Username, user.HashedPassword)

}
