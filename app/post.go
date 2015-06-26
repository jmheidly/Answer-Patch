package app

import (
	"time"
)

type Question struct {
	ID          int       `json:"questionID"`
	Title       string    `json:"questionTitle"`
	Author      string    `json:"questionAuthor"`
	Content     string    `json:"questionContent"`
	Upvotes     int       `json:"questionUpvotes"`
	SubmittedAt time.Time `json:"questionSubmittedAt"`
	EditCount   int       `json:"answerEditCount"`
}
type Answer struct {
	ID              int       `json:"AnswerID"`
	QuestionID      int       `json:"AnswerquestionID"`
	IsCurrentAnswer bool      `json:"AnswerCurrent"`
	Author          string    `json:"AnswerAuthor"`
	Content         string    `json:"AnswerContent"`
	Upvotes         int       `json:"AnswerUpvotes"`
	LastEditedAt    time.Time `json:"AnswerLastEditedAt"`
}

// change the Qauthor and CurrentAnsAuthor to User values, once the User struct has been declared

func NewQuestionStruct() *Question {
	return &Question{}
}

func NewAnswerStruct() *Answer {
	return &Answer{}
}
