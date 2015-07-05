package datastore

import (
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/patelndipen/AP1/models"
)

var GlobalPostStore *PostStore

func init() {
	GlobalPostStore = NewPostStore(ConnectToDatabase("postgres", "test", "ap1"))
	//	initializeDatabase(GlobalPostStore.DB)
	//	populateDatabase(GlobalPostStore.DB)
}

func populateDatabase(db *sql.DB) {

	//Users
	if _, err := db.Exec(`INSERT INTO ap_user(id, username) VALUES('{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}'::uuid, 'Tester1')`); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO ap_user(id, username) VALUES ('{95954f28-a8c3-4e76-8c80-18de07931639}'::uuid, 'Tester2')`); err != nil {
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

func TestFindByID(t *testing.T) {

	expectedQuestion := &models.Question{ID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", UserID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Username: "Tester1", Title: "What is should my squat to bench ratio be?", Content: "I need gains", Upvotes: 13, EditCount: 4}

	expectedAnswer := &models.Answer{ID: "f46fd5c9-ea9b-4677-ba8a-433b27fc097c", QuestionID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", IsCurrentAnswer: true, Content: "Always to never", Upvotes: 20}

	retrievedQuestion, retrievedAnswer := GlobalPostStore.FindByID(expectedQuestion.ID)
	if retrievedQuestion == nil {
		t.Error("Expected and did not recieve %#v", expectedQuestion)
	} else if retrievedAnswer == nil {

		t.Error("Expected and did not recieve %#v", expectedAnswer)
	}

	// Avoids the complication of parsing postgres timestamp values to golang time.Time values by setting the time.Time values of expected post components equal to the time.Time values of retrieved post components
	standardizeTime(&expectedQuestion.SubmittedAt, &retrievedQuestion.SubmittedAt)
	standardizeTime(&expectedAnswer.LastEditedAt, &retrievedAnswer.LastEditedAt)

	if !reflect.DeepEqual(retrievedQuestion, expectedQuestion) {
		t.Error("Expected %#v, but recieved %#v", expectedQuestion, retrievedQuestion)
	} else if !reflect.DeepEqual(expectedAnswer, retrievedAnswer) {
		t.Error("Expected %#v, but recieved %#v", expectedAnswer, retrievedAnswer)
	}

}

func TestFindByUser(t *testing.T) {

	expectedQuestions := []*models.Question{&models.Question{ID: "526c4576-0e49-4e90-b760-e6976c698574", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", Title: "Where is the best sushi place?", Content: "I have cravings", Upvotes: 15, EditCount: 5}, &models.Question{ID: "0a24c4cd-4c73-42e4-bcca-3844d088de85", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", Title: "Will Jordans make me a sick baller?", Content: "I need to improve my game", Upvotes: 10, EditCount: 1}}

	retrievedQuestions := GlobalPostStore.FindByUser("posted-by", "Tester2")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

	retrievedQuestions = GlobalPostStore.FindByUser("answered-by", "Tester1")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

}

func TestFindByFilter(t *testing.T) {

	//Test postComponent: "question", filter: "upvotes", order: "desc"
	expectedQuestions := []*models.Question{&models.Question{ID: "526c4576-0e49-4e90-b760-e6976c698574", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", Title: "Where is the best sushi place?", Content: "I have cravings", Upvotes: 15, EditCount: 5}, &models.Question{ID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", UserID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Username: "Tester1", Title: "What is should my squat to bench ratio be?", Content: "I need gains", Upvotes: 13, EditCount: 4}, &models.Question{ID: "0a24c4cd-4c73-42e4-bcca-3844d088de85", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", Title: "Will Jordans make me a sick baller?", Content: "I need to improve my game", Upvotes: 10, EditCount: 1}}

	retrievedQuestions := GlobalPostStore.FindByFilter("question", "upvotes", "DESC", "0")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

	//Test postComponent: "answer", filter: "date", order: "asc"
	expectedQuestions = []*models.Question{&models.Question{ID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", UserID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Username: "Tester1", Title: "What is should my squat to bench ratio be?", Content: "I need gains", Upvotes: 13, EditCount: 4}, &models.Question{ID: "526c4576-0e49-4e90-b760-e6976c698574", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", Title: "Where is the best sushi place?", Content: "I have cravings", Upvotes: 15, EditCount: 5}, &models.Question{ID: "0a24c4cd-4c73-42e4-bcca-3844d088de85", UserID: "95954f28-a8c3-4e76-8c80-18de07931639", Username: "Tester2", Title: "Will Jordans make me a sick baller?", Content: "I need to improve my game", Upvotes: 10, EditCount: 1}}

	retrievedQuestions = GlobalPostStore.FindByFilter("answer", "date", "ASC", "0")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)
}

func TestCheckQuestionExistence(t *testing.T) {

	//Test for the recognition of an existing question with a matching title
	if id := GlobalPostStore.CheckQuestionExistence("Where is the best sushi place?"); id != "526c4576-0e49-4e90-b760-e6976c698574" {
		t.Errorf("CreateQuestion did not return the id of an existing question, instead CreateQuestion returned %d", id)
	}

}

func TestStoreQuestion(t *testing.T) {

	GlobalPostStore.StoreQuestion("{0c1b2b91-9164-4d52-87b0-9c4b444ee62d}", "Test title", "Content and stuff")

	row, err := GlobalPostStore.DB.Query(`SELECT title FROM question WHERE title = 'Test title'`)
	if err != nil {
		t.Error(err)
	} else if !row.Next() {
		t.Errorf("Failed to insert question into the database through PostStore's StoreQuestion function")
	}
}
func checkQuestionsForEquality(t *testing.T, x []*models.Question, y []*models.Question) {

	if x == nil {
		t.Errorf("Recieved nil, but expected %#v", y)
	}

	for i, _ := range x {
		standardizeTime(&y[i].SubmittedAt, &x[i].SubmittedAt)
		if !reflect.DeepEqual(x[i], y[i]) {
			t.Errorf("\n\nExpected %#v,\n but recieved %#v\n\n", y[i], x[i])
		}
	}
}

func standardizeTime(x *time.Time, y *time.Time) {
	*x = *y
}
