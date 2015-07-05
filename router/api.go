package router

import (
	"github.com/gorilla/mux"
)

const (
	ReadPost              = "get:post"
	ReadQuestionsByAuthor = "get:questions_by_author"
	ReadFilteredQuestions = "get:questions_by_filter"
	CreateQuestion        = "post:question"
	UpdateAnswer          = "put:answer"
)

func NewAPIRouter() *mux.Router {
	apiRouter := mux.NewRouter().StrictSlash(false)

	api := apiRouter.PathPrefix("/api").Subrouter()

	//GET
	api.Path("/posts/{id:[0-9]+}").Methods("GET").Name(ReadPost)
	api.Path("/questions/{filterBy:posted-by|answered-by}/{user:[A-Za-z0-9]+}").Methods("GET").Name(ReadQuestionsByAuthor)
	api.Path("/{postComponent:question|answer}/{filter:upvotes|edits|date}/{order:DESC|ASC}/{offset:[0-9]+}").Methods("GET").Name(ReadFilteredQuestions)

	//POST
	api.Path("/posts/question").Methods("POST").Name(CreateQuestion)

	//PUT
	api.Path("/posts/{id:[0-9]+}").Methods("PUT").Name(UpdateAnswer)

	return apiRouter
}
