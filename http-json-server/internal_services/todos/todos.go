package todos

import (
	"context"
	"encoding/json"
	"http-json-server/datastore_services/sql_service"
	"http-json-server/logger"
)


//request structs
type getTodoRequest struct {
	id string
}

//return structs
type todo struct {
	name        string
	description string
}

type TodoService interface {
	GetTodo(ctx context.Context, request getTodoRequest) (todo, error)
	UnmarshalGetTodoRequest(bts []byte) (getTodoRequest, error)
	ListTodos(ctx context.Context) ([]todo, error)
}

type todoservice struct {
	sqlservice sql_service.PostgresService
	logger.Logger
}

func NewTodoService(sqlservice sql_service.PostgresService, logger logger.Logger) *todoservice {
	return &todoservice{sqlservice:sqlservice, Logger: logger}
}

func (t todoservice) UnmarshalGetTodoRequest(bts []byte) (getTodoRequest, error) {
	req := getTodoRequest{}
	err := json.Unmarshal(bts, req)
	if err != nil {
		t.LogError(err.Error(), bts)
		return req, err
	}
	return req, nil
}

func (t *todoservice) GetTodo(ctx context.Context, request getTodoRequest) (todo, error) {
	t.LogInfo("getting todo with id", request.id)
	retodo := &todo{}
	err := t.sqlservice.Query(retodo, "SELECT * FROM table WHERE id = ?", request.id)
	if err != nil {
		t.LogError(err.Error(), retodo)
		return *retodo, err
	}
	return *retodo, nil
}

func (t *todoservice) ListTodos(ctx context.Context) ([]todo, error) {
	t.LogInfo("listing all todos")
	retodos := []todo{}
	err := t.sqlservice.Query(retodos, "SELECT * FROM table")
	if err != nil {
		t.LogError(err.Error(), retodos)
		return retodos, err
	}
	return retodos, nil
}
