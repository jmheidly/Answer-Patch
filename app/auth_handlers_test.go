package app

import (
	//	"testing"

	auth "github.com/patelndipen/AP1/services"
)

type MockAuthContext struct {
}

func (ac *MockAuthContext) Login(enteredpassword, hashedpassword string) *auth.Token {
	return nil
}

func (ac *MockAuthContext) Logout(signedToken string) {
}

func (ac *MockAuthContext) RefreshToken() *auth.Token {
	return nil
}
