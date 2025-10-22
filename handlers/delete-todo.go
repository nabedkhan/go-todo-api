package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nabedkhan/go-todo-api/db"
	"github.com/nabedkhan/go-todo-api/types"
	"github.com/nabedkhan/go-todo-api/utils"
)

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, r, "Invalid route param", http.StatusBadRequest)
		return
	}

	if intId > db.GetTodosLength() {
		utils.SendError(w, r, "Todo does not exist with given id", http.StatusNotFound)
		return
	}

	db.DeleteTodo(intId)

	utils.SendJSON(w, types.Response{
		Message: fmt.Sprintf("Todo %s deleted successfully", id),
		Success: true,
		Data:    db.GetTodos(),
	})
}
