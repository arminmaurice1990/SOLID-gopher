package http_json_server

import (
	"http-json-server/datastore_services/blob_service"
	"http-json-server/datastore_services/sql_service"
	"http-json-server/http_server"
	"http-json-server/internal_services/messages"
	"http-json-server/internal_services/todos"
	"http-json-server/logger"
	"log"
)

const port = 8080
const debug = true
const database_url = "DATABASE_URL"
const blob_url = "BLOB_URL"
func Main() {
	sqlservice, err := sql_service.NewPostgresService(database_url)
	if err != nil {
		panic(err)
	}
	blobservice := blob_service.NewBlobConnection(blob_url)
	logHelper := logger.NewLogger(debug)
	todoservice := todos.NewTodoService(sqlservice, logHelper)
	messagesservice := messages.NewMessageService(blobservice, logHelper)
	server := http_server.NewHttpServer(port, todoservice, messagesservice, logHelper)

	log.Fatal(server.Serve())
}
