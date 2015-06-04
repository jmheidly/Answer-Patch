package datastore

import (
	"testing"
)

func TestFind(t *testing.T) {
	postStore := NewPostStore()
	postStore.DB = ConnectToDatabase()
	InitializeDatabase(postStore.DB)

	//   Uncomment this block when you push to github
	/*
			_, err := postStore.DB.Exec(`INSERT INTO post VALUES( 1, 'firstUser', 'does this find function work?', 'I hope the find function works', 10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 2, 'Run this function to find out' , 20, 'secondUser')`)

		if err != nil {
			t.Fatal(err)
		}
	*/
	retrievedPost, err := postStore.FindByID("1")
	if err != nil {
		t.Fatal(err)
	}

	if retrievedPost.Qtitle != "does this find function work?" {
		t.Error("The post struct retrieved form the Find function of the post store does not contain the correct question title of \"does this find funciton work?\". The question title retrieved is ", retrievedPost.Qtitle)
	}
	if retrievedPost.SecondPendingAns != "" {
		t.Error("The field SecondPendingAns is supposed to equal to \"\", because the call to the scan function in the postStore's Find function should not have returned a value for SecondPendingAns due to the fact that no value for SecondPendingAns was supplied to the DB. The SecondPendingAns value of the retrived post is ", retrievedPost.SecondPendingAns)
	}

}
