package http_server

import (
	"http-json-server/route_handler"
	"net/http"
	"strconv"
)

type http_server struct {
	port   int
	router route_handler.RouteHandler
}

func NewHttpServer(port int, router route_handler.RouteHandler) http_server {
	return http_server{port: port, router: router}
}

func (h http_server) Serve() error {
	http.Handle("/todo", h.router.GetTodoRouteHandler())
	http.Handle("/todos", h.router.ListTodosRouteHandler())
	http.Handle("/messages", h.router.ListMessagesRouteHandler())

	return http.ListenAndServe(strconv.Itoa(h.port), nil)
}
