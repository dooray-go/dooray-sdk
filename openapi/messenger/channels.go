package messenger

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	model "github.com/dooray-go/dooray-sdk/openapi/model/messenger"
)

const (
	IDTypeEmail    = "email"
	IDTypeMemberID = "member-id"
)

func (m Messenger) GetChannels(apikey string) (*model.ListChannelsResponse, error) {
	return m.GetChannelsCustomHTTPContext(context.Background(), apikey, m.httpClient)
}
func (m Messenger) GetChannelsContext(ctx context.Context, apikey string) (*model.ListChannelsResponse, error) {
	return m.GetChannelsCustomHTTPContext(ctx, apikey, m.httpClient)
}
func (m Messenger) GetChannelsCustomHTTP(apikey string, httpClient *http.Client) (*model.ListChannelsResponse, error) {
	return m.GetChannelsCustomHTTPContext(context.Background(), apikey, httpClient)
}
func (m Messenger) GetChannelsCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client) (*model.ListChannelsResponse, error) {
	var out model.ListChannelsResponse
	body, err := m.doJSON(ctx, httpClient, apikey, http.MethodGet, "/messenger/v1/channels", nil, nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (m Messenger) CreateChannel(apikey, idType string, req model.CreateChannelRequest) (*model.IDResponse, error) {
	return m.CreateChannelCustomHTTPContext(context.Background(), apikey, m.httpClient, idType, req)
}
func (m Messenger) CreateChannelContext(ctx context.Context, apikey, idType string, req model.CreateChannelRequest) (*model.IDResponse, error) {
	return m.CreateChannelCustomHTTPContext(ctx, apikey, m.httpClient, idType, req)
}
func (m Messenger) CreateChannelCustomHTTP(apikey string, httpClient *http.Client, idType string, req model.CreateChannelRequest) (*model.IDResponse, error) {
	return m.CreateChannelCustomHTTPContext(context.Background(), apikey, httpClient, idType, req)
}
func (m Messenger) CreateChannelCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, idType string, req model.CreateChannelRequest) (*model.IDResponse, error) {
	q := url.Values{}
	if idType != "" {
		q.Set("idType", idType)
	}
	var out model.IDResponse
	body, err := m.doJSON(ctx, httpClient, apikey, http.MethodPost, "/messenger/v1/channels", q, req, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (m Messenger) JoinChannelMembers(apikey, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.JoinChannelMembersCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, memberIDs)
}
func (m Messenger) JoinChannelMembersContext(ctx context.Context, apikey, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.JoinChannelMembersCustomHTTPContext(ctx, apikey, m.httpClient, channelID, memberIDs)
}
func (m Messenger) JoinChannelMembersCustomHTTP(apikey string, httpClient *http.Client, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.JoinChannelMembersCustomHTTPContext(context.Background(), apikey, httpClient, channelID, memberIDs)
}
func (m Messenger) JoinChannelMembersCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.headerCall(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/messenger/v1/channels/%s/members/join", channelID), model.MemberIDsRequest{MemberIDs: memberIDs})
}

func (m Messenger) LeaveChannelMembers(apikey, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.LeaveChannelMembersCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, memberIDs)
}
func (m Messenger) LeaveChannelMembersContext(ctx context.Context, apikey, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.LeaveChannelMembersCustomHTTPContext(ctx, apikey, m.httpClient, channelID, memberIDs)
}
func (m Messenger) LeaveChannelMembersCustomHTTP(apikey string, httpClient *http.Client, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.LeaveChannelMembersCustomHTTPContext(context.Background(), apikey, httpClient, channelID, memberIDs)
}
func (m Messenger) LeaveChannelMembersCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID string, memberIDs []string) (*model.HeaderOnlyResponse, error) {
	return m.headerCall(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/messenger/v1/channels/%s/members/leave", channelID), model.MemberIDsRequest{MemberIDs: memberIDs})
}

func (m Messenger) UpdateChannelLog(apikey, channelID, logID, text string) (*model.HeaderOnlyResponse, error) {
	return m.UpdateChannelLogCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, logID, text)
}
func (m Messenger) UpdateChannelLogContext(ctx context.Context, apikey, channelID, logID, text string) (*model.HeaderOnlyResponse, error) {
	return m.UpdateChannelLogCustomHTTPContext(ctx, apikey, m.httpClient, channelID, logID, text)
}
func (m Messenger) UpdateChannelLogCustomHTTP(apikey string, httpClient *http.Client, channelID, logID, text string) (*model.HeaderOnlyResponse, error) {
	return m.UpdateChannelLogCustomHTTPContext(context.Background(), apikey, httpClient, channelID, logID, text)
}
func (m Messenger) UpdateChannelLogCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID, logID, text string) (*model.HeaderOnlyResponse, error) {
	return m.headerCall(ctx, httpClient, apikey, http.MethodPut, fmt.Sprintf("/messenger/v1/channels/%s/logs/%s", channelID, logID), model.TextRequest{Text: text})
}

