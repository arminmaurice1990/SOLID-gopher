package todos

import (
	"context"
	"database/sql"
)

type todo struct {
	name        string
	description string
}

type TodoService interface {
	GetTodo(ctx context.Context, id string) (todo, error)
	ListTodos(ctx context.Context) ([]todo, error)
}

type todoservice struct {
	db *sql.DB
}

func NewTodoService(db *sql.DB) *todoservice {
	return &todoservice{db: db}
}

func (t *todoservice) GetTodo(ctx context.Context, id string) (todo, error) {
	retodo := &todo{}
	rows, err := t.db.Query("SELECT * FROM table WHERE id = ?", id)
	if err != nil {
		return *retodo, err
	}
	err = rows.Scan(retodo)
	if err != nil {
		return *retodo, err
	}
	return *retodo, nil
}

func (t *todoservice) ListTodos(ctx context.Context) ([]todo, error) {
	retodos := []todo{}
	rows, err := t.db.Query("SELECT * FROM table")
	if err != nil {
		return retodos, err
	}
	err = rows.Scan(retodos)
	if err != nil {
		return retodos, err
	}
	return retodos, nil
}
