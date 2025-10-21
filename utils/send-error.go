package utils

import (
	"net/http"

	"github.com/nabedkhan/go-todo-api/types"
)

func SendError(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	w.WriteHeader(statusCode)

	SendJSON(w, types.Response{
		Message: message,
		Success: false,
		Data:    nil,
	})
}
