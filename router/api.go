package router

import (
	"github.com/gorilla/mux"
)

const (
	ViewPost          = "get:post"
	ViewFilteredPosts = "get:filtered_posts"
	ViewUserProfile   = "get:user_profile"
	ViewPendingEdits  = "get:pending_edit"
)

func NewAPIRouter() *mux.Router {
	apiRouter := mux.NewRouter().StrictSlash(false)
	api := apiRouter.PathPrefix("/api").Subrouter()
	api.Path("/posts/{id:[0-9]+}").Methods("GET").Name(ViewPost)
	api.Path("/posts/{filterBy:[a-z]+\\/[a-z]+}/{filterVal}").Methods("GET").Name(ViewFilteredPosts)

	//	api.Path("/{sortBy:[a-z]+\\/[a-z]+}/{sortVal}").Methods("GET").Name(ViewFilteredPosts)
	/*
		apiRouter.Path("/api/user/{usr_id:/\\w+/}").Methods("GET").Name(ViewUserProfile)
			apiRouter.Path("/api/post/{post_id:/[0-9]+/}/?edit={edit_id:/[0-9]+/}").Methods("GET").Path(ViewPendingEdits)
	*/
	return apiRouter
}
