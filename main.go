package main

import (
	"fmt"
	"net/http"

	"github.com/patelndipen/AP1/api"
	"github.com/patelndipen/AP1/app"
	"github.com/patelndipen/AP1/datastore"
	m "github.com/patelndipen/AP1/middleware"
	auth "github.com/patelndipen/AP1/services"
	"github.com/patelndipen/AP1/settings"
)

func main() {

	settings.SetPreproductionEnv() // Set GO_ENV to "preproduction"

	postgresConn := datastore.ConnectToPostgres()
	postStore := &datastore.PostStore{DB: postgresConn}
	userStore := &datastore.UserStore{DB: postgresConn}

	ac := auth.NewAuthContext(&datastore.JWTStore{datastore.ConnectToRedis()})
	c := &m.Context{ac, &datastore.RepStore{datastore.ConnectToMongoCol()}, nil} // may have to add nil for the model

	r := http.NewServeMux()
	r.Handle("/api/", api.Handler(c, postStore))
	r.Handle("/", app.Handler(c, userStore))

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}
