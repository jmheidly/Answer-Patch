package datastore

import (
	"database/sql"
	"log"

	"github.com/patelndipen/AP1/models"
)

type UserStoreServices interface {
	FindByID(string) *models.User
	StoreUser(string, string, string)
}

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

func (store UserStore) FindUserByID(id string) *models.User {

	user := models.NewUser()

	row, err := store.DB.Query(`SELECT username, hashed_password, reputation, created_at FROM ap_user WHERE id=$1`, id)
	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		return nil
	}

	err = row.Scan(&user.Username, &user.HashedPassword, &user.Reputation, &user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	user.ID = id

	return user
}

func (store UserStore) StoreUser(id, username, hashedPassword string) {

	transact(store.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`INSERT INTO user(id, username, hashed_password) values($1::uuid, $2, $3)`, id, username, hashedPassword)
		return err
	})
}
