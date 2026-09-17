package messenger

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dooray-go/dooray-sdk/openapi/model/messenger"
)

const okNull = `{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":null}`

const logSendBody = `{
    "header": {
        "resultCode": 0,
        "resultMessage": "",
        "isSuccessful": true
    },
    "result": {
        "id": "3986497071236383013",
        "channelId": "3986497069711082184"
    }
}`

func TestGetChannels(t *testing.T) {
	body := `{
    "header": {"resultCode": 0, "resultMessage": "", "isSuccessful": true},
    "totalCount": 2,
    "result": [{
        "id": "2138186748229007787",
        "title": "",
        "organization": {"id": "2131218346506734372"},
        "type": "direct",
        "users": {"participants": [
            {"type": "member", "member": {"organizationMemberId": "2138167606271073201"}},
            {"type": "member", "member": {"organizationMemberId": "1136202552584980936"}}
        ]},
        "me": {"type": "member", "member": {"organizationMemberId": "2138167606271073201"}, "role": "member"},
        "capacity": 2,
        "status": "normal",
        "createdAt": "2020-08-16T12:30:00+09:00",
        "updatedAt": "2020-08-25T12:30:00+09:00",
        "displayed": true,
        "role": "member",
        "archivedAt": null
    }]
}`
	mux := http.NewServeMux()
	mux.HandleFunc("GET /messenger/v1/channels", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "dooray-api key" {
			t.Errorf("auth: %s", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewMessenger(server.URL).GetChannels("key")
	if err != nil {
		t.Fatal(err)
	}
	if resp.TotalCount != 2 || len(resp.Result) != 1 || resp.Result[0].Type != "direct" {
		t.Errorf("channels: %#v", resp)
	}
	if resp.Result[0].Me.Role != "member" || !resp.Result[0].Displayed {
		t.Errorf("me/displayed: %#v", resp.Result[0])
	}
}

func TestCreateChannel_SendsIDType(t *testing.T) {
	var gotQuery, gotType, gotCapacity string
	var gotMembers []string
	mux := http.NewServeMux()
	mux.HandleFunc("POST /messenger/v1/channels", func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query().Get("idType")
		raw, _ := io.ReadAll(r.Body)
		var req model.CreateChannelRequest
		_ = json.Unmarshal(raw, &req)
		gotType = req.Type
		gotCapacity = req.Capacity
		gotMembers = req.MemberIDs
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":{"id":"2790449730960545466"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewMessenger(server.URL).CreateChannel("key", IDTypeMemberID, model.CreateChannelRequest{
		Type:      "private",
		Capacity:  "100",
		MemberIDs: []string{"20000000000000"},
		Title:     "Title",
	})
	if err != nil {
		t.Fatal(err)
	}
	if gotQuery != IDTypeMemberID || gotType != "private" || gotCapacity != "100" || len(gotMembers) != 1 {
		t.Errorf("query=%s type=%s capacity=%s members=%v", gotQuery, gotType, gotCapacity, gotMembers)
	}
	if resp.Result.ID != "2790449730960545466" {
		t.Errorf("id: %s", resp.Result.ID)
	}
}

func TestJoinAndLeaveChannelMembers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /messenger/v1/channels/{channelId}/members/join", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(okNull))
	})
	mux.HandleFunc("POST /messenger/v1/channels/{channelId}/members/leave", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(okNull))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewMessenger(server.URL)
	if _, err := client.JoinChannelMembers("key", "1", []string{"2131762672128832968"}); err != nil {
		t.Fatalf("join: %v", err)
	}
	if _, err := client.LeaveChannelMembers("key", "1", []string{"2131762672128832968"}); err != nil {
		t.Fatalf("leave: %v", err)
	}
}

func TestUpdateDeleteReplyChannelLog(t *testing.T) {
	var putText, replyText string
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /messenger/v1/channels/{channelId}/logs/{logId}", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var req model.TextRequest
		_ = json.Unmarshal(raw, &req)
		putText = req.Text
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(okNull))
	})
	mux.HandleFunc("DELETE /messenger/v1/channels/{channelId}/logs/{logId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(okNull))
	})
	mux.HandleFunc("POST /messenger/v1/channels/{channelId}/logs/{logId}/reply", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var req model.TextRequest
		_ = json.Unmarshal(raw, &req)
		replyText = req.Text
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(logSendBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewMessenger(server.URL)
	if _, err := client.UpdateChannelLog("key", "c", "l", "modified text"); err != nil {
		t.Fatal(err)
	}
	if putText != "modified text" {
		t.Errorf("put text: %s", putText)
	}
	if _, err := client.DeleteChannelLog("key", "c", "l"); err != nil {
		t.Fatal(err)
	}
	reply, err := client.ReplyChannelLog("key", "c", "l", "reply")
	if err != nil {
		t.Fatal(err)
	}
	if replyText != "reply" {
		t.Errorf("reply: %s", replyText)
	}
	if reply.Result.ID != "3986497071236383013" || reply.Result.ChannelID != "3986497069711082184" {
		t.Errorf("reply result: %#v", reply.Result)
	}
}

func TestCreateAndSendThread(t *testing.T) {
	var got model.CreateAndSendThreadRequest
	mux := http.NewServeMux()
	mux.HandleFunc("POST /messenger/v1/channels/{channelId}/threads/create-and-send", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &got)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(logSendBody))
	})
	mux.HandleFunc("POST /messenger/v1/channels/{channelId}/logs/{logId}/threads/create-and-send", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(logSendBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewMessenger(server.URL)
	created, err := client.CreateAndSendThread("key", "c", model.CreateAndSendThreadRequest{Text: "root", ThreadText: "thread"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Text != "root" || got.ThreadText != "thread" {
		t.Errorf("thread req: %#v", got)
	}
	if created.Result.ID != "3986497071236383013" || created.Result.ChannelID != "3986497069711082184" {
		t.Errorf("thread result: %#v", created.Result)
	}
	fromLog, err := client.CreateAndSendThreadFromLog("key", "c", "l", "from-log")
	if err != nil {
		t.Fatal(err)
	}
	if fromLog.Result.ID != "3986497071236383013" {
		t.Errorf("from-log result: %#v", fromLog.Result)
	}
}

func TestSendMessage_ReturnsLogID(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /messenger/v1/channels/{channelId}/logs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(logSendBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewMessenger(server.URL).SendMessage("key", "3986497069711082184", &SendMessageRequest{Text: "hi"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result.ID != "3986497071236383013" || resp.Result.ChannelID != "3986497069711082184" {
		t.Errorf("send: %#v", resp.Result)
	}
}

func TestCreateChannel_NotOK(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /messenger/v1/channels", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewMessenger(server.URL).CreateChannel("key", IDTypeEmail, model.CreateChannelRequest{Type: "direct"})
	if err == nil {
		t.Fatal("expected error")
	}
}
