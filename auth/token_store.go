package auth

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type TokenStoreServices interface {
	StoreToken(string, string, int)
	IsTokenStored(string) bool
}

type JWTStore struct {
	Conn redis.Conn
}

func (store *JWTStore) StoreToken(userID, signedToken string, exp int) {
	_, err := store.Conn.Do("SET", userID, signedToken)
	if err != nil {
		log.Fatal(err)
	}

	_, err = store.Conn.Do("EXPIRE", userID, exp)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *JWTStore) IsTokenStored(userID string) bool {

	val, err := store.Conn.Do("GET", userID)
	if err != nil {
		log.Fatal(err)
	} else if val == nil {
		return false
	}

	return true
}
