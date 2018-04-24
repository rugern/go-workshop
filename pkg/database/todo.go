package database

import (
	"github.com/rugern/go-workshop/pkg/model"
	"errors"
)

type TodoDB struct {
	todos map[int]model.Todo
}

func (t TodoDB) SelectTodo(id int) (model.Todo, error) {
	// Check if map contains the requested todo
	if _, ok := t.todos[id]; !ok {
		return model.Todo{}, errors.New("Could not find todo")
	}
	return t.todos[id], nil
}
func (t TodoDB) SelectTodos() (map[int]model.Todo, error) {
	if len(t.todos) == 0 {
		return nil, errors.New("No todos in database")
	}
	return t.todos, nil
}
func (t *TodoDB) InsertTodo(description string) model.Todo {
	id := len(t.todos)
	t.todos[id] = model.Todo{ID: id, Description: description, Checked: false}
	return t.todos[id]
}
func (t *TodoDB) PatchTodo(id int, checked bool) (model.Todo, error) {
	// Check if map contains the requested todo
	if _, ok := t.todos[id]; !ok {
		return model.Todo{}, errors.New("Could not find todo")
	}
	todo := t.todos[id]
	todo.Checked = checked
	t.todos[id] = todo
	return todo, nil
}

func NewTodoDB() TodoDB {
	todos := make(map[int]model.Todo)
	return TodoDB{todos}
}
