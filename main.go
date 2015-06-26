package main

import (
	"fmt"
	"net/http"

	"github.com/patelndipen/AP1/api"
	"github.com/patelndipen/AP1/datastore"
)

func main() {
	postStore := datastore.NewPostStore(datastore.ConnectToDatabase("postgres", "test", "testap1poststore"))

	m := api.Handler(postStore)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", m)
}
