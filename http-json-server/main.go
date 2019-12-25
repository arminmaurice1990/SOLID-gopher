package http_json_server

import (
	"http-json-server/http_server"
	"http-json-server/router"
	"http-json-server/services/todos"
	"log"
)

const port = 8080

func Main() {
	todoservice := todos.NewTodoService()
	router := router.NewRouteHandler(todoservice)
	server := http_server.NewHttpServer(port, router)

	log.Fatal(server.Serve())
}