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

func initializeDatabase(db *sql.DB) {

	// Run 'create extension "uuid-ossp";' in your psql shell in order for uuid_generate_v4() to be a valid value for columns

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS ap_user(id uuid PRIMARY KEY DEFAULT uuid_generate_v4(), username varchar(20) NOT NULL UNIQUE, hashed_password char(60) NOT NULL UNIQUE, reputation integer DEFAULT 0, created_at TIME WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)`)
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

//Populates DB with questions, answers, and users for unit testing
func populatePostgres(db *sql.DB) {

	//Users
	if _, err := db.Exec(`INSERT INTO ap_user(id, username, hashed_password, reputation) VALUES('{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'Tester1', '$2a$10$UyVxgEPxf.cS4V7QzuGfcOUm7mxBP8J.Rp6zqbZyppjiP8UvbU57a',10)`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO ap_user(id, username, hashed_password, reputation) VALUES ('{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Tester2', '$2a$10$Jx7qNvdifqyda30n8SetceMc0B0kWpvKKqx7GxEBR4ptKJeaVn82i', 15)`); err != nil {
		log.Fatal(err)
	}

	//Questions
	if _, err := db.Exec(`INSERT INTO question(id, user_id, title, content, upvotes, edit_count) VALUES('{38681976-4d2d-4581-8a68-1e4acfadcfa0}'::uuid,'{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'What is should my squat to bench ratio be?', 'I need gains', 13, 4)`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO question(id, user_id, title, content, upvotes, edit_count) VALUES('{526c4576-0e49-4e90-b760-e6976c698574}'::uuid,'{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Where is the best sushi place?', 'I have cravings', 15, 5)`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO question(id, user_id, title, content, upvotes, edit_count) VALUES('{0a24c4cd-4c73-42e4-bcca-3844d088de85}'::uuid,'{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Will Jordans make me a sick baller?', 'I need to improve my game', 10, 1)`); err != nil {
		log.Fatal(err)
	}

	//Answers
	if _, err := db.Exec("INSERT INTO answer(id, question_id, user_id, is_current_answer, content, upvotes) VALUES ('{f46fd5c9-ea9b-4677-ba8a-433b27fc097c}'::uuid, '{38681976-4d2d-4581-8a68-1e4acfadcfa0}'::uuid, '{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'true', 'Always to never', 20)"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("INSERT INTO answer(id, question_id, user_id, is_current_answer, content, upvotes) VALUES ('{c6f753ea-8b55-468f-9eb2-3ac03f6ed179}'::uuid, '{526c4576-0e49-4e90-b760-e6976c698574}'::uuid,'{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'true', 'Not Massachusetts', 40)"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("INSERT INTO answer(id, question_id, user_id, is_current_answer, content, upvotes) VALUES ('{b50f0224-3fda-435b-a8a6-8257fcbf5aa7}'::uuid, '{0a24c4cd-4c73-42e4-bcca-3844d088de85}'::uuid,'{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'true', 'Yeah get the ones with the neon laces', 50)"); err != nil {
		log.Fatal(err)
	}

}
