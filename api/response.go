package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

// sendError sending response with error based on error + status code.
func sendError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(Error{Error: err.Error()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func sendResponse(w http.ResponseWriter, statusCode int, v ...interface{}) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
