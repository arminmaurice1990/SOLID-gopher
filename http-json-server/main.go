package http_json_server

import (
	"http-json-server/datastore_connectors"
	"http-json-server/http_server"
	"http-json-server/logger"
	"http-json-server/services/messages"
	"http-json-server/services/todos"
	"log"
)

const port = 8080
const debug = true

func Main() {
	db, err := datastore_connectors.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	logHelper := logger.NewLogger(debug)
	todoservice := todos.NewTodoService(db, logHelper)
	messagesservice := messages.NewMessageService(logHelper)
	server := http_server.NewHttpServer(port, todoservice, messagesservice, logHelper)

	log.Fatal(server.Serve())
}
