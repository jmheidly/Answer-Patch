package models

import (
	"time"
)

type Question struct {
	ID          string    `json:"questionID"`
	UserID      string    `json:"userID"`
	Username    string    `json:"questionUsername"`
	Title       string    `json:"questionTitle"`
	Content     string    `json:"questionContent"`
	Upvotes     int       `json:"questionUpvotes"`
	EditCount   int       `json:"answerEditCount"`
	SubmittedAt time.Time `json:"questionSubmittedAt"`
}

func NewQuestion() *Question {
	return &Question{}
}
