package models

import (
	"time"

	"github.com/patelndipen/AP1/auth"
)

type Context struct {
	*auth.AuthContext
	ParsedModel ModelServices
}

func NewContext(store auth.TokenStoreServices) *Context {

	return &Context{&auth.AuthContext{UserID: "", Exp: time.Time{}, TokenStore: store}, nil}

}
