package datastore

import (
	"testing"
)

func TestConnectToPostgres(t *testing.T) {
	if err := ConnectToPostgres().Ping(); err != nil {
		t.Error(err)
	}
}
