package server

import (
	"github.com/gorilla/mux"
	"github.com/rugern/go-workshop/pkg/database"
)


func MakeRouter(todoDB database.TodoDB) *mux.Router {
	router := mux.NewRouter()
	todoController := NewTodoController(todoDB)
	router.HandleFunc("/todo/{id:[0-9]+}", todoController.GetTodo).Methods("GET")

	// FIXME: Get all todos

	// FIXME: Add todo

	// FIXME: Patch todo

	return router
}
