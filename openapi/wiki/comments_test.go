package wiki

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dooray-go/dooray-sdk/openapi/model/wiki"
)

func TestCreateAndGetComments(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages/{pageId}/comments", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var req model.CommentRequest
		_ = json.Unmarshal(raw, &req)
		if req.Body.Content != "댓글 내용 작성" {
			t.Errorf("content: %s", req.Body.Content)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":{"id":"3972742540415907781"}}`))
	})
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/comments/{commentId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":{"id":"3950295078642684620","page":{"id":"3521165468947041024"},"createdAt":"2024-12-03T17:51:10+09:00","modifiedAt":"2024-12-03T17:51:10+09:00","creator":{"type":"member","member":{"organizationMemberId":"3521165460461543659","name":"두레이"}},"body":{"mimeType":"text/x-markdown","content":"hello"}}}`))
	})
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/comments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":[{"id":"3950295078642684620","page":{"id":"p"},"createdAt":"2024-12-03T17:51:10+09:00","modifiedAt":"2024-12-03T17:51:10+09:00","creator":{"type":"member","member":{"organizationMemberId":"1","name":"두레이"}},"body":{"mimeType":"text/x-markdown","content":"hello"}}],"totalCount":27}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewWiki(server.URL)
	created, err := client.CreateComment("key", "w", "p", "댓글 내용 작성")
	if err != nil || created.Result.ID != "3972742540415907781" {
		t.Fatalf("CreateComment: %v %#v", err, created)
	}
	one, err := client.GetComment("key", "w", "p", "c")
	if err != nil || one.Result.Body.Content != "hello" {
		t.Fatalf("GetComment: %v %#v", err, one)
	}
	list, err := client.GetComments("key", "w", "p", 0, 20)
	if err != nil || list.TotalCount != 27 {
		t.Fatalf("GetComments: %v %#v", err, list)
	}
}

func TestGetSharedLinks(t *testing.T) {
	var gotQuery string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/shared-links", func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query().Get("valid")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":[{"id":"1","sharedLink":"https://example.dooray.com/share/pages/x","scope":"member","includeDescendants":true}],"totalCount":1}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	valid := true
	resp, err := NewWiki(server.URL).GetSharedLinks("key", "w", "p", 0, 0, &valid)
	if err != nil {
		t.Fatalf("GetSharedLinks: %v", err)
	}
	if gotQuery != "true" {
		t.Errorf("valid query: %s", gotQuery)
	}
	if len(resp.Result) != 1 || resp.Result[0].Scope != "member" {
		t.Errorf("result: %#v", resp.Result)
	}
}

func TestDeleteComment_NotFound(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /wiki/v1/wikis/{wikiId}/pages/{pageId}/comments/{commentId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":false}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	if _, err := NewWiki(server.URL).DeleteComment("key", "w", "p", "c"); err == nil {
		t.Fatal("expected error")
	}
}
