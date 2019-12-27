package messages

import (
	"context"
	"http-json-server/datastore_services/blob_service"
	"http-json-server/logger"
	"time"
)

type message struct {
	key  string
	text []byte
}

type MessageService interface {
	ListMessages(ctx context.Context) ([]message, error)
}

type messageservice struct {
	logger.Logger
	blobservice blob_service.BlobService
}

func NewMessageService(blobservice blob_service.BlobService, logger logger.Logger) *messageservice {
	return &messageservice{Logger: logger, blobservice:blobservice}
}

func (m *messageservice) ListMessages(ctx context.Context) ([]message, error) {
	m.LogInfo("Getting all messages at time: ", time.Now())
	messagemap, err := m.blobservice.ListBlobs(ctx)
	if err != nil {
		m.LogError(err.Error())
		return nil, err
	}
	messages := make([]message, len(messagemap))
	index := 0
	for k, v := range messagemap {
		messages[index] =  message{key:k, text:v}
		index++
	}
	return messages, nil
}
