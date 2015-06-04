package datastore

import "errors"

var (
	errInvalidRequestParam = errors.New("No post(s) exists with the provided information")
	errInvalidFilter       = errors.New("The requested filter does not exist")
)
