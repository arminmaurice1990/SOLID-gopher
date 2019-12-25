package router

import (
	"encoding/json"
	"http-json-server/services/todos"
	"net/http"
)

type Router interface {
	GetTodoRouteHandler() http.HandlerFunc
	ListTodosRouteHandler() http.HandlerFunc
}

type router struct {
	todoservice todos.TodoService
}

func NewRouteHandler(todoservice todos.TodoService) *router {
	return &router{todoservice:todoservice}
}

func (ro *router) GetTodoRouteHandler () http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo := ro.todoservice.GetTodo(r.Context(), "id")
		jsonBytes, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}

func (ro *router) ListTodosRouteHandler () http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todos := ro.todoservice.ListTodos(r.Context())
		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}