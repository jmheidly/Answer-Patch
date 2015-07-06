package app

import (
	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/router"
)

func Handler(store *datastore.UserStore) *mux.Router {

	m := router.NewAppRouter()

	m.Get(router.ReadUser).Handler(ServeUserByID(store))
	m.Get(router.CreateUser).Handler(middleware.CheckRequestBody(ServePostUser(store)))
	return m
}
