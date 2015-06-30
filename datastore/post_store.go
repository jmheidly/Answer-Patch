package datastore

import (
	"database/sql"
	"log"

	"github.com/patelndipen/AP1/models"
)

type PostStoreActions interface {
	FindByID(string) (*models.Question, *models.Answer)
	FindByAuthor(string, string) []*models.Question
	FindByFilter(string, string) []*models.Question
	CheckQuestionExistence(string) int
	StoreQuestion(string, string, string)
}

type PostStore struct {
	DB *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db}
}

func (store *PostStore) FindByID(id string) (*models.Question, *models.Answer) {

	row, err := store.DB.Query(`SELECT id, title, author, content, upvotes, submitted_at, edit_count FROM question WHERE id = $1`, id)

	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		return nil, nil
	}

	question := models.NewQuestion()
	err = row.Scan(&question.ID, &question.Title, &question.Author,
		&question.Content, &question.Upvotes,
		&question.SubmittedAt, &question.EditCount)

	if err != nil {
		log.Fatal(err)
	}

	row, err = store.DB.Query(`SELECT id, question_id, is_current_answer, author, content, upvotes, last_edited_at FROM answer WHERE question_id = $1 AND is_current_answer = 'true'`, id)

	if err != nil {
		log.Fatal(err)
	} else if !row.Next() {
		return question, nil
	}

	answer := models.NewAnswer()
	err = row.Scan(&answer.ID, &answer.QuestionID, &answer.IsCurrentAnswer, &answer.Author, &answer.Content, &answer.Upvotes, &answer.LastEditedAt)
	if err != nil {
		log.Fatal(err)
	}

	return question, answer
}

func (store *PostStore) FindByAuthor(filter, author string) []*models.Question {
	var query string

	if filter == "posted-by" {
		query = `SELECT id, title, author, upvotes, submitted_at, edit_count FROM question WHERE author = $1`
	} else if filter == "answered-by" {
		query = `SELECT q.id, q.title, q.author, q.upvotes, q.submitted_at, q.edit_count FROM question q INNER JOIN answer a ON (a.author = $1 AND a.question_id = q.id)`
	} else {
		return nil
	}

	rows, err := store.DB.Query(query, author)
	if err != nil {
		log.Fatal(err)
	}

	return scanQuestions(rows)
}

func (store *PostStore) FindByFilter(filter, offset string) []*models.Question {
	filteredQueries := map[string]string{
		"question/upvotes/desc": `SELECT id, title, author, upvotes, submitted_at, edit_count FROM question ORDER BY upvotes DESC LIMIT 10 OFFSET $1`,
		"question/upvotes/asc":  `SELECT  id, title, author, upvotes, submitted_at, edit_count FROM  question ORDER BY upvotes ASC LIMIT 10 OFFSET $1`,
		"answer/upvotes/desc":   `SELECT  q.id, q.title, q.author, q.upvotes, q.submitted_at, q.edit_count FROM question q INNER JOIN answer a ON q.id = a.question_id ORDER BY a.upvotes DESC LIMIT 10 OFFSET $1`,
		"answer/upvotes/asc":    `SELECT  q.id, q.title, q.author, q.upvotes, q.submitted_at, q.edit_count FROM question q INNER JOIN answer a ON q.id = a.question_id ORDER BY a.upvotes ASC LIMIT 10 OFFSET $1`,
		"question/date/desc":    `SELECT id, title, author, upvotes, submitted_at, edit_count FROM question ORDER BY submitted_at DESC LIMIT 10 OFFSET $1`,
		"question/date/asc":     `SELECT  id, title, author, upvotes, submitted_at, edit_count FROM question ORDER BY submitted_at ASC LIMIT 10 OFFSET $1`,
		"answer/edits/desc":     `SELECT id, title, author, upvotes, submitted_at, edit_count FROM question ORDER BY edit_count DESC LIMIT 10 OFFSET $1`,
		"answer/edits/asc":      `SELECT id, title, author, upvotes, submitted_at, edit_count FROM question ORDER BY answer_edits ASC LIMIT 10 OFFSET $1`,
		"answer/date/desc":      `SELECT  q.id, q.title, q.author, q.upvotes, q.submitted_at, q.edit_count FROM question q INNER JOIN answer a ON q.id = a.question_id ORDER BY a.last_edited_at DESC LIMIT 10 OFFSET $1`,
		"answer/date/asc":       `SELECT  q.id, q.title, q.author, q.upvotes, q.submitted_at, q.edit_count FROM question q INNER JOIN answer a ON q.id = a.question_id ORDER BY a.last_edited_at ASC LIMIT 10 OFFSET $1`,
	}

	query, ok := filteredQueries[filter]
	if !ok {
		return nil
	}

	rows, err := store.DB.Query(query, offset)
	if err != nil {
		log.Fatal(err)
	}

	return scanQuestions(rows)
}

func (store *PostStore) CheckQuestionExistence(title string) int {

	var id int

	row, err := store.DB.Query("SELECT id FROM question WHERE title = $1", title)
	if err != nil {
		log.Fatal(err)
	} else if row.Next() {
		err = row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		return id
	}

	return 0
}

func (store *PostStore) StoreQuestion(title, author, content string) {

	transact(store.DB, func(tx *sql.Tx) error {
		_, err := tx.Exec(`INSERT INTO question(title, author, content) values( $1, $2, $3)`, title, author, content)
		return err
	})

}

func scanQuestions(rows *sql.Rows) []*models.Question {
	var questions []*models.Question

	for rows.Next() {
		tempQuestion := models.NewQuestion()
		err := rows.Scan(&tempQuestion.ID, &tempQuestion.Title, &tempQuestion.Author, &tempQuestion.Upvotes, &tempQuestion.SubmittedAt, &tempQuestion.EditCount)
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
