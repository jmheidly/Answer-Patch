package app

import (
	"time"
)

type Post struct {
	Postid                  int       `json:"postId"`
	Qauthor                 string    `json:"questionAuthor"`
	Qtitle                  string    `json:"questionTitle"`
	Qcontent                string    `json:"questionContent"`
	Qupvotes                int       `json:"questionUpvotes"`
	QsubmittedAt            time.Time `json:"questionSubmittedAt"`
	AnsLastEditedAt         time.Time `json:"answerLastEditedAt"`
	AnsEdits                int       `json:"answerEdits"`
	CurrentAns              string    `json:"currentAnswer"`
	CurrentAnsUpvotes       int       `json:"currentAnswerUpvotes"`
	CurrentAnsAuthor        string    `json:"currentAnswerAuthor"`
	FirstPendingAns         string    `json:"firstPendingAnswer"`
	FirstPendingAnsUpvotes  int       `json:"firstPendingAnswerUpvotes"`
	FirstPendingAnsAuthor   string    `json:"firstPendingAnswerAuthor"`
	SecondPendingAns        string    `json:"secondPendingAnswer"`
	SecondPendingAnsUpvotes int       `json:"secondPendingAnswerUpvotes"`
	SecondPendingAnsAuthor  string    `json:"secondPendingAnswerAuthor"`
}

// change the Qauthor and CurrentAnsAuthor to User values, once the User struct has been declared

func NewPostStruct() *Post {
	return &Post{}
}
