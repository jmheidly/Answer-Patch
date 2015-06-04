package datastore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "dbname=ap1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func InitializeDatabase(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS post
	(post_id SERIAL PRIMARY KEY,
	 question_author varchar(50) NOT NULL,
	 question_title varchar(255) NOT NULL,
	 question_content text NOT NULL,
	 question_upvotes integer DEFAULT 0,
	 question_submitted_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, answer_last_edited_at TIME WITH TIME ZONE,
	 answer_edits integer DEFAULT 0,
	 current_answer text,
	 current_answer_upvotes integer DEFAULT 0,
	 current_answer_author varchar(50) NOT NULL
        )`)

	if err != nil {
		log.Fatal(err)
		return
	}
}
