package datastore

import (
	"database/sql"
	"log"

	"github.com/patelndipen/AP1/models"
)

type UserStoreServices interface {
	FindUser(string, string) *models.User
	StoreUser(string, string)
	IsUsernameRegistered(string) bool
}

type UserStore struct {
	DB *sql.DB
}

func (store *UserStore) FindUser(filter, searchVal string) *models.User {

	queryStmt := `SELECT id, username, hashed_password, reputation, created_at FROM  ap_user WHERE ` + filter + ` =$1`

	row, err := store.DB.Query(queryStmt, searchVal)
	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		return nil
	}

	user := new(models.User)

	err = row.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Reputation, &user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func (store *UserStore) StoreUser(username, hashedpassword string) {

	transact(store.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`INSERT INTO ap_user(username, hashed_password) values($1, $2)`, username, hashedpassword)
		return err
	})
}

func (store *UserStore) IsUsernameRegistered(username string) bool {

	row, err := store.DB.Query(`SELECT username FROM ap_user WHERE username = $1`, username)
	if err != nil {
		log.Fatal(err)
	}

	return row.Next()
}
