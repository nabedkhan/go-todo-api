# go-todo-api

A minimal Go HTTP server that exposes a simple in-memory TODO API. It's a small demo project intended to show basic routing and JSON handling in Go.

## Features

- In-memory TODO list (no database)
- JSON API
- Endpoints: list todos, get todo by id, delete todo

## Getting started

1. Clone the repository (or open this folder).
2. Run the server:

```powershell
go run main.go
```

The server listens on port 8080 by default.

## API

Base URL: http://localhost:8080

Content-Type: application/json for all requests and responses.

Endpoints:

- GET /

  - Response: { "message": "Go server is running on" }

- GET /todos

  - Returns the full TODO list (array of todos).

- GET /todos/{id}

  - Returns a single todo by numeric id. Returns 400 for invalid id, 404 if not found.

- DELETE /todos/{id}
  - Deletes the todo with the given id and returns the updated list. Returns 400 for invalid id, 404 if not found.

Example responses:

GET /todos

```json
[
  { "id": 1, "title": "Learning Go", "completed": false },
  { "id": 2, "title": "Learning React", "completed": false }
]
```

GET /todos/1

```json
{ "id": 1, "title": "Learning Go", "completed": false }
```

Error response example (400):

```json
{ "message": "Invalid route param" }
```
