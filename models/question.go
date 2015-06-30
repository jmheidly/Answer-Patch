package models

import (
	"time"
)

type Question struct {
	ID    int    `json:"questionID"`
	Title string `json:"questionTitle"`
	//User        *User     `json:"questionAuthor"`
	Author      string    `json:"questionAuthor"`
	Content     string    `json:"questionContent"`
	Upvotes     int       `json:"questionUpvotes"`
	SubmittedAt time.Time `json:"questionSubmittedAt"`
	EditCount   int       `json:"answerEditCount"`
}

func NewQuestion() *Question {
	return &Question{}
}
