Lightweight TODO REST API written in Go. This repo implements a small JSON HTTP API for creating, reading, updating and deleting todo items. It's intended as a minimal, learning-focused example of building a simple web service in Go.

## Project layout

Top-level files

- `go.mod` - module definition and Go version
- `main.go` - application entry point, routes and server startup

Folders

- `db/` - simple data store (in-memory or file-backed helpers)
- `handlers/` - HTTP handler functions for each endpoint
- `middlewares/` - HTTP middlewares (CORS/headers etc.)
- `types/` - shared types (response envelope, todo model)
- `utils/` - small helper functions for JSON/error responses

## API Endpoints

All responses use `application/json` and follow a envelope with `Success`, `Message`, and `Data` fields.

- `GET /` - root health/info handler
- `GET /todos` - return all todos
- `POST /todos` - create a new todo (JSON body)
- `GET /todos/{id}` - get todo by id
- `PATCH /todos/{id}` - update an existing todo (JSON body)
- `DELETE /todos/{id}` - delete a todo by id

Example request/response (create):

Request:

```json
{
  "title": "Buy milk",
  "completed": false
}
```

Response envelope:

```json
{
  "Success": true,
  "Message": "Success",
  "Data": {
    "id": "1",
    "title": "Buy milk",
    "completed": false
  }
}
```

## Run locally

Ensure you have Go installed (1.20+ recommended). From the project root:

```powershell
go run .\main.go
```

The server listens on `:8080` by default.
