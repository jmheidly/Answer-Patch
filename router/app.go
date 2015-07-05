package router

import "github.com/gorilla/mux"

const (
	ReadUser   = "get:user"
	CreateUser = "post:user"
)

func NewAppRouter() *mux.Router {

	appRouter := mux.NewRouter().StrictSlash(false)

	//GET
	appRouter.Path("/user/{id:[A-Za-z0-9]+}").Methods("GET").Name(ReadUser)

	//POST
	appRouter.Path("/register").Methods("POST").Name(CreateUser)

	return appRouter
}
