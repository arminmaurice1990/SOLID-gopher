package http_server

import (
	"http-json-server/router"
	"net/http"
	"strconv"
)

type http_server struct {
	port int
	router router.Router
}

func NewHttpServer(port int, router router.Router) http_server {
	return http_server{port:port, router:router}
}

func (h http_server) Serve () error {
	http.Handle("/todo", h.router.GetTodoRouteHandler())
	http.Handle("/todos", h.router.ListTodosRouteHandler())

	return http.ListenAndServe(strconv.Itoa(h.port), nil)
}