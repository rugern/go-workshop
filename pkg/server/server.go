package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func MakeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")
	return router
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
