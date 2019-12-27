package messages

import (
	"context"
	"gocloud.dev/blob"
	"http-json-server/datastore_connectors"
)

type message struct {
	text []byte
}

type MessageService interface {
	ListMessages(ctx context.Context) ([]message, error)
}

type messageservice struct{}

func NewMessageService() *messageservice {
	return &messageservice{}
}

func (m *messageservice) ListMessages(ctx context.Context) (messages []message, err error) {
	bucket, err := datastore_connectors.GetBucket(ctx)
	if err != nil {
		return messages, err
	}
	opts := &blob.ListOptions{Prefix: "ITEM", Delimiter: ","}
	messageIter := bucket.List(opts)

	for {
		obj, err := messageIter.Next(ctx)
		if err != nil {
			return messages, err
		}
		message := message{}
		message.text = obj.MD5
		messages = append(messages, message)
	}

	return messages, nil
}
