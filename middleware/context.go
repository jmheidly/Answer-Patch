package middleware

import (
	"github.com/patelndipen/AP1/datastore"
	"github.com/patelndipen/AP1/models"
	auth "github.com/patelndipen/AP1/services"
)

type Context struct {
	*auth.AuthContext
	RepStore    datastore.RepStoreServices
	ParsedModel models.ModelServices
}

func NewContext() *Context {
	return &Context{new(auth.AuthContext), nil, nil}
}
