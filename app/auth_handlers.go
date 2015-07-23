package app

import (
	"net/http"

	"github.com/patelndipen/AP1/middleware"
	"github.com/patelndipen/AP1/models"
)

func ServeLogout() middleware.ContextRequiredHandlerFunc {
	return func(c *models.Context, w http.ResponseWriter, r *http.Request) {
		c.Logout(r.Header.Get("Authorization")[7:]) //Sends the signed token without the "BEARER:" prefix
	}
}
