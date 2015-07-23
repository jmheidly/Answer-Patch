package models

import (
	"time"
)

type Question struct {
	ID          string    `json:"questionID"`
	AuthorID    string    `json:"authorID"`
	Author      string    `json:"questionAuthor"`
	Title       string    `json:"questionTitle"`
	Content     string    `json:"questionContent"`
	Upvotes     int       `json:"questionUpvotes"`
	EditCount   int       `json:"answerEditCount"`
	SubmittedAt time.Time `json:"questionSubmittedAt"`
}

func (question *Question) GetMissingFields() string {

	var missing string

	switch {
	case question.AuthorID == "":
		missing = "Author's ID\n"
	case question.Author == "":
		missing += "Author's username\n"
	case question.Title == "":
		missing += "Title\n"
	}

	return missing
}
