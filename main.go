package main

import (
	"fmt"
	"net/http"

	"github.com/nabedkhan/go-todo-api/handlers"
	"github.com/nabedkhan/go-todo-api/middlewares"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", middlewares.HeadersMiddleware(handlers.RootHandler))
	mux.Handle("GET /todos", middlewares.HeadersMiddleware(handlers.GetTodosHandler))
	mux.Handle("POST /todos", middlewares.HeadersMiddleware(handlers.CreateTodoHandler))
	mux.Handle("GET /todos/{id}", middlewares.HeadersMiddleware(handlers.GetTodoHandler))
	mux.Handle("PATCH /todos/{id}", middlewares.HeadersMiddleware(handlers.UpdateTodoHandler))
	mux.Handle("DELETE /todos/{id}", middlewares.HeadersMiddleware(handlers.DeleteTodoHandler))

	fmt.Println("Server is running on port:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server is failed to run", err)
	}
}
