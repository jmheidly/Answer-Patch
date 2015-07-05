package models

import (
	"time"
)

type Answer struct {
	ID              string    `json:"answerID"`
	QuestionID      string    `json:"questionID"`
	UserID          string    `json:"userID"`
	Username        string    `json:"answerUsername"`
	IsCurrentAnswer bool      `json:"answerCurrent"`
	Content         string    `json:"answerContent"`
	Upvotes         int       `json:"answerUpvotes"`
	LastEditedAt    time.Time `json:"answerLastEditedAt"`
}

func NewAnswer() *Answer {
	return &Answer{}
}
