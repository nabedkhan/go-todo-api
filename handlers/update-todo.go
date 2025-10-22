package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nabedkhan/go-todo-api/db"
	"github.com/nabedkhan/go-todo-api/types"
	"github.com/nabedkhan/go-todo-api/utils"
)

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, r, "Invalid param id", http.StatusNotFound)
		return
	}

	if intId > db.GetTodosLength() {
		utils.SendError(w, r, "Todo does not exist with given id", http.StatusBadRequest)
		return
	}

	var body types.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	updatedTodo := db.UpdateTodo(intId, body)

	utils.SendJSON(w, types.Response{
		Message: fmt.Sprintf("Todo %s updated successfully", id),
		Success: true,
		Data:    updatedTodo,
	})
}
