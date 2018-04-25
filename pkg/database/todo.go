package database

import (
	"database/sql"
	"log"
	"github.com/rugern/go-workshop/pkg/model"
	"net/http"
)

type TodoDB struct {
	db                  *sql.DB
	selectTodoStatement *sql.Stmt
}

func (t TodoDB) SelectTodo(id int) (model.Todo, error) {
	var (
		description string
		checked     bool
	)
	err := t.selectTodoStatement.QueryRow(id).Scan(&description, &checked)
	if err == sql.ErrNoRows {
		return model.Todo{}, model.CustomError{Message: "Could not find todo", StatusCode: http.StatusNotFound}
	}
	if err != nil {
		log.Fatalf("Failed to prepare select todo by ID: %v", err)
		return model.Todo{}, model.CustomError{Message: "Internal server error", StatusCode: http.StatusInternalServerError}
	}

	return model.Todo{ID: id, Description: description, Checked: checked}, nil
}

func NewTodoDB(db *sql.DB) TodoDB {
	selectTodoStatement, err := db.Prepare(`
		SELECT description, checked 
		FROM todos
		WHERE id=?
	`)
	if err != nil {
		log.Fatalf("Failed to prepare select todo by ID: %v", err)
		panic(err)
	}

	return TodoDB{
		db:                  db,
		selectTodoStatement: selectTodoStatement,
	}
}
