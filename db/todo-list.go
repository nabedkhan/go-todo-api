package db

import (
	"slices"

	"github.com/nabedkhan/go-todo-api/types"
)

var todoList = []types.Todo{
	{Id: 1, Title: "Learning Go", Completed: false},
	{Id: 2, Title: "Learning React", Completed: false},
}

func GetLastTodoId() int {
	return todoList[GetTodosLength()-1].Id
}

func GetTodosLength() int {
	return len(todoList)
}

func GetTodos() []types.Todo {
	return todoList
}

func GetTodoById(id int) types.Todo {
	idx := slices.IndexFunc(todoList, func(todo types.Todo) bool {
		return todo.Id == id
	})

	return todoList[idx]
}

func CreateTodo(body types.Todo) {
	todoList = append(todoList, body)
}

func DeleteTodo(id int) {
	todoList = slices.DeleteFunc(todoList, func(todo types.Todo) bool {
		return todo.Id == id
	})
}

func UpdateTodo(id int, body types.Todo) types.Todo {
	idx := slices.IndexFunc(todoList, func(todo types.Todo) bool {
		return todo.Id == id
	})

	todoList[idx].Title = body.Title
	todoList[idx].Completed = body.Completed

	return todoList[idx]
}
