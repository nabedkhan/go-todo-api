package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
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

	if intId > len(db.TodoList) {
		utils.SendError(w, r, "Todo does not exist with given id", http.StatusBadRequest)
		return
	}

	var body types.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	idx := slices.IndexFunc(db.TodoList, func(todo types.Todo) bool {
		return todo.Id == intId
	})

	if body.Title != "" {
		db.TodoList[idx].Title = body.Title
	}

	db.TodoList[idx].Completed = body.Completed

	utils.SendJSON(w, types.Response{
		Data:    db.TodoList[idx],
		Success: true,
		Message: "Todo " + id + " updated successfully",
	})
}
