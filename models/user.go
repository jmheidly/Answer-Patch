package models

import (
	"time"

	"code.google.com/p/go.crypto/bcrypt"
)

const (
	hashCost = 10
)

type User struct {
	ID             string
	Username       string
	HashedPassword string
	Reputation     int
	CreatedAt      time.Time
}

func NewUser() *User {
	return &User{}
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
