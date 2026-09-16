package wiki

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dooray-go/dooray-sdk/openapi/model/wiki"
)

const listWikisBody = `{
  "header": {"isSuccessful": true, "resultCode": 0, "resultMessage": "Success"},
  "result": [{"id": "100", "project": {"id": "10"}, "name": "Dooray-공지사항", "type": "public", "scope": "public", "home": {"pageId": "1001"}}],
  "totalCount": 1
}`

const pageDetailBody = `{
    "header": {"isSuccessful": true, "resultMessage": "", "resultCode": 0},
    "result": {
        "id": "100",
        "wikiId": "1",
        "version": "2",
        "parentPageId": "10",
        "subject": "공지사항",
        "body": {"mimeType": "text/x-markdown", "content": "위키 본문 내용"},
        "root": true,
        "createdAt": "2023-04-21T15:47:14+09:00",
        "updatedAt": "2023-09-14T12:26:19+09:00",
        "creator": {"type": "member", "member": {"organizationMemberId": "3521165460461543659"}},
        "referrers": [{"type": "member", "member": {"organizationMemberId": "1"}}],
        "files": [{"id": "f1", "name": "test.xlsx", "size": 5167119, "attachFileId": "a1"}],
        "images": [{"id": "i1", "name": "img.png", "size": 52911, "attachFileId": "a2"}]
    }
}`

func TestGetWikis(t *testing.T) {
	var gotPath, gotQuery string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis", func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(listWikisBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetWikisWithOptions("key", 1, 20)
	if err != nil {
		t.Fatalf("GetWikis: %v", err)
	}
	if gotPath != "/wiki/v1/wikis" {
		t.Errorf("path: %s", gotPath)
	}
	if gotQuery != "page=1&size=20" && gotQuery != "size=20&page=1" {
		t.Errorf("query: %s", gotQuery)
	}
	if len(resp.Result) != 1 || resp.Result[0].Home.PageID != "1001" || resp.TotalCount != 1 {
		t.Errorf("result: %#v", resp.Result)
	}
}

func TestGetPage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/pages/{pageId}", func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("pageId") != "100" {
			t.Errorf("pageId: %s", r.PathValue("pageId"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(pageDetailBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetPage("key", "100")
	if err != nil {
		t.Fatalf("GetPage: %v", err)
	}
	if resp.Result.Subject != "공지사항" || resp.Result.Body.Content != "위키 본문 내용" {
		t.Errorf("page: %#v", resp.Result)
	}
	if resp.Result.Version.String() != "2" {
		t.Errorf("version string: %s", resp.Result.Version)
	}
	if len(resp.Result.Files) != 1 || resp.Result.Files[0].AttachFileID != "a1" {
		t.Errorf("files: %#v", resp.Result.Files)
	}
}

func TestCreatePage_Accepts201(t *testing.T) {
	var gotSubject string
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		var req model.CreatePageRequest
		_ = json.Unmarshal(raw, &req)
		gotSubject = req.Subject
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":{"id":"100","wikiId":"1","parentPageId":"10","version":2}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).CreatePage("key", "wiki-1", model.CreatePageRequest{
		ParentPageID: "10",
		Subject:      "두레이 사용법",
		Body:         model.Body{MimeType: "text/x-markdown", Content: "위키 본문 내용"},
	})
	if err != nil {
		t.Fatalf("CreatePage: %v", err)
	}
	if gotSubject != "두레이 사용법" || resp.Result.ID != "100" || resp.Result.Version != 2 {
		t.Errorf("got subject=%s result=%#v", gotSubject, resp.Result)
	}
}

func TestGetPages_OmitsEmptyParent(t *testing.T) {
	var gotQuery string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages", func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":[{"id":"100","wikiId":"1","version":"2","parentPageId":"10","subject":"공지사항","root":true,"creator":{"type":"member","member":{"organizationMemberId":"1"}}}]}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetPages("key", "wiki-1", "")
	if err != nil {
		t.Fatalf("GetPages: %v", err)
	}
	if gotQuery != "" {
		t.Errorf("expected no query, got %s", gotQuery)
	}
	if len(resp.Result) != 1 || resp.Result[0].Subject != "공지사항" {
		t.Errorf("result: %#v", resp.Result)
	}
}

func TestGetPages_ParsesTotalCount(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveGetPagesBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetPages("key", "wiki-1", "")
	if err != nil {
		t.Fatal(err)
	}
	if resp.TotalCount != 1 {
		t.Errorf("totalCount: %d", resp.TotalCount)
	}
}

func TestGetWikiPage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(pageDetailBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetWikiPage("key", "1", "100")
	if err != nil {
		t.Fatalf("GetWikiPage: %v", err)
	}
	if resp.Result.ID != "100" || resp.Result.Body.MimeType != "text/x-markdown" {
		t.Errorf("result: %#v", resp.Result)
	}
}

func TestUpdatePageTitleAndDeletePage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /wiki/v1/wikis/{wikiId}/pages/{pageId}/title", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		if string(raw) != `{"subject":"두레이 사용법"}` {
			t.Errorf("body: %s", raw)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":null}`))
	})
	mux.HandleFunc("DELETE /wiki/v1/wikis/{wikiId}/pages/{pageId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":null}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewWiki(server.URL)
	title, err := client.UpdatePageTitle("key", "w", "p", "두레이 사용법")
	if err != nil || !title.Header.IsSuccessful {
		t.Fatalf("UpdatePageTitle: %v %#v", err, title)
	}
	del, err := client.DeletePage("key", "w", "p")
	if err != nil || !del.Header.IsSuccessful {
		t.Fatalf("DeletePage: %v %#v", err, del)
	}
}

func TestUpdatePageContentAndReferrers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /wiki/v1/wikis/{wikiId}/pages/{pageId}/content", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":null}`))
	})
	mux.HandleFunc("PUT /wiki/v1/wikis/{wikiId}/pages/{pageId}/referrers", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":null}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewWiki(server.URL)
	if _, err := client.UpdatePageContent("key", "w", "p", model.Body{MimeType: "text/x-markdown", Content: "본문"}); err != nil {
		t.Fatalf("UpdatePageContent: %v", err)
	}
	if _, err := client.UpdatePageReferrers("key", "w", "p", []model.Actor{{Type: "member", Member: model.Member{OrganizationMemberID: "1"}}}); err != nil {
		t.Fatalf("UpdatePageReferrers: %v", err)
	}
}

func TestUpdatePage_NotFound(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /wiki/v1/wikis/{wikiId}/pages/{pageId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":false,"resultCode":404,"resultMessage":"not found"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewWiki(server.URL).UpdatePage("key", "w", "p", model.UpdatePageRequest{Subject: "x", Body: model.Body{Content: "y"}})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMovePage(t *testing.T) {
	var got model.MovePageRequest
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages/{pageId}/move", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &got)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":"Success"},"result":null}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	children := true
	resp, err := NewWiki(server.URL).MovePage("key", "w", "p", model.MovePageRequest{
		TargetParentPageID: "parent",
		WithChildren:       &children,
		BeforePageID:       "0",
	})
	if err != nil {
		t.Fatalf("MovePage: %v", err)
	}
	if got.TargetParentPageID != "parent" || got.BeforePageID != "0" {
		t.Errorf("payload: %#v", got)
	}
	if !resp.Header.IsSuccessful {
		t.Errorf("header: %#v", resp.Header)
	}
}
