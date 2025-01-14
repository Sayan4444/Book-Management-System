package controllers

import (
	"encoding/json"
	"net/http"
)

func readJSON(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, message string, status int) error {
	data := map[string]string{"error": message}
	return writeJSON(w, status, data)
}
