package router

import (
	"github.com/gorilla/mux"
)

const (
	ReadPost              = "get:post"
	ReadQuestionsByFilter = "get:questions_by_filter"
	ReadSortedQuestions   = "get:sorted_questions"
	CreateQuestion        = "post:question"
	CreatePendingAnswer   = "put:pending_answer"
)

func NewAPIRouter() *mux.Router {
	apiRouter := mux.NewRouter().StrictSlash(false)

	api := apiRouter.PathPrefix("/api").Subrouter()

	//GET
	api.Path("/posts/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").Methods("GET").Name(ReadPost)
	api.Path("/questions/{filter:posted-by|answered-by|category}/{val:[A-Za-z0-9]+}").Methods("GET").Name(ReadQuestionsByFilter)
	api.Path("/{postComponent:questions|answers}/{sortedBy:upvotes|edits|date}/{order:desc|asc}/{offset:[0-9]+}").Methods("GET").Name(ReadSortedQuestions)

	//POST
	api.Path("question/{category:[a-z]+}").Methods("POST").Name(CreateQuestion)

	//PUT
	api.Path("/answer/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}").Methods("PUT").Name(CreatePendingAnswer)

	return apiRouter
}
