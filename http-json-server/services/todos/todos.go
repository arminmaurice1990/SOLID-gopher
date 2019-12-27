package todos

import (
	"context"
	"database/sql"
	"http-json-server/logger"
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
	logger.Logger
}

func NewTodoService(db *sql.DB, logger logger.Logger) *todoservice {
	return &todoservice{db: db, Logger:logger}
}

func (t *todoservice) GetTodo(ctx context.Context, id string) (todo, error) {
	t.LogInfo("getting todo with id", id)
	retodo := &todo{}
	rows, err := t.db.Query("SELECT * FROM table WHERE id = ?", id)
	if err != nil {
		t.LogError(err.Error(), id)
		return *retodo, err
	}
	err = rows.Scan(retodo)
	if err != nil {
		t.LogError(err.Error(), retodo)
		return *retodo, err
	}
	return *retodo, nil
}

func (t *todoservice) ListTodos(ctx context.Context) ([]todo, error) {
	t.LogInfo("listing all todos")
	retodos := []todo{}
	rows, err := t.db.Query("SELECT * FROM table")
	if err != nil {
		t.LogError(err.Error())
		return retodos, err
	}
	err = rows.Scan(retodos)
	if err != nil {
		t.LogError(err.Error(), retodos)
		return retodos, err
	}
	return retodos, nil
}
