package datastore

import (
	"reflect"
	"testing"

	"github.com/patelndipen/AP1/models"
)

var GlobalUserStore *UserStore

func init() {

	GlobalUserStore = NewUserStore(ConnectToDatabase("postgres", "test", "ap1"))

}

func TestFindUserByID(t *testing.T) {

	expectedUser := &models.User{ID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Username: "Tester1", HashedPassword: "$2a$10$UyVxgEPxf.cS4V7QzuGfcOUm7mxBP8J.Rp6zqbZyppjiP8UvbU57a", Reputation: 10}

	retrievedUser := GlobalUserStore.FindUserByID("0c1b2b91-9164-4d52-87b0-9c4b444ee62d")

	if retrievedUser == nil {
		t.Errorf("Expected and did not recieve %#v", expectedUser)
	} else {
		// Avoids the complication of parsing postgres timestamp values to golang time.Time values by setting the time.Time values of expected post components equal to the time.Time values of retrieved post components
		standardizeTime(&expectedUser.CreatedAt, &retrievedUser.CreatedAt)
	}

	if !reflect.DeepEqual(retrievedUser, expectedUser) {
		t.Errorf("Expected %#v, but recieved %#v", expectedUser, retrievedUser)
	}

}

func TestStoreUser(t *testing.T) {
	GlobalUserStore.StoreUser("{40ca4830-fdb0-48b6-b880-8b1677624223}", "TestUser", "$2a$10$iziTEDykz1SgOVWhLuBxeeBiZFJdD6GfTO0vA06IJTafiPfSu4QYq")
	row, err := GlobalUserStore.DB.Query(`SELECT username FROM ap_user WHERE username = 'TestUser'`)
	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		t.Errorf("Failed to insert user into the database through Userstore's StoreUser function")
	}
}
