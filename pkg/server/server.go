package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var todos = []string{
	"Commit",
	"Clean",
	"Sleep",
}

func MakeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todo/0", GetTodo).Methods("GET")

	// FIXME: Create endpoint to return all todos

	// FIXME: Create endpoint to remove the first todo in the list, and return the remaining todos

	return router
}

func GetTodo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, todos[0])
}
