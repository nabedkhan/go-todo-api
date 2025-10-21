package handlers

import (
	"net/http"

	"github.com/nabedkhan/go-todo-api/types"
	"github.com/nabedkhan/go-todo-api/utils"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, types.Response{
		Message: "Go server is running on",
		Success: true,
	})
}
