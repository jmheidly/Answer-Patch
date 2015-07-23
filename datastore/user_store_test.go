package datastore

import (
	"reflect"
	"testing"

	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/settings"
)

var GlobalUserStore *UserStore

func init() {
	settings.SetPreproductionEnv()
	GlobalUserStore = &UserStore{DB: ConnectToPostgres()}

}

func TestFindUser(t *testing.T) {

	expectedUser := &models.User{ID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Username: "Tester1", HashedPassword: "$2a$10$UyVxgEPxf.cS4V7QzuGfcOUm7mxBP8J.Rp6zqbZyppjiP8UvbU57a", Reputation: 10}

	retrievedUser := GlobalUserStore.FindUser("id", "0c1b2b91-9164-4d52-87b0-9c4b444ee62d")

	if retrievedUser == nil {
		t.Errorf("Expected and did not recieve %#v", expectedUser)
	} else {
		compareUsers(t, expectedUser, retrievedUser)
	}

	retrievedUser = GlobalUserStore.FindUser("username", "Tester1")

	if retrievedUser == nil {
		t.Errorf("Expected and did not recieve %#v", expectedUser)
	} else {
		compareUsers(t, expectedUser, retrievedUser)
	}
}

func TestIsUsernameRegistered(t *testing.T) {
	if !GlobalUserStore.IsUsernameRegistered("Tester2") {
		t.Errorf("IsUsernameUnique function failed to recognize \"Tester2\" as a non unique username")
	}
}

func TestStoreUser(t *testing.T) {
	GlobalUserStore.StoreUser("TestUser", "$2a$10$iziTEDykz1SgOVWhLuBxeeBiZFJdD6GfTO0vA06IJTafiPfSu4QYq")
	//	t.Errorf("\n%t\n", GlobalUserStore.IsUsernameRegistered("TestUser"))
	if !GlobalUserStore.IsUsernameRegistered("TestUser") {
		t.Errorf("Failed to insert user into the database through Userstore's StoreUser function")
	}
}

func compareUsers(t *testing.T, x *models.User, y *models.User) {

	// Avoids the complication of parsing postgres timestamp values to golang time.Time values by setting the time.Time values of expected post components equal to the time.Time values of retrieved post components
	standardizeTime(&x.CreatedAt, &y.CreatedAt)

	if !reflect.DeepEqual(x, y) {
		t.Errorf("Expected %#v, but recieved %#v", x, y)
	}
}
