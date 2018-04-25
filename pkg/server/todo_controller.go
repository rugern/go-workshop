package server

import (
	"net/http"
	"github.com/rugern/go-workshop/pkg/database"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

type updateTodoRequest struct {
	Checked bool `json:"Checked"`
}

type TodoController struct {
	todoDB database.TodoDB
}

func (t TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err) // FIXME: Handle error
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var body updateTodoRequest
	err = decoder.Decode(&body)
	if err != nil {
		panic(err) // FIXME: Handle error
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

func NewTodoController(todoDB database.TodoDB) TodoController {
	return TodoController{todoDB: todoDB}
}
