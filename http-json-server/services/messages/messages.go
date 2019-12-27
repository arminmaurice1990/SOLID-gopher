package messages

import (
	"context"
	"gocloud.dev/blob"
	"http-json-server/datastore_connectors"
	"http-json-server/logger"
	"time"
)

type message struct {
	text []byte
}

type MessageService interface {
	ListMessages(ctx context.Context) ([]message, error)
}

type messageservice struct {
	logger.Logger
}

func NewMessageService(logger logger.Logger) *messageservice {
	return &messageservice{logger}
}

func (m *messageservice) ListMessages(ctx context.Context) (messages []message, err error) {
	m.LogInfo("Getting all messages at time: ", time.Now())
	bucket, err := datastore_connectors.GetBucket(ctx)
	if err != nil {
		m.Logger.LogError(err.Error(), bucket)
		return messages, err
	}
	opts := &blob.ListOptions{Prefix: "ITEM", Delimiter: ","}
	messageIter := bucket.List(opts)

	for {
		obj, err := messageIter.Next(ctx)
		if err != nil {
			m.Logger.LogError(err.Error(), obj)
			return messages, err
		}
		message := message{}
		message.text = obj.MD5
		messages = append(messages, message)
	}

	return messages, nil
}
