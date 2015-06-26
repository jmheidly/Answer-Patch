package datastore

import (
	"database/sql"
	"fmt"
	"log"

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

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS question (id SERIAL PRIMARY KEY, title varchar(255) NOT NULL, author varchar(50) NOT NULL, content text NOT NULL, upvotes integer DEFAULT 0, submitted_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, edit_count integer DEFAULT 0)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS answer (id SERIAL PRIMARY KEY, question_id integer REFERENCES question ON DELETE CASCADE, is_current_answer boolean DEFAULT false, author varchar(50) NOT NULL, content text, upvotes integer DEFAULT 0, last_edited_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)

	if err != nil {
		log.Fatal(err)
	}
}

func populateDatabase(db *sql.DB) {

	db.Exec(`INSERT INTO question(title, author, content, upvotes, edit_count) values( 'Where is the best sushi place?', 'postStoreTester1', 'I have cravings', 10, 2)`)
	db.Exec(`INSERT INTO question(title, author, content, upvotes, edit_count) values( 'How many days are there in spring?', 'postStoreTester2', 'I love spring weather', 13, 4)`)
	db.Exec(`INSERT INTO question(title, author, content, upvotes, edit_count) values( 'What should my bench to squat ratio be?', 'postStoreTester1', 'I need some gains', 100, 32)`)

	db.Exec(`INSERT INTO answer(question_id, is_current_answer, author, content, upvotes) values (1, 'false','postStoreTester2', 'I think New York has a couple of good ones', 1)`)
	db.Exec(`INSERT INTO answer(question_id, is_current_answer, author, content, upvotes) values (2, 'true', 'postStoreTester1', 'Spring 2016 starts on March 20th', 20)`)
	db.Exec(`INSERT INTO answer(question_id, is_current_answer, author, content, upvotes) values (3, 'true', 'postStoreTester1', 'Easy, always to never', 100)`)

}
