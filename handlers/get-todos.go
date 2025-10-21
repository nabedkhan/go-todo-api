package handlers

import (
	"net/http"

	"github.com/nabedkhan/go-todo-api/db"
	"github.com/nabedkhan/go-todo-api/types"
	"github.com/nabedkhan/go-todo-api/utils"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, types.Response{
		Message: "Todos fetched successfully",
		Success: true,
		Data:    db.GetTodos(),
	})
}
