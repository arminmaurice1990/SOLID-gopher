package http_json_server

import (
	"http-json-server/http_server"
	"http-json-server/route_handler"
	"http-json-server/services/todos"
	"log"
)

const port = 8080

func Main() {
	todoservice := todos.NewTodoService()
	routehandler := route_handler.NewRouteHandler(todoservice)
	server := http_server.NewHttpServer(port, routehandler)

	log.Fatal(server.Serve())
}