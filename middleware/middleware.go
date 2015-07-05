package middleware

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ParseRequestBody(r *http.Request, loc interface{}) (string, int) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return "Error reading request body", http.StatusInternalServerError
	}

	if err = r.Body.Close(); err != nil {
		return "", http.StatusInternalServerError
	}

	//Unmarshals JSON data into loc param
	err = json.Unmarshal(body, loc)
	if err != nil {
		return err.Error() + "\n", 422 //unprocessable entity
	}

	return "", 0

}

func PrintJSON(w http.ResponseWriter, content interface{}) {

	w.Header().Set("Content-Type", "application/json")

	postJSON, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(postJSON)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

// Ensures that POST requests contain data in the request body
func CheckRequestBody(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			http.Error(w, "This api only accepts JSON payloads. Be sure to specify the \"Content-Type\" of the payload in the request header.", http.StatusBadRequest)
		} else if r.Body == nil {
			http.Error(w, "No data recieved through the request", http.StatusBadRequest)
		} else {
			fn(w, r)
		}
	}
}
