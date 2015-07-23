package datastore

import (
	"reflect"
	"testing"

	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/settings"
)

var GlobalPostStore *PostStore

func init() {
	settings.SetPreproductionEnv()
	GlobalPostStore = &PostStore{ConnectToPostgres()}
}

func TestFindPostByID(t *testing.T) {

	expectedQuestion := &models.Question{ID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", AuthorID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Author: "Tester1", Title: "What is should my squat to bench ratio be?", Content: "I need gains", Upvotes: 13, EditCount: 4}

	expectedAnswer := &models.Answer{ID: "f46fd5c9-ea9b-4677-ba8a-433b27fc097c", QuestionID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", IsCurrentAnswer: true, Content: "Always to never", Upvotes: 20}

	retrievedQuestion, retrievedAnswer := GlobalPostStore.FindPostByID(expectedQuestion.ID)
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

func TestFindQuestionsByAuthor(t *testing.T) {

	expectedQuestions := []*models.Question{&models.Question{ID: "526c4576-0e49-4e90-b760-e6976c698574", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", Title: "Where is the best sushi place?", Content: "I have cravings", Upvotes: 15, EditCount: 5}, &models.Question{ID: "0a24c4cd-4c73-42e4-bcca-3844d088de85", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", Title: "Will Jordans make me a sick baller?", Content: "I need to improve my game", Upvotes: 10, EditCount: 1}}

	retrievedQuestions := GlobalPostStore.FindQuestionsByAuthor("posted-by", "Tester2")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

	retrievedQuestions = GlobalPostStore.FindQuestionsByAuthor("answered-by", "Tester1")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

}

func TestFindQuestionsByFilter(t *testing.T) {

	//Test postComponent: "question", filter: "upvotes", order: "desc"
	expectedQuestions := []*models.Question{&models.Question{ID: "526c4576-0e49-4e90-b760-e6976c698574", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", Title: "Where is the best sushi place?", Content: "I have cravings", Upvotes: 15, EditCount: 5}, &models.Question{ID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", AuthorID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Author: "Tester1", Title: "What is should my squat to bench ratio be?", Content: "I need gains", Upvotes: 13, EditCount: 4}, &models.Question{ID: "0a24c4cd-4c73-42e4-bcca-3844d088de85", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", Title: "Will Jordans make me a sick baller?", Content: "I need to improve my game", Upvotes: 10, EditCount: 1}}

	retrievedQuestions := GlobalPostStore.FindQuestionsByFilter("question", "upvotes", "DESC", "0")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

	//Test postComponent: "answer", filter: "date", order: "asc"
	expectedQuestions = []*models.Question{&models.Question{ID: "38681976-4d2d-4581-8a68-1e4acfadcfa0", AuthorID: "0c1b2b91-9164-4d52-87b0-9c4b444ee62d", Author: "Tester1", Title: "What is should my squat to bench ratio be?", Content: "I need gains", Upvotes: 13, EditCount: 4}, &models.Question{ID: "526c4576-0e49-4e90-b760-e6976c698574", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", Title: "Where is the best sushi place?", Content: "I have cravings", Upvotes: 15, EditCount: 5}, &models.Question{ID: "0a24c4cd-4c73-42e4-bcca-3844d088de85", AuthorID: "95954f28-a8c3-4e76-8c80-18de07931639", Author: "Tester2", Title: "Will Jordans make me a sick baller?", Content: "I need to improve my game", Upvotes: 10, EditCount: 1}}

	retrievedQuestions = GlobalPostStore.FindQuestionsByFilter("answer", "date", "ASC", "0")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)
}

func TestIsTitleUnique(t *testing.T) {

	//Test for the recognition of an existing question with a matching title
	if _, id := GlobalPostStore.IsTitleUnique("Where is the best sushi place?"); id != "526c4576-0e49-4e90-b760-e6976c698574" {
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

	for i, _ := range y {
		standardizeTime(&y[i].SubmittedAt, &x[i].SubmittedAt)
		if !reflect.DeepEqual(x[i], y[i]) {
			t.Errorf("\n\nExpected %#v,\n but recieved %#v\n\n", y[i], x[i])
		}
	}
}
