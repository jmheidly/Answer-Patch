package router

import "github.com/gorilla/mux"

const (
	ViewIndex         = "get:index"
	ViewRegisteration = "get:registration_form"
	ViewSubmission    = "get:post_submission_form"
)

func NewNonAuthenticatedRouter() *mux.Router {
	nonAuthRouter := mux.NewRouter()
	nonAuthRouter.Path("/").Methods("GET").Name(ViewIndex)
	nonAuthRouter.Path("/register").Methods("GET").Name(ViewRegisteration)

	return nonAuthRouter
}

func NewAuthenticatedRouter() *mux.Router {
	authRouter := mux.NewRouter()
	authRouter.Path("submit").Methods("GET").Name(ViewSubmission)

	return authRouter
}
