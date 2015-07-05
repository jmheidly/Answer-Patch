package api

import (
	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/router"
)

func Handler(store *datastore.PostStore) *mux.Router {
	m := router.NewAPIRouter()

	m.Get(router.ReadPost).Handler(ServePostByID(store))
	m.Get(router.ReadQuestionsByAuthor).Handler(ServeQuestionsByUser(store))
	m.Get(router.ReadFilteredQuestions).Handler(ServeQuestionsByFilter(store))
	m.Get(router.CreateQuestion).Handler(middleware.CheckRequestBody(ServeSubmitQuestion(store)))

	return m
}
