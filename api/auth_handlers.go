package api

import (
	"fmt"
	"net/http"

	"github.com/patelndipen/AP1/datastore"
	m "github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

func ServeSubmitQuestion(store datastore.PostStoreServices) m.HandlerFunc {
	return func(c *m.Context, w http.ResponseWriter, r *http.Request) {

		//Code to check the user is authorized to post a question
		newQuestion := c.ParsedModel.(*models.Question)
		isUnique, existingID := store.IsTitleUnique(newQuestion.Title)
		// Redirects the user to the post with the matching question title
		if !isUnique {
			url := fmt.Sprintf("http://localhost:8080/api/posts/%d", existingID)
			http.Redirect(w, r, url, 303) // Status 303 - See Other
		} else {
			store.StoreQuestion(newQuestion.AuthorID, newQuestion.Title, newQuestion.Content, newQuestion.Category)
			w.WriteHeader(http.StatusCreated)
		}
	}
}
