package server

import (
	"net/http"
	"github.com/rugern/go-workshop/pkg/database"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

type addTodoRequest struct {
	Description string `json:"Description"`
}

type patchTodoRequest struct {
	Checked bool `json:"Checked"`
}

type TodoController struct {
	todoDB database.TodoDB
}

func (t TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	todo, err := t.todoDB.SelectTodo(id)
	encodeResponse(w, todo, err)
}

func (t TodoController) GetTodos(w http.ResponseWriter, _ *http.Request) {
	todos, err := t.todoDB.SelectTodos()
	encodeResponse(w, todos, err)
}

func (t TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var body addTodoRequest
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	todo := t.todoDB.InsertTodo(body.Description)
	encodeResponse(w, todo, nil)
}

func (t TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var body patchTodoRequest
	err = decoder.Decode(&body)
	if err != nil {
		panic(err)
	}

	todo, err := t.todoDB.PatchTodo(id, body.Checked)
	encodeResponse(w, todo, err)
}

func encodeResponse(w http.ResponseWriter, body interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func NewTodoController() TodoController {
	return TodoController{todoDB: database.NewTodoDB()}
}
