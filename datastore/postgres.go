package datastore

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/patelndipen/AP1/settings"
)

func ConnectToPostgres() *sql.DB {

	dns := settings.GetPostgresDSN()

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

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

func standardizeTime(x *time.Time, y *time.Time) {
	*x = *y
}

func initializePostgres(db *sql.DB) {

	// Create Database

	_, err := db.Exec(`CREATE DATABASE ap1`)
	if err != nil {
		log.Fatal(err)
	}

	// Run 'create extension "uuid-ossp";' in your psql shell in order for uuid_generate_v4() to be a valid value for columns

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ap_user(id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), username varchar(20) NOT NULL UNIQUE, hashed_password char(60) NOT NULL UNIQUE, reputation integer DEFAULT 5, created_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS category (category_name varchar(15) PRIMARY KEY, user_id uuid REFERENCES ap_user NOT NULL, created_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS question (id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), user_id uuid REFERENCES ap_user NOT NULL, title varchar(255) NOT NULL UNIQUE, content text NOT NULL, upvotes integer DEFAULT 0, edit_count integer DEFAULT 0, category_name varchar(15) REFERENCES category NOT NULL, submitted_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS answer (id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), question_id uuid REFERENCES question ON DELETE CASCADE NOT NULL, user_id uuid REFERENCES ap_user NOT NULL, is_current_answer boolean DEFAULT false, content text, upvotes integer DEFAULT 0, last_edited_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}

}
