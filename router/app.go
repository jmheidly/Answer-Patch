package router

import "github.com/gorilla/mux"

const (
	ReadUser   = "get:user"
	CreateUser = "post:user"
)

func NewAppRouter() *mux.Router {

	appRouter := mux.NewRouter().StrictSlash(false)

	//GET
	appRouter.Path("/user/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").Methods("GET").Name(ReadUser)

	//POST
	appRouter.Path("/register").Methods("POST").Name(CreateUser)

	return appRouter
}
