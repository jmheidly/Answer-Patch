package datastore

import (
	"database/sql"
	"log"
	"testing"
)

func TestConnectToPostgres(t *testing.T) {
	if err := ConnectToPostgres().Ping(); err != nil {
		t.Error(err)
	}
}

//Populates DB with questions, answers, and users for unit testing
func populatePostgres(db *sql.DB) {

	var err error

	//Users
	if _, err = db.Exec(`INSERT INTO ap_user(id, username, hashed_password, reputation) VALUES('{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'Tester1', '$2a$10$UyVxgEPxf.cS4V7QzuGfcOUm7mxBP8J.Rp6zqbZyppjiP8UvbU57a',10)`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO ap_user(id, username, hashed_password, reputation) VALUES ('{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Tester2', '$2a$10$Jx7qNvdifqyda30n8SetceMc0B0kWpvKKqx7GxEBR4ptKJeaVn82i', 15)`); err != nil {
		log.Fatal(err)
	}

	//Categories
	if _, err = db.Exec(`INSERT INTO category(category_name, user_id) VALUES ('Gains', '{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid)`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO category(category_name, user_id) VALUES ('City Dining', '{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid)`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO category(category_name, user_id) VALUES ('Balling', '{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid)`); err != nil {
		log.Fatal(err)
	}

	//Questions
	if _, err = db.Exec(`INSERT INTO question(id, user_id, title, content, upvotes, edit_count, category_name) VALUES('{38681976-4d2d-4581-8a68-1e4acfadcfa0}'::uuid,'{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'What is should my squat to bench ratio be?', 'I need gains', 13, 4, 'Gains')`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO question(id, user_id, title, content, upvotes, edit_count, category_name) VALUES('{526c4576-0e49-4e90-b760-e6976c698574}'::uuid,'{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Where is the best sushi place?', 'I have cravings', 15, 5, 'City Dining')`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO question(id, user_id, title, content, upvotes, edit_count, category_name) VALUES('{0a24c4cd-4c73-42e4-bcca-3844d088de85}'::uuid,'{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Will Jordans make me a sick baller?', 'I need to improve my game', 10, 1, 'Balling')`); err != nil {
		log.Fatal(err)
	}

	//Answers
	if _, err = db.Exec(`INSERT INTO answer(id, question_id, user_id, is_current_answer, content, upvotes) VALUES ('{f46fd5c9-ea9b-4677-ba8a-433b27fc097c}'::uuid, '{38681976-4d2d-4581-8a68-1e4acfadcfa0}'::uuid, '{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'true', 'Always to never', 20)`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO answer(id, question_id, user_id, is_current_answer, content, upvotes) VALUES ('{c6f753ea-8b55-468f-9eb2-3ac03f6ed179}'::uuid, '{526c4576-0e49-4e90-b760-e6976c698574}'::uuid,'{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'true', 'Not Massachusetts', 40)`); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(`INSERT INTO answer(id, question_id, user_id, is_current_answer, content, upvotes) VALUES ('{b50f0224-3fda-435b-a8a6-8257fcbf5aa7}'::uuid, '{0a24c4cd-4c73-42e4-bcca-3844d088de85}'::uuid,'{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'true', 'Yeah get the ones with the neon laces', 50)`); err != nil {
		log.Fatal(err)
	}

}