func (m Messenger) DeleteChannelLog(apikey, channelID, logID string) (*model.HeaderOnlyResponse, error) {
	return m.DeleteChannelLogCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, logID)
}
func (m Messenger) DeleteChannelLogContext(ctx context.Context, apikey, channelID, logID string) (*model.HeaderOnlyResponse, error) {
	return m.DeleteChannelLogCustomHTTPContext(ctx, apikey, m.httpClient, channelID, logID)
}
func (m Messenger) DeleteChannelLogCustomHTTP(apikey string, httpClient *http.Client, channelID, logID string) (*model.HeaderOnlyResponse, error) {
	return m.DeleteChannelLogCustomHTTPContext(context.Background(), apikey, httpClient, channelID, logID)
}
func (m Messenger) DeleteChannelLogCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID, logID string) (*model.HeaderOnlyResponse, error) {
	return m.headerCall(ctx, httpClient, apikey, http.MethodDelete, fmt.Sprintf("/messenger/v1/channels/%s/logs/%s", channelID, logID), nil)
}

func (m Messenger) ReplyChannelLog(apikey, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.ReplyChannelLogCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, logID, text)
}
func (m Messenger) ReplyChannelLogContext(ctx context.Context, apikey, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.ReplyChannelLogCustomHTTPContext(ctx, apikey, m.httpClient, channelID, logID, text)
}
func (m Messenger) ReplyChannelLogCustomHTTP(apikey string, httpClient *http.Client, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.ReplyChannelLogCustomHTTPContext(context.Background(), apikey, httpClient, channelID, logID, text)
}
func (m Messenger) ReplyChannelLogCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.logSendCall(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/messenger/v1/channels/%s/logs/%s/reply", channelID, logID), model.TextRequest{Text: text})
}

func (m Messenger) CreateAndSendThread(apikey, channelID string, req model.CreateAndSendThreadRequest) (*model.SendMessageResponse, error) {
	return m.CreateAndSendThreadCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, req)
}
func (m Messenger) CreateAndSendThreadContext(ctx context.Context, apikey, channelID string, req model.CreateAndSendThreadRequest) (*model.SendMessageResponse, error) {
	return m.CreateAndSendThreadCustomHTTPContext(ctx, apikey, m.httpClient, channelID, req)
}
func (m Messenger) CreateAndSendThreadCustomHTTP(apikey string, httpClient *http.Client, channelID string, req model.CreateAndSendThreadRequest) (*model.SendMessageResponse, error) {
	return m.CreateAndSendThreadCustomHTTPContext(context.Background(), apikey, httpClient, channelID, req)
}
func (m Messenger) CreateAndSendThreadCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID string, req model.CreateAndSendThreadRequest) (*model.SendMessageResponse, error) {
	return m.logSendCall(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/messenger/v1/channels/%s/threads/create-and-send", channelID), req)
}

func (m Messenger) CreateAndSendThreadFromLog(apikey, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.CreateAndSendThreadFromLogCustomHTTPContext(context.Background(), apikey, m.httpClient, channelID, logID, text)
}
func (m Messenger) CreateAndSendThreadFromLogContext(ctx context.Context, apikey, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.CreateAndSendThreadFromLogCustomHTTPContext(ctx, apikey, m.httpClient, channelID, logID, text)
}
func (m Messenger) CreateAndSendThreadFromLogCustomHTTP(apikey string, httpClient *http.Client, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.CreateAndSendThreadFromLogCustomHTTPContext(context.Background(), apikey, httpClient, channelID, logID, text)
}
func (m Messenger) CreateAndSendThreadFromLogCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, channelID, logID, text string) (*model.SendMessageResponse, error) {
	return m.logSendCall(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/messenger/v1/channels/%s/logs/%s/threads/create-and-send", channelID, logID), model.TextRequest{Text: text})
}

func (m Messenger) logSendCall(ctx context.Context, httpClient *http.Client, apikey, method, path string, payload any) (*model.SendMessageResponse, error) {
	var out model.SendMessageResponse
	body, err := m.doJSON(ctx, httpClient, apikey, method, path, nil, payload, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}
