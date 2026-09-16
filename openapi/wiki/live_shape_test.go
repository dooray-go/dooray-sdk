package wiki

import (
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dooray-go/dooray-sdk/openapi/model/wiki"
)

// Fixtures are truncated copies of live api.dooray.com responses from TEST wiki.

const liveGetWikisBody = `{
  "header": {"resultCode": 0, "resultMessage": "", "isSuccessful": true},
  "result": [{
    "name": "TEST",
    "type": "public",
    "scope": "private",
    "home": {"pageId": "3617610205741988427"},
    "project": {"id": "3617610203123698120", "projectCategoryId": null},
    "id": "3617610205158345977"
  }],
  "totalCount": 5
}`

const liveGetPagesBody = `{
  "header": {"resultCode": 0, "resultMessage": "", "isSuccessful": true},
  "result": [{
    "id": "3617610205741988427",
    "wikiId": "3617610205158345977",
    "version": 0,
    "root": true,
    "creator": {"type": "member", "member": {"organizationMemberId": "1900555447797833080"}},
    "subject": "Home"
  }],
  "totalCount": 1
}`

const liveGetPageBody = `{
  "header": {"isSuccessful": true, "resultMessage": "", "resultCode": 0},
  "result": {
    "id": "3617610205741988427",
    "wikiId": "3617610205158345977",
    "version": 0,
    "parentPageId": null,
    "subject": "Home",
    "body": {"mimeType": "text/x-markdown", "content": "# Home"},
    "root": true,
    "createdAt": "2023-09-01T17:25:42+09:00",
    "updatedAt": "2023-09-01T17:25:42+09:00",
    "creator": {"type": "member", "member": {"organizationMemberId": "1900555447797833080"}},
    "referrers": [],
    "files": [],
    "images": []
  }
}`

const liveCreatePageBody = `{
  "header": {"resultCode": 0, "resultMessage": "", "isSuccessful": true},
  "result": {
    "id": "4422720094904263399",
    "wikiId": "3617610205158345977",
    "parentPageId": "3617610205741988427",
    "version": 0
  }
}`

const liveGetCommentsBody = `{
  "header": {"resultCode": 0, "resultMessage": "", "isSuccessful": true},
  "result": [{
    "id": "4422725310956091620",
    "page": {"id": "4422720094904263399"},
    "createdAt": "2026-09-16T13:45:42+09:00",
    "modifiedAt": "2026-09-16T13:45:42+09:00",
    "creator": {"type": "member", "member": {"organizationMemberId": "1900555447797833080", "name": "정지범"}},
    "body": {"mimeType": "text/x-markdown", "content": "sdk-inspect-comment"}
  }],
  "totalCount": 1
}`

const liveSharedLinksBody = `{
  "header": {"resultCode": 0, "resultMessage": "", "isSuccessful": true},
  "result": [{
    "scope": "member",
    "includeDescendants": false,
    "id": "4422719802240581747",
    "sharedLink": "https://manty.dooray.com/share/pages/mftqK0qISL-8p3a47Q_Nig"
  }],
  "totalCount": 1
}`

const liveMovePageBody = `{
  "header": {"isSuccessful": true, "resultMessage": "", "resultCode": 0},
  "result": {
    "wikiId": "3617610205158345977",
    "pageId": "4422720094904263399",
    "parentPageId": "3617610205741988427",
    "version": 3
  }
}`

const liveUploadFileBody = `{
  "header": {"isSuccessful": true, "resultCode": 0, "resultMessage": ""},
  "result": {
    "id": "4422726371991293556",
    "pageFileId": "4422726371991293556",
    "attachFileId": "4422726371981396325",
    "name": "sdk-inspect.txt",
    "mimeType": "application/octet-stream",
    "type": "general",
    "size": 14,
    "createdAt": "2026-09-16T13:47:48.682521151+09:00",
    "extension": "txt"
  }
}`

