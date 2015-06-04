package datastore

// filteredByOrderQueries := map[string]string{
// 		"descending_question_upvotes":       `SELECT * FROM post ORDER BY question_upvotes DESC LIMIT 10`,
// 		"ascending_question_upvotes":        `SELECT * FROM post ORDER BY question_upvotes ASC LIMIT 10`,
// 		"descending_current_answer_upvotes": `SELECT * FROM post ORDER BY current_answer_upvotes DESC LIMIT 10`,
// 		"ascending_current_answer_upvotes":  `SELECT * FROM post ORDER BY current_answer_upvotes ASC LIMIT 10`,
// 		"descending_question_date":          `SELECT * FROM post ORDER BY question_submitted_at DESC LIMIT 10`,
// 		"ascending_question_date":           `SELECT * FROM post ORDER BY question_submitted_at ASC LIMIT 10`,
// 		"descending_answers_edits":          `SELECT * FROM post ORDER BY answer_edits DESC LIMIT 10`,
// 		"ascending_answer_edits":            `SELECT * FROM post ORDER BY answer_edits ASC LIMIT 10`,
// 		"descending_answer_edits":           `SELECT * FROM post ORDER BY answer_last_edited_at DESC LIMIT 10`,
// 		"ascending_answer_edits":            `SELECT * FROM post ORDER BY answer_last_edited_at ASC LIMIT 10`,
// 	}
//
import (
	"database/sql"
	"log"

	"github.com/patelndipen/AP1/app"
)

type PostStoreActions interface {
	FindByID(string) (*app.Post, error)
	FindByFilter(string, string) ([]*app.Post, error)
}

type PostStore struct {
	DB *sql.DB
}

func NewPostStore() *PostStore {
	return &PostStore{}
}

func (store *PostStore) FindByID(id string) (*app.Post, error) {
	row, err := store.DB.Query(`SELECT * FROM post WHERE post_id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	// Have to call ScanPosts even if rows does not actually contain any returned rows, because the only way of checking if rows contains sql rows is to call the Next() function on rows. The Next() function also iterates through rows if rows exists, therefore if rows does exists, Next() is called on rows, then rows has lost a row because one rows has been iterated by the call to Next()
	posts, err := ScanPosts(row)
	if err != nil {
		return nil, err
	}
	return posts[0], err
}

func (store *PostStore) FindByFilter(filterBy, filterVal string) ([]*app.Post, error) {
	filteredByValQueries := map[string]string{
		"user/questions":       `SELECT * FROM post WHERE question_author = $1`,
		"user/posted_answers":  `SELECT * FROM post WHERE current_answer_author = $1`,
		"user/pending_answers": `SELECT * FROM post WHERE first_pending_answer_author = $1 OR WHERE second_pending_answer_author = $1 OR WHERE third_pending_answer_author = $1 OR WHERE fourth_pending_answer_author = $1 OR WHERE fifth_pending_answer_author = $1`,
	}
	if query, ok := filteredByValQueries[filterBy]; ok {
		rows, err := store.DB.Query(query, filterVal)
		if err != nil {
			log.Fatal(err)
		}
		return ScanPosts(rows)
	}

	return nil, errInvalidFilter

}

func ScanPosts(rows *sql.Rows) ([]*app.Post, error) {

	defer rows.Close()

	var posts []*app.Post
	post := app.NewPostStruct()
	scanablePost := []interface{}{
		&post.Postid, &post.Qauthor, &post.Qtitle,
		&post.Qcontent, &post.Qupvotes,
		&post.QsubmittedAt, &post.AnsLastEditedAt,
		&post.AnsEdits, &post.CurrentAns,
		&post.CurrentAnsUpvotes,
		&post.CurrentAnsAuthor,
		&post.FirstPendingAns,
		&post.FirstPendingAnsUpvotes,
		&post.FirstPendingAnsAuthor,
		&post.SecondPendingAns,
		&post.SecondPendingAnsUpvotes,
		&post.SecondPendingAnsAuthor,
	}

	for rows.Next() {
		columns, err := rows.Columns()
		columnCount := len(columns)
		if err != nil {
			log.Fatal(err)
		}
		scanablePtrs := make([]interface{}, columnCount)
		for i := 0; i < columnCount; i++ {
			scanablePtrs[i] = scanablePost[i]
		}
		err = rows.Scan(scanablePtrs...)

		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
		post = app.NewPostStruct()
	}
	if len(posts) == 0 {
		return nil, errInvalidRequestParam
	}
	return posts, nil

}
