package models

import (
	"time"
)

type Answer struct {
	ID              int  `json:"AnswerID"`
	QuestionID      int  `json:"AnswerquestionID"`
	IsCurrentAnswer bool `json:"AnswerCurrent"`
	//User            *User     `json:"AnswerAuthor"`
	Author       string    `json:"AnswerAuthor"`
	Content      string    `json:"AnswerContent"`
	Upvotes      int       `json:"AnswerUpvotes"`
	LastEditedAt time.Time `json:"AnswerLastEditedAt"`
}

func NewAnswer() *Answer {
	return &Answer{}
}
