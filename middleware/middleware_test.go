package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patelndipen/AP1/models"
	"github.com/patelndipen/AP1/settings"
)

type MockModel struct {
	Field string
}

type MockTokenStore struct {
	IsStored bool
}

func init() {
	settings.SetPreproductionEnv()
}

func (model *MockModel) GetMissingFields() string {
	if model.Field == "" {
		return "Field"
	}
	return ""
}

func (store *MockTokenStore) StoreToken(key, val string, exp int) {
}

func (store *MockTokenStore) IsTokenStored(key string) bool {
	return store.IsStored
}

func TestParseRequestBody(t *testing.T) {

	model := &MockModel{Field: "value"}

	body, err := json.Marshal(model)
	if err != nil {
		t.Error(err)
	}

	r, err := http.NewRequest("", "", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ParseRequestBody(new(MockModel), func(c *models.Context, w http.ResponseWriter, r *http.Request) {

		parsedModel, ok := c.ParsedModel.(*MockModel)
		if !ok {
			http.Error(w, "context.ParsedModel is not of type*MockModel", http.StatusInternalServerError)
		}

		w.Write([]byte(parsedModel.Field))
	})(models.NewContext(nil), w, r)

	if parsedField := w.Body.String(); parsedField != model.Field {
		t.Errorf("Expected parsedModel.Field to equal %s, but instead %s was retrieved by parsing the request body ", model.Field, parsedField)
	}
}

func TestParseRequestBodyOnEmptyBody(t *testing.T) {

	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ParseRequestBody(new(MockModel), func(c *models.Context, w http.ResponseWriter, r *http.Request) {})(models.NewContext(nil), w, r)

	if w.Code != 400 {
		t.Errorf("Expected a status code of 400 due to the absence of a request body, recieved a status code of %d", w.Code)
	} else if errMessage := w.Body.String(); errMessage != "No data recieved through the request\n" {
		t.Errorf("Expected \"No data recieved through the request\" to be written to the responsewriter body, but the responsewriter body contains: %s", errMessage)
	}
}

func TestRefreshToken(t *testing.T) {

	w := httptest.NewRecorder()

	RefreshExpiringToken(func(c *models.Context, w http.ResponseWriter, r *http.Request) {})(models.NewContext(nil), w, nil)

	if w.Body == nil {
		t.Errorf("Expected RefreshToken to print a token to the responsewriter body")
	}

}

func TestAuthenticateTokenWithInvalidToken(t *testing.T) {

	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Authorization", "BEARER:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")

	w := httptest.NewRecorder()

	AuthenticateToken(models.NewContext(&MockTokenStore{IsStored: false}), func(c *models.Context, w http.ResponseWriter, r *http.Request) {})(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected the status code to be 401, because of the request contained an invalid JWT, but instead recieved a status code of %d", w.Code)
	} else if w.Body.String() != "Unrecognized signing method: HS256\n" {
		t.Errorf("Expected the responsewriter body to be set to \"Unrecognized signing method: HS256\", but instead the responsewriter body is set to \"%s\"", w.Body.String())
	}
}

func TestAuthenticateTokenParseOperations(t *testing.T) {

	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}

	//JWT token with a "sub" claim and UserID of "0"
	r.Header.Set("Authorization", "BEARER:eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDE1LTA3LTI1VDAxOjEzOjAzLjY3NTc2NjY1LTA0OjAwIiwiaWF0IjoxNDM3NTQxOTgzLCJzdWIiOiIwIn0.NlF6R66jEnBXUPYP4yQI2tI4rkV_JSFCFkjrvMmLXTDafXoLjCU7zWzruEGHiceaDyjuOV4poUz0riBrWk2Cx2148NFMdtwL8uviC-g5jJsRLpjzzL35pPLJ3y5vsqFRCzKuPRTTffIxda2GbjMqFe6uErzgH03VP2G3b_jDdWQ")

	w := httptest.NewRecorder()

	AuthenticateToken(models.NewContext(&MockTokenStore{IsStored: false}), func(c *models.Context, w http.ResponseWriter, r *http.Request) { w.Write([]byte(c.UserID)) })(w, r)

	if w.Body.String() != "0" {
		t.Errorf("Expected the UserID that AuthenticateToken is supposed to determine by parsing the JWT to be \"0\", but the UserID retrieved from the context struct was %s", w.Body.String())
	}
}

func TestAuthenticateTokenWithStoredToken(t *testing.T) {

	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Error(err)
	}

	r.Header.Set("Authorization", "BEARER:eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDE1LTA3LTI1VDAxOjEzOjAzLjY3NTc2NjY1LTA0OjAwIiwiaWF0IjoxNDM3NTQxOTgzLCJzdWIiOiIwIn0.NlF6R66jEnBXUPYP4yQI2tI4rkV_JSFCFkjrvMmLXTDafXoLjCU7zWzruEGHiceaDyjuOV4poUz0riBrWk2Cx2148NFMdtwL8uviC-g5jJsRLpjzzL35pPLJ3y5vsqFRCzKuPRTTffIxda2GbjMqFe6uErzgH03VP2G3b_jDdWQ")

	w := httptest.NewRecorder()

	AuthenticateToken(models.NewContext(&MockTokenStore{IsStored: true}), func(c *models.Context, w http.ResponseWriter, r *http.Request) {})(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected the status code to be a 401, because AuthenticateToken recognized that the token is stored in Redis due to the mock IsTokenStored method always returning true, but recieved a status code of %d", w.Code)
	} else if w.Body.String() != "Token is no longer valid\n" {
		t.Errorf("Expected the responsewriter body to contain \"Token is no longer valid\", because AuthenticateToken recognized that the token is stored in Redis due to the mock IsTokenStored method always returning true, but the responsewriter contained %s", w.Body.String())
	}
}
