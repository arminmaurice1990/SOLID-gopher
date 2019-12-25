package route_handler

import (
	"encoding/json"
	"http-json-server/services/todos"
	"net/http"
)

type RouteHandler interface {
	GetTodoRouteHandler() http.HandlerFunc
	ListTodosRouteHandler() http.HandlerFunc
}

type routehandler struct {
	todoservice todos.TodoService
}

func NewRouteHandler(todoservice todos.TodoService) *routehandler {
	return &routehandler{todoservice: todoservice}
}

func (ro *routehandler) GetTodoRouteHandler () http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo, err := ro.todoservice.GetTodo(r.Context(), "id")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}

func (ro *routehandler) ListTodosRouteHandler () http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todos, err := ro.todoservice.ListTodos(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}