package app

import (
	"github.com/gorilla/mux"
)

func Handler(store *datastore.UserStore) *mux.Router {

	m := router.NewAppRouter()

	m.Get(router.ReadUser).Handler(ServeUserByID)
	m.Get(router.CreateUser).Handler(middleware.CheckRequestBody(ServePostUser))
	return m
}