func TestGetWikis_LiveShape(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveGetWikisBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetWikis("key")
	if err != nil {
		t.Fatal(err)
	}
	if resp.TotalCount != 5 || len(resp.Result) != 1 {
		t.Fatalf("count: total=%d n=%d", resp.TotalCount, len(resp.Result))
	}
	w := resp.Result[0]
	if w.ID != "3617610205158345977" || w.Name != "TEST" || w.Type != "public" || w.Scope != "private" {
		t.Errorf("wiki: %#v", w)
	}
	if w.Home.PageID != "3617610205741988427" || w.Project.ID != "3617610203123698120" {
		t.Errorf("home/project: %#v", w)
	}
}

func TestGetPages_LiveShape(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveGetPagesBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetPages("key", "3617610205158345977", "")
	if err != nil {
		t.Fatal(err)
	}
	if resp.TotalCount != 1 || resp.Result[0].Subject != "Home" || resp.Result[0].Version.String() != "0" {
		t.Errorf("pages: %#v", resp)
	}
	if resp.Result[0].ParentPageID != "" {
		t.Errorf("omitted parentPageId should be empty, got %q", resp.Result[0].ParentPageID)
	}
}

func TestGetPage_LiveShapeNullParent(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/pages/{pageId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveGetPageBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetPage("key", "3617610205741988427")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result.Body.Content != "# Home" || resp.Result.Body.MimeType != "text/x-markdown" {
		t.Errorf("body: %#v", resp.Result.Body)
	}
	if resp.Result.ParentPageID != "" {
		t.Errorf("null parentPageId: %q", resp.Result.ParentPageID)
	}
	if !resp.Result.Root || resp.Result.Version.String() != "0" {
		t.Errorf("page: %#v", resp.Result)
	}
}

func TestCreatePage_Live200VersionZero(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveCreatePageBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).CreatePage("key", "3617610205158345977", model.CreatePageRequest{
		ParentPageID: "3617610205741988427",
		Subject:      "sdk-inspect-wiki-create 13:35:20",
		Body:         model.Body{MimeType: "text/x-markdown", Content: "sdk live create body"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result.ID != "4422720094904263399" || resp.Result.Version != 0 {
		t.Errorf("create: %#v", resp.Result)
	}
}

func TestGetComments_LiveShape(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/comments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveGetCommentsBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetComments("key", "w", "p", 0, 20)
	if err != nil {
		t.Fatal(err)
	}
	if resp.TotalCount != 1 || resp.Result[0].Body.Content != "sdk-inspect-comment" {
		t.Errorf("comments: %#v", resp)
	}
	if resp.Result[0].Creator.Member.Name != "정지범" {
		t.Errorf("creator: %#v", resp.Result[0].Creator)
	}
}

func TestGetComment_LiveNotFoundBody(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/comments/{commentId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"header":{"resultMessage":"null"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewWiki(server.URL).GetComment("key", "w", "p", "4422725310956091620")
	if err == nil {
		t.Fatal("live GetComment on a newly created comment returned 404")
	}
}

func TestGetSharedLinks_LiveShape(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/shared-links", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveSharedLinksBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).GetSharedLinks("key", "w", "p", 0, 0, nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.TotalCount != 1 || resp.Result[0].IncludeDescendants || resp.Result[0].Scope != "member" {
		t.Errorf("shared: %#v", resp.Result)
	}
}

func TestMovePage_LiveResultObject(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages/{pageId}/move", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveMovePageBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).MovePage("key", "3617610205158345977", "4422720094904263399", model.MovePageRequest{
		TargetParentPageID: "3617610205741988427",
		BeforePageID:       "0",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result.Version != 3 || resp.Result.PageID != "4422720094904263399" {
		t.Errorf("move: %#v", resp.Result)
	}
}

func TestUploadFile_LiveExtraFields(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages/{pageId}/files", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveUploadFileBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).UploadPageFile("key", "w", "p", FileTypeGeneral, "sdk-inspect.txt", []byte("sdk live file\n"))
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result.PageFileID != "4422726371991293556" || resp.Result.Extension != "txt" || resp.Result.Size != 14 {
		t.Errorf("upload: %#v", resp.Result)
	}
	if resp.Result.AttachFileID != "4422726371981396325" {
		t.Errorf("attachFileId: %s", resp.Result.AttachFileID)
	}
}
