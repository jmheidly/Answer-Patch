package main

import (
	"fmt"
	"net/http"

	"github.com/patelndipen/AP1/api"
	"github.com/patelndipen/AP1/datastore"
)

func main() {
	postStore := datastore.NewPostStore(datastore.ConnectToDatabase("postgres", "test", "ap1"))

	m := http.NewServeMux()
	m.Handle("/api/", api.Handler(postStore))
	//	m.Handle("/", app.Handler(postStore))

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", m)
}
