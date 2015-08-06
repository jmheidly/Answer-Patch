package app

import (
	"net/http"

	m "github.com/patelndipen/AP1/middleware"
)

func ServeLogout() m.HandlerFunc {
	return func(c *m.Context, w http.ResponseWriter, r *http.Request) {
		c.Logout(r.Header.Get("Authorization")[7:]) //Sends the signed token without the "BEARER:" prefix
	}
}
