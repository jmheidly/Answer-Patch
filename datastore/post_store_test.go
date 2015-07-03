package datastore

import (
	"reflect"
	"testing"
	"time"

	"github.com/patelndipen/AP1/models"
)

var testingGlobalPostStore *PostStore

func init() {

	testingGlobalPostStore = NewPostStore(ConnectToDatabase("postgres", "test", "testap1poststore"))
	//	initializeDatabase(testingGlobalPostStore.DB)
	//	populateDatabase(testingGlobalPostStore.DB)

}

func TestFindByID(t *testing.T) {

	expectedQuestion := &models.Question{2, "How many days are there in spring?", "postStoreTester2", "I love spring weather", 13, parseTimeStamp("0000-01-01T22:20:05.714972-04:00"), 4}

	expectedAnswer := &models.Answer{2, 2, true, "postStoreTester1", "Spring 2016 starts on March 20th", 20, parseTimeStamp("0000-01-01T22:20:05.748316-04:00")}

	retrievedQuestion, retrievedAnswer := testingGlobalPostStore.FindByID("2")

	if retrievedQuestion == nil {
		t.Errorf("Expected and did not recieve %#v", expectedQuestion)
	}

	// Standardizes the time zone for time.Time values
	retrievedQuestion.SubmittedAt = retrievedQuestion.SubmittedAt.Local()
	expectedQuestion.SubmittedAt = expectedQuestion.SubmittedAt.Local()

	if !reflect.DeepEqual(retrievedQuestion, expectedQuestion) {
		t.Errorf("Expected %#v, but recieved %#v", expectedQuestion, retrievedQuestion)
	}

	if retrievedAnswer == nil {
		t.Errorf("Expected and did not recieve %#v", expectedAnswer)
	}

	// Standardizes the time zone for time.Time values
	retrievedAnswer.LastEditedAt = expectedAnswer.LastEditedAt.Local()
	expectedAnswer.LastEditedAt = expectedAnswer.LastEditedAt.Local()

	if !reflect.DeepEqual(expectedAnswer, retrievedAnswer) {
		t.Errorf("Expected %#v, but recieved %#v", expectedAnswer, retrievedAnswer)
	}

}

func TestFindByAuthor(t *testing.T) {

	// Tests FindByAuthor based off of providing the author of a questions
	expectedQuestions := []*models.Question{&models.Question{1, "Where is the best sushi place?", "postStoreTester1", "", 10, parseTimeStamp("0000-01-01T22:20:05.70397-04:00"), 2}, &models.Question{3, "What should my bench to squat ratio be?", "postStoreTester1", "", 100, parseTimeStamp("0000-01-01T22:20:05.726089-04:00"), 32}}

	retrievedQuestions := testingGlobalPostStore.FindByAuthor("posted-by", "postStoreTester1")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

	// Tests FindByAuthor based off of providing the author of an answer
	expectedQuestions = []*models.Question{{2, "How many days are there in spring?", "postStoreTester2", "", 13, parseTimeStamp("0000-01-01T22:20:05.714972-04:00"), 4}, {3, "What should my bench to squat ratio be?", "postStoreTester1", "", 100, parseTimeStamp("0000-01-01T22:20:05.726089-04:00"), 32}}

	retrievedQuestions = testingGlobalPostStore.FindByAuthor("answered-by", "postStoreTester1")

	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

}

func TestFindByFilter(t *testing.T) {

	//Test for "question/upvotes/desc"
	expectedQuestions := []*models.Question{{3, "What should my bench to squat ratio be?", "postStoreTester1", "", 100, parseTimeStamp("0000-01-01T22:20:05.726089-04:00"), 32}, {2, "How many days are there in spring?", "postStoreTester2", "", 13, parseTimeStamp("0000-01-01T22:20:05.714972-04:00"), 4}, {1, "Where is the best sushi place?", "postStoreTester1", "", 10, parseTimeStamp("0000-01-01T22:20:05.70397-04:00"), 2}}
	retrievedQuestions := testingGlobalPostStore.FindByFilter("question", "upvotes", "DESC", "0")
	checkQuestionsForEquality(t, retrievedQuestions, expectedQuestions)

}

func TestCheckQuestionExistence(t *testing.T) {

	//Test for recognition of existing question with exact title
	if id := testingGlobalPostStore.CheckQuestionExistence("Where is the best sushi place?"); id != 1 {
		t.Errorf("CreateQuestion did not return the id of an existing question, instead CreateQuestion returned %d", id)
	}

}

func parseTimeStamp(timestamp string) time.Time {

	// Parses a string representing the time.Time value into the corresponding time.Time value
	formattedTime, err := time.Parse(time.RFC3339Nano, timestamp)
	if err != nil {
		panic(err)
	}
	return formattedTime
}

func checkQuestionsForEquality(t *testing.T, x []*models.Question, y []*models.Question) {

	if x == nil {
		t.Errorf("Recieved nil, but expected %#v", y)
	}

	for i, _ := range x {
		// Standardizes the time zone for time.Time values
		x[i].SubmittedAt = x[i].SubmittedAt.Local()
		y[i].SubmittedAt = y[i].SubmittedAt.Local()

		if !reflect.DeepEqual(x[i], y[i]) {
			t.Errorf("\n\nExpected %#v,\n but recieved %#v\n\n", y[i], x[i])
		}
	}
}
