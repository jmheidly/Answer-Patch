package router

import (
	"github.com/gorilla/mux"
)

const (
	ViewPost              = "get:post"
	ViewQuestionsByAuthor = "get:questions_by_author"
	ViewFilteredQuestions = "get:questions_by_filter"
	ViewUserProfile       = "get:user_profile"
	ViewPendingEdits      = "get:pending_edit"
)

func NewAPIRouter() *mux.Router {
	apiRouter := mux.NewRouter().StrictSlash(false)

	api := apiRouter.PathPrefix("/api").Subrouter()
	api.Path("/posts/{id:[0-9]+}").Methods("GET").Name(ViewPost)
	api.Path("/questions/{filterBy:[a-z\\-a-z]+}/{author:[A-Za-z0-9]+}").Methods("GET").Name(ViewQuestionsByAuthor)
	api.Path("/{filter:[a-z]+\\/[a-z]+\\/[a-z]+}/{offset:[0-9]+}").Methods("GET").Name(ViewFilteredQuestions)

	return apiRouter
}
