package datastore

import (
	"testing"
)

func TestConnectToDatabase(t *testing.T) {
	db := ConnectToDatabase()

	if db == nil {
		t.Error("Expected db to be of type *sql.DB, however DB is nil")
	}

}

func TestInitializeDatabase(t *testing.T) {
	db := ConnectToDatabase()
	InitializeDatabase(db)
}
