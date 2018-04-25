package server

import (
	"net/http"
	"github.com/rugern/go-workshop/pkg/database"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/rugern/go-workshop/pkg/model"
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
		encodeResponse(w, model.Todo{}, model.CustomError{
			Message: fmt.Sprintf("Invalid parameter, reason: %s", err.Error()),
			StatusCode: http.StatusBadRequest,
		})
	}

	todo, err := t.todoDB.SelectTodo(id)
	encodeResponse(w, todo, err)
}

func encodeResponse(w http.ResponseWriter, body interface{}, err error) {
	if err != nil {
		switch err.(type) {
		case model.CustomError:
			customError := err.(model.CustomError)
			w.WriteHeader(customError.Status())
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
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
