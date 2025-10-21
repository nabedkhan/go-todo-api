package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

// Structs
type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Text struct {
	Message string `json:"message"`
}

// global data
var TodoList = []Todo{
	{Id: 1, Title: "Learning Go", Completed: false},
	{Id: 2, Title: "Learning React", Completed: false},
}

// Route Handlers
func RootHandler(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, Text{Message: "Go server is running on"})
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, TodoList)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		SendError(w, r, "Invalid route param", http.StatusBadRequest)
		return
	}

	if intId > len(TodoList) {
		SendError(w, r, "Todo does not exist with given id", http.StatusNotFound)
		return
	}

	idx := slices.IndexFunc(TodoList, func(todo Todo) bool {
		return todo.Id == intId
	})

	SendJSON(w, TodoList[idx])
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		SendError(w, r, "Invalid route param", http.StatusBadRequest)
		return
	}

	if intId > len(TodoList) {
		SendError(w, r, "Todo does not exist with given id", http.StatusNotFound)
		return
	}

	updatedTodoList := slices.DeleteFunc(TodoList, func(todo Todo) bool {
		return todo.Id == intId
	})

	SendJSON(w, updatedTodoList)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var body Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	if err != nil {
		SendError(w, r, "Title is required!", http.StatusBadRequest)
		return
	}

	body.Completed = false
	body.Id = TodoList[len(TodoList)-1].Id + 1

	TodoList := append(TodoList, body)

	SendJSON(w, TodoList)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		SendError(w, r, "Invalid param id", http.StatusNotFound)
		return
	}

	if intId > len(TodoList) {
		SendError(w, r, "Todo does not exist with given id", http.StatusBadRequest)
		return
	}

	var body Todo
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	idx := slices.IndexFunc(TodoList, func(todo Todo) bool {
		return todo.Id == intId
	})

	if body.Title != "" {
		TodoList[idx].Title = body.Title
	}

	TodoList[idx].Completed = body.Completed

	SendJSON(w, TodoList[idx])
}

// Middlewares
// func HeadersMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		next.ServeHTTP(w, r)
// 	})

// }

func HeadersMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}

	return http.HandlerFunc(handler)

}

// Utils
func SendError(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	w.WriteHeader(statusCode)

	SendJSON(w, Text{Message: message})
}

func SendJSON(w http.ResponseWriter, data any) {
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", HeadersMiddleware(RootHandler))
	mux.Handle("GET /todos", HeadersMiddleware(GetTodosHandler))
	mux.Handle("POST /todos", HeadersMiddleware(CreateTodoHandler))
	mux.Handle("GET /todos/{id}", HeadersMiddleware(GetTodoHandler))
	mux.Handle("PATCH /todos/{id}", HeadersMiddleware(UpdateTodoHandler))
	mux.Handle("DELETE /todos/{id}", HeadersMiddleware(DeleteTodoHandler))

	fmt.Println("Server is running on port:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server is failed to run", err)
	}
}
