package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var todos = []string{
	"Commit",
	"Push",
	"Sleep",
}

func MakeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todo/0", GetTodo).Methods("GET")
	router.HandleFunc("/todo/", GetTodos).Methods("GET")
	router.HandleFunc("/todo/0", RemoveFirstTodo).Methods("DELETE")

	return router
}

func GetTodo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, todos[0])
}

func GetTodos(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, todos)
}

func RemoveFirstTodo(w http.ResponseWriter, _ *http.Request) {
	if len(todos) > 0 {
		todos = todos[1:]
	}
	fmt.Fprint(w, todos)
}
