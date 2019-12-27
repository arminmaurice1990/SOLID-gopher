package route_handler

import (
	"encoding/json"
	"http-json-server/services/messages"
	"http-json-server/services/todos"
	"net/http"
)

type RouteHandler interface {
	GetTodoRouteHandler() http.HandlerFunc
	ListTodosRouteHandler() http.HandlerFunc
	ListMessagesRouteHandler() http.HandlerFunc
}

type routehandler struct {
	todoservice    todos.TodoService
	messageservice messages.MessageService
}

func NewRouteHandler(todoservice todos.TodoService, messageservice messages.MessageService) *routehandler {
	return &routehandler{todoservice: todoservice, messageservice: messageservice}
}

func (ro *routehandler) GetTodoRouteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo, err := ro.todoservice.GetTodo(r.Context(), "id")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}

func (ro *routehandler) ListTodosRouteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todos, err := ro.todoservice.ListTodos(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}

func (ro *routehandler) ListMessagesRouteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo, err := ro.messageservice.ListMessages(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}
