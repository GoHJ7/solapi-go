package solapi

import (
	"github.com/solapi/solapi-go/apirequest"
	"github.com/solapi/solapi-go/cash"
	"github.com/solapi/solapi-go/messages"
	"github.com/solapi/solapi-go/storage"
)

// Client struct
type Client struct {
	Messages messages.Messages
	Storage  storage.Storage
	Cash     cash.Cash
}

// NewClient return a new client
func NewClient(config apirequest.ApiRequester) *Client {
	client := Client{}
	client.Messages.ApiRequester = config
	client.Storage.ApiRequester = config
	client.Cash.ApiRequester = config
	return &client
}
