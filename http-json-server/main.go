package http_json_server

import (
	"http-json-server/database"
	"http-json-server/http_server"
	"http-json-server/route_handler"
	"http-json-server/services/todos"
	"log"
)

const port = 8080

func Main() {
	db, err := database.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	todoservice := todos.NewTodoService(db)
	routehandler := route_handler.NewRouteHandler(todoservice)
	server := http_server.NewHttpServer(port, routehandler)

	log.Fatal(server.Serve())
}