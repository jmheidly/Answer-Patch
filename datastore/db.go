package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDatabase(db_user, db_password, db_name string) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", db_user, db_password, db_name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func initializeDatabase(db *sql.DB) {

	// Run 'create extension "uuid-ossp";' in your psql shell in order to use the uuid_generate_v4() function

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS ap_user(id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), username varchar(20), hashed_password char(60), reputation integer DEFAULT 0, created_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS question (id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), user_id uuid REFERENCES ap_user, title varchar(255) NOT NULL UNIQUE, content text NOT NULL, upvotes integer DEFAULT 0, edit_count integer DEFAULT 0, submitted_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS answer (id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), question_id uuid REFERENCES question ON DELETE CASCADE, user_id uuid REFERENCES ap_user, is_current_answer boolean DEFAULT false, content text, upvotes integer DEFAULT 0, last_edited_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}
}

//Populates DB with questions, answers, and users for unit testing the post store
func transact(db *sql.DB, fn func(*sql.Tx) error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	if err = fn(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func checkExistence(stmt *sql.Stmt, val string) string {

	var id string

	row, err := stmt.Query(val)

	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		return ""
	}

	err = row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func standardizeTime(x *time.Time, y *time.Time) {
	*x = *y
}
