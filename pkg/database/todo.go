package database

import (
	"github.com/rugern/go-workshop/pkg/model"
)

type TodoDB struct {
	// FIXME: Decide on data structure to store todos
}

func (t *TodoDB) PatchTodo(id int, checked bool) (model.Todo, error) {
	// FIXME: Update todo identified by 'id' with the value of 'checked', and return the updated todo
	return model.Todo{}, nil
}

func NewTodoDB() TodoDB {
	return TodoDB{}
}
