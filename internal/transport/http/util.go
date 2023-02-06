package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func getIdFromRequest(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("id", id)

	if id == "" {
		return "", errors.New("id can't be 0")
	}

	return id, nil
}

func sendResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("sendResponse() error:", err)
		}
	}
}

// sendErrorResponse sends error response to the client
func sendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
