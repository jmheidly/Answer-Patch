package app

import (
	"github.com/gorilla/mux"
	"github.com/patelndipen/AP1/datastore"
	m "github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/router"
)

func Handler(c *models.Context, store *datastore.UserStore) *mux.Router {

	r := router.NewAppRouter()

	r.Get(router.ReadUser).Handler(ServeFindUser(store))
	r.Get(router.CreateUser).Handler(m.ServeHTTP(m.ParseRequestBody(new(models.UnauthUser), ServeRegisterUser(store))))
	r.Get(router.Login).Handler(m.ServeHTTP(m.ParseRequestBody(new(models.UnauthUser), ServeLogin(store))))
	r.Get(router.Logout).Handler(m.AuthenticateToken(c, ServeLogout()))

	return r
}
