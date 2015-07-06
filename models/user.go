package models

import (
	"time"
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
