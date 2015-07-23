package router

import "github.com/gorilla/mux"

const (
	ReadUser   = "get:user"
	CreateUser = "post:user"
	Login      = "post:login"
	Logout     = "post:logout"
)

func NewAppRouter() *mux.Router {

	appRouter := mux.NewRouter().StrictSlash(false)

	//GET
	appRouter.Path("/{filter:id|username}/{searchVal:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|[a-z0-9]}").Methods("GET").Name(ReadUser)

	//POST
	appRouter.Path("/register").Methods("POST").Name(CreateUser)
	appRouter.Path("/login").Methods("POST").Name(Login)
	appRouter.Path("/logout").Methods("POST").Name(Logout)

	return appRouter
}
