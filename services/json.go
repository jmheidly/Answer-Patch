package services

import (
	"encoding/json"
	"net/http"
)

//Marshals and writes JSON to the http.ResponseWriter
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
