package messenger

import "github.com/dooray-go/dooray-sdk/openapi/model"

type SendMessageResult struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelId"`
}

type SendMessageResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  SendMessageResult    `json:"result"`
	RawJSON string               `json:"-"`
}
