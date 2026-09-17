package messenger

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/dooray-go/dooray-sdk/openapi/model/messenger"
)

type SendMessageRequest struct {
	Text string `json:"text"`
}

func (m Messenger) SendMessage(apikey string, channelID string, msg *SendMessageRequest) (*model.SendMessageResponse, error) {
	return m.SendMessageCustomHTTPContext(context.Background(), apikey, channelID, m.httpClient, msg)
}

func (m Messenger) SendMessageContext(ctx context.Context, apikey string, channelID string, msg *SendMessageRequest) (*model.SendMessageResponse, error) {
	return m.SendMessageCustomHTTPContext(ctx, apikey, channelID, m.httpClient, msg)
}

func (m Messenger) SendMessageCustomHTTP(apikey string, channelID string, httpClient *http.Client, msg *SendMessageRequest) (*model.SendMessageResponse, error) {
	return m.SendMessageCustomHTTPContext(context.Background(), apikey, channelID, httpClient, msg)
}

func (m Messenger) SendMessageCustomHTTPContext(ctx context.Context, apikey string, channelID string, httpClient *http.Client, msg *SendMessageRequest) (*model.SendMessageResponse, error) {
	return m.logSendCall(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/messenger/v1/channels/%s/logs", channelID), msg)
}
