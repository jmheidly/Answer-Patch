package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/services"
	"github.com/patelndipen/AP1/settings"
)

type ContextRequiredHandlerFunc func(*models.Context, http.ResponseWriter, *http.Request)

func ServeHTTP(fn ContextRequiredHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := models.NewContext(nil)
		fn(c, w, r)
	}
}

func ParseRequestBody(model models.ModelServices, fn ContextRequiredHandlerFunc) ContextRequiredHandlerFunc {
	return func(c *models.Context, w http.ResponseWriter, r *http.Request) {

		//Checks whether the request body is in JSON format
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			http.Error(w, "This api only accepts JSON payloads. Be sure to specify the \"Content-Type\" of the payload in the request header.", http.StatusBadRequest)
			return
		} else if r.Body == nil {
			http.Error(w, "No data recieved through the request", http.StatusBadRequest)
			return
		}

		//Parses request body

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			log.Fatal(err)
		}

		defer r.Body.Close()

		err = json.Unmarshal(body, model)
		if err != nil {
			// 422 -unprocessable entity
			http.Error(w, err.Error()+"\n", 422)
			return
		}

		missing := model.GetMissingFields()
		if missing != "" {
			http.Error(w, "The following fields were not recieved:\n"+missing, http.StatusBadRequest)
			return
		}

		c.ParsedModel = model

		fn(c, w, r)
	}
}

func RefreshExpiringToken(fn ContextRequiredHandlerFunc) ContextRequiredHandlerFunc {
	return func(c *models.Context, w http.ResponseWriter, r *http.Request) {

		//Refreshes token, if the token expires in less than 24 hours
		if c.Exp.Sub(time.Now()) < (time.Duration(24) * time.Hour) {
			services.PrintJSON(w, (c.RefreshToken()))
		}
	}
}

func AuthenticateToken(c *models.Context, fn ContextRequiredHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := jwt.ParseFromRequest(r, func(parsedToken *jwt.Token) (interface{}, error) {
			if _, ok := parsedToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unrecognized signing method: %v", parsedToken.Header["alg"])

			} else {
				return settings.GetPublicKey(), nil
			}
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid JWT", http.StatusUnauthorized)
			return
		}

		var ok bool

		c.UserID, ok = token.Claims["sub"].(string)
		if !ok {
			log.Fatal("The underlying type of sub is not string")
		}

		if c.TokenStore.IsTokenStored(c.UserID) {
			http.Error(w, "Token is no longer valid", http.StatusUnauthorized)
			return
		}

		exp, ok := token.Claims["exp"].(string)
		if !ok {
			log.Fatal("The underlying type of exp is not string")
		}
		c.Exp, err = time.Parse(time.RFC3339Nano, exp)
		if err != nil {
			log.Fatal(err)
		}

		fn(c, w, r)
	}

}
