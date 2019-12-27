package http_json_server

import (
	"http-json-server/datastore_connectors"
	"http-json-server/http_server"
	"http-json-server/route_handler"
	"http-json-server/services/messages"
	"http-json-server/services/todos"
	"log"
)

const port = 8080

func Main() {
	db, err := datastore_connectors.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	todoservice := todos.NewTodoService(db)
	messagesservice := messages.NewMessageService()
	routehandler := route_handler.NewRouteHandler(todoservice, messagesservice)
	server := http_server.NewHttpServer(port, routehandler)

	log.Fatal(server.Serve())
}
