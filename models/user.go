package models

type User struct {
	ID             int
	Username       string
	HashedPassword string
	Reputation     int
}

func NewUser() *User {
	return &User{}
}

/*
func CreateNewUser(
*/
