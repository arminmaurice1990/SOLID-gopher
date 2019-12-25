package todos

import "context"

type todo struct {
	name string
	description string
}

type TodoService interface {
	GetTodo(ctx context.Context, id string) todo
	ListTodos(ctx context.Context) []todo
}

type todoservice struct {}

func NewTodoService() TodoService {
	return &todoservice{}
}

func (t *todoservice) GetTodo (ctx context.Context, id string) todo {
	return todo{name:"test", description:"a test"}
}

func (t *todoservice) ListTodos(ctx context.Context) []todo {
	return []todo{{name:"test", description:"a test"}, todo{name:"test2", description:"a test 2"}}
}