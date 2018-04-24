package server

import (
	"github.com/gorilla/mux"
	"github.com/rugern/go-workshop/pkg/database"
)

func MakeRouter(todoDB database.TodoDB) *mux.Router {
	router := mux.NewRouter()
	todoController := NewTodoController(todoDB)
	router.HandleFunc("/todo/{id:[0-9]+}", todoController.GetTodo).Methods("GET")
	router.HandleFunc("/todo/", todoController.GetTodos).Methods("GET")
	router.HandleFunc("/todo/", todoController.AddTodo).Methods("POST")
	router.HandleFunc("/todo/{id:[0-9]+}", todoController.UpdateTodo).Methods("PATCH")
	return router
}
