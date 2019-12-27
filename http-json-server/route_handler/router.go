package route_handler

import (
	"encoding/json"
	"http-json-server/logger"
	"http-json-server/services/messages"
	"http-json-server/services/todos"
	"io"
	"io/ioutil"
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
	logger.Logger
}

func NewRouteHandler(todoservice todos.TodoService, messageservice messages.MessageService, logger logger.Logger) *routehandler {
	return &routehandler{todoservice: todoservice, messageservice: messageservice, Logger:logger}
}

//request structs
type GetTodoRequest struct {
	id string
}

func (ro *routehandler) GetTodoRouteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ro.LogDebug("Get Todo Route called")
		reqBody, err := GetRequestBodyBytes(r.Body)
		if err != nil {
			ro.LogError(err.Error(), r.Body)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reqStruct := &GetTodoRequest{}
		err = json.Unmarshal(reqBody, reqStruct)
		if err != nil {
			ro.LogError(err.Error(), r.Body)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todo, err := ro.todoservice.GetTodo(r.Context(), "id")
		if err != nil {
			ro.LogError(err.Error(), r.Body)
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		ro.LogDebug("Get Todos Route called")
		todos, err := ro.todoservice.ListTodos(r.Context())
		if err != nil {
			ro.LogError(err.Error(), r.Body)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(todos)
		if err != nil {
			ro.LogError(err.Error(), todos)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}

func (ro *routehandler) ListMessagesRouteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ro.LogDebug("Get Messages Route called")
		messages, err := ro.messageservice.ListMessages(r.Context())
		if err != nil {
			ro.LogError(err.Error(), r.Body)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(messages)
		if err != nil {
			ro.LogError(err.Error(), messages)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
	})
}


func GetRequestBodyBytes(reader io.Reader) ([]byte,error) {
	reqBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return []byte{}, nil
	}
	return reqBody, nil
}