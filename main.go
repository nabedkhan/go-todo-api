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
	// w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(Text{
		Message: "Go server is running on",
	})
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(TodoList)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		errMessage := Text{
			Message: "Invalid route param",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusBadRequest)
		return
	}

	if intId > len(TodoList) {
		errMessage := Text{
			Message: "Todo does not exist with given id",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusNotFound)
		return
	}

	idx := slices.IndexFunc(TodoList, func(todo Todo) bool {
		return todo.Id == intId
	})

	encoder := json.NewEncoder(w)
	encoder.Encode(TodoList[idx])
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		errMessage := Text{
			Message: "Invalid route param",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusBadRequest)
	}

	if intId > len(TodoList) {
		errMessage := Text{
			Message: "Todo does not exist with given id",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusNotFound)
	}

	updatedTodoList := slices.DeleteFunc(TodoList, func(todo Todo) bool {
		return todo.Id == intId
	})

	encoder := json.NewEncoder(w)
	encoder.Encode(updatedTodoList)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var body Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	if err != nil {
		errMessage := Text{
			Message: "Title is required!",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusBadRequest)
		return
	}

	body.Completed = false
	body.Id = TodoList[len(TodoList)-1].Id + 1

	updatedTodoList := append(TodoList, body)

	encoder := json.NewEncoder(w)
	encoder.Encode(updatedTodoList)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		errMessage := Text{
			Message: "Invalid param id",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusBadRequest)
	}

	if intId > len(TodoList) {
		errMessage := Text{
			Message: "Todo does not exist with given id",
		}

		errJson, _ := json.Marshal(errMessage)
		http.Error(w, string(errJson), http.StatusBadRequest)
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

	encoder := json.NewEncoder(w)
	encoder.Encode(TodoList[idx])
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
