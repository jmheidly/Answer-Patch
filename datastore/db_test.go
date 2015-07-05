package datastore

import (
	"testing"
)

func TestConnectToDatabase(t *testing.T) {
	db := ConnectToDatabase("postgres", "test", "ap1")

	if db == nil {
		t.Error("Expected db to be of type *sql.DB, however DB is nil")
	}
}
