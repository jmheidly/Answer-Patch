package api

import (
	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	m "github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/router"
)

func Handler(c *m.Context, store *datastore.PostStore) *mux.Router {
	r := router.NewAPIRouter()

	r.Get(router.ReadPost).Handler(m.AuthenticateToken(c, m.RefreshExpiringToken(ServePostByID(store))))
	r.Get(router.ReadQuestionsByFilter).Handler(m.AuthenticateToken(c, m.RefreshExpiringToken(ServeQuestionsByFilter(store))))
	r.Get(router.ReadSortedQuestions).Handler(m.AuthenticateToken(c, m.RefreshExpiringToken(ServeSortedQuestions(store))))
	r.Get(router.CreateQuestion).Handler(m.AuthenticateToken(c, m.RefreshExpiringToken(m.ParseRequestBody(new(models.Question), ServeSubmitQuestion(store)))))

	return r
}
