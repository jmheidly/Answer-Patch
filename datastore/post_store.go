package datastore

import (
	"database/sql"
	"log"
	"strings"

	"github.com/patelndipen/AP1/models"
)

type PostStoreServices interface {
	FindPostByID(string) (*models.Question, *models.Answer)
	FindQuestionsByUser(string, string) []*models.Question
	FindQuestionsByFilter(string, string, string, string) []*models.Question
	CheckQuestionExistence(string) string
	StoreQuestion(string, string, string)
}

type PostStore struct {
	DB *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db}
}

func (store *PostStore) FindPostByID(id string) (*models.Question, *models.Answer) {

	row, err := store.DB.Query(`SELECT q.user_id, u.username, q.title, q.content, q.upvotes, q.edit_count, q.submitted_at FROM question q INNER JOIN ap_user u ON q.user_id = u.id WHERE q.id =$1`, id)
	if err != nil {
		log.Fatal(err)
	} else if !row.Next() { // Checks if rows were returned by the query
		return nil, nil
	}

	question := models.NewQuestion()
	err = row.Scan(&question.UserID, &question.Username, &question.Title, &question.Content, &question.Upvotes, &question.EditCount, &question.SubmittedAt)
	if err != nil {
		log.Fatal(err)
	}

	row, err = store.DB.Query(`SELECT a.id, a.user_id, u.username, a.is_current_answer, a.content, a.upvotes, a.last_edited_at FROM answer a INNER JOIN ap_user u ON a.user_id = u.id WHERE a.question_id = $1 AND is_current_answer = 'true'`, id)
	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		return question, nil // A question could possibly lack any valid answer at the current moment
	}

	answer := models.NewAnswer()
	err = row.Scan(&answer.ID, &answer.UserID, &answer.Username, &answer.IsCurrentAnswer, &answer.Content, &answer.Upvotes, &answer.LastEditedAt)
	if err != nil {
		log.Fatal(err)
	}

	question.ID = id
	answer.QuestionID = id

	return question, answer
}

func (store *PostStore) FindQuestionsByUser(filter, user string) []*models.Question {

	queryStmt := `SELECT q.id, q.user_id, u.username, q.title, q.content, q.upvotes, q.edit_count, q.submitted_at FROM question q`

	if filter == "posted-by" { // The user param is the username of auser that asked a question
		queryStmt += ` INNER JOIN ap_user u ON q.user_id = u.id WHERE u.username = $1`
	} else if filter == "answered-by" { // The user param is the username of a user that answered a question
		queryStmt += ` JOIN ap_user answer_author ON answer_author.username = $1 JOIN answer a ON (answer_author.id = a.user_id AND a.question_id=q.id AND a.is_current_answer='true') INNER JOIN ap_user u ON q.user_id = u.id`
	}

	rows, err := store.DB.Query(queryStmt, user)
	if err != nil {
		log.Fatal(err)
	}

	return scanQuestions(rows)
}

func (store *PostStore) FindQuestionsByFilter(postComponent, filter, order, offset string) []*models.Question {

	var ok bool

	// These maps convert the url param "filter" into a valid field that can be used as by ORDER BY in an sql query
	questionFilters := map[string]string{
		"upvotes": "q.upvotes",
		"date":    "q.submitted_at",
		"edits":   "q.edit_count",
	}

	answerFilters := map[string]string{
		"upvotes": "a.upvotes",
		"date":    "a.last_edited_at",
	}

	queryStmt := `SELECT q.id, q.user_id, u.username, q.title, q.content, q.upvotes, q.edit_count, q.submitted_at FROM question q`
	if postComponent == "question" {
		queryStmt += ` INNER JOIN ap_user u ON q.user_id = u.id`
		filter, ok = questionFilters[filter]
	} else if postComponent == "answer" {
		queryStmt += ` JOIN answer a ON (a.question_id=q.id AND a.is_current_answer='true') INNER JOIN ap_user u ON q.user_id = u.id`
		filter, ok = answerFilters[filter]
	}

	if !ok { // Returns nil if the url param "filter" is not a field capable of being used as a ORDER BY specification in an sql query
		return nil
	}

	queryStmt += ` ORDER BY ` + filter + ` ` + strings.ToUpper(order) + ` LIMIT 10 OFFSET $1`

	rows, err := store.DB.Query(queryStmt, offset)
	if err != nil {
		log.Fatal(err)
	}

	return scanQuestions(rows)
}

func (store *PostStore) CheckQuestionExistence(title string) string {

	var existingID string

	row, err := store.DB.Query("SELECT id FROM question WHERE title = $1", title)
	if err != nil {
		log.Fatal(err)
	} else if row.Next() {
		err = row.Scan(&existingID)
		if err != nil {
			log.Fatal(err)
		}
		return existingID
	}

	return ""
}

func (store *PostStore) StoreQuestion(user_id, title, content string) {

	transact(store.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`INSERT INTO question(user_id, title, content) values($1::uuid, $2, $3)`, user_id, title, content)
		return err
	})

}

func scanQuestions(rows *sql.Rows) []*models.Question {
	var questions []*models.Question

	for rows.Next() {
		tempQuestion := models.NewQuestion()
		err := rows.Scan(&tempQuestion.ID, &tempQuestion.UserID, &tempQuestion.Username, &tempQuestion.Title, &tempQuestion.Content, &tempQuestion.Upvotes, &tempQuestion.EditCount, &tempQuestion.SubmittedAt)
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, tempQuestion)
	}

	if len(questions) == 0 {
		return nil
	}
	return questions
}
