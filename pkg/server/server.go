package server

import (
	"github.com/gorilla/mux"
	"github.com/rugern/go-workshop/pkg/database"
)

func MakeRouter(todoDB database.TodoDB) *mux.Router {
	router := mux.NewRouter()
	todoController := NewTodoController(todoDB)

	// FIXME: Add route to retrieve a todo by ID

	// FIXME: Add route to get all todos

	// FIXME: Add route to add a todo with description

	router.HandleFunc("/todo/{id:[0-9]+}", todoController.UpdateTodo).Methods("PATCH")
	return router
}
