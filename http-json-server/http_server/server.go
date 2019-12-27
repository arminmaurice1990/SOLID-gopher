package http_server

import (
	"encoding/json"
	"http-json-server/logger"
	"http-json-server/internal_services/messages"
	"http-json-server/internal_services/todos"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type http_server struct {
	port           int
	todoservice    todos.TodoService
	messageservice messages.MessageService
	logger.Logger
}

func NewHttpServer(port int, todoservice todos.TodoService, messageservice messages.MessageService, logger logger.Logger) http_server {
	return http_server{
		port:           port,
		todoservice:    todoservice,
		messageservice: messageservice,
		Logger:         logger,
	}
}

func (h http_server) Serve() error {
	http.Handle("/todo", h.GetTodoRouteHandler())
	http.Handle("/todos", h.ListTodosRouteHandler())
	http.Handle("/messages", h.ListMessagesRouteHandler())

	return http.ListenAndServe(strconv.Itoa(h.port), nil)
}


func (ro *http_server) GetTodoRouteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ro.LogDebug("Get Todo Route called")
		reqBody, err := GetRequestBodyBytes(r.Body)
		if err != nil {
			ro.LogError(err.Error(), r.Body)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reqStruct := &todos.GetTodoRequest{}
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

func (ro *http_server) ListTodosRouteHandler() http.HandlerFunc {
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

func (ro *http_server) ListMessagesRouteHandler() http.HandlerFunc {
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

func GetRequestBodyBytes(reader io.Reader) ([]byte, error) {
	reqBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return []byte{}, nil
	}
	return reqBody, nil
}
