package datastore

import "errors"

var (
	errInvalidRequestParam = errors.New("No post(s) exists with the provided information")
	errInvalidAuthorRequest = errors.New("No results found from the provided user")
)
