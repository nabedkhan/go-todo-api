package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nabedkhan/go-todo-api/db"
	"github.com/nabedkhan/go-todo-api/types"
	"github.com/nabedkhan/go-todo-api/utils"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var body types.Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	if err != nil {
		utils.SendError(w, r, "Title is required!", http.StatusBadRequest)
		return
	}

	body.Completed = false
	body.Id = db.TodoList[len(db.TodoList)-1].Id + 1

	TodoList := append(db.TodoList, body)

	utils.SendJSON(w, types.Response{
		Message: "New Todo created successfully",
		Success: true,
		Data:    TodoList,
	})
}
