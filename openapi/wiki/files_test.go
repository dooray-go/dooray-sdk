package wiki

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadPageFile_SendsTypeBeforeFile(t *testing.T) {
	var gotType, gotFileName string
	mux := http.NewServeMux()
	mux.HandleFunc("POST /wiki/v1/wikis/{wikiId}/pages/{pageId}/files", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			t.Errorf("parse: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		gotType = r.FormValue("type")
		_, hdr, err := r.FormFile("file")
		if err != nil {
			t.Errorf("file: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		gotFileName = hdr.Filename
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":""},"result":{"id":"2541304532468051951","attachFileId":"2541304532468051951","name":"jwt.pdf","mimeType":"application/pdf","type":"general","size":1728914,"createdAt":"2019-08-08T16:58:27+09:00"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).UploadPageFile("key", "w", "p", FileTypeGeneral, "jwt.pdf", []byte("pdf"))
	if err != nil {
		t.Fatalf("UploadPageFile: %v", err)
	}
	if gotType != FileTypeGeneral || gotFileName != "jwt.pdf" {
		t.Errorf("type=%s file=%s", gotType, gotFileName)
	}
	if resp.Result.AttachFileID == "" || resp.Result.Size != 1728914 {
		t.Errorf("result: %#v", resp.Result)
	}
}

func TestDownloadAttachFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/attachFiles/{attachFileId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("PDFDATA"))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	dl, err := NewWiki(server.URL).DownloadAttachFile("key", "w", "a1")
	if err != nil {
		t.Fatalf("DownloadAttachFile: %v", err)
	}
	if string(dl.Content) != "PDFDATA" || dl.ContentType != "application/pdf" {
		t.Errorf("download: %#v", dl)
	}
}

func TestUploadPageFile_ForwardsAuthorizationOn307(t *testing.T) {
	var fileAPIAuth string
	fileAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileAPIAuth = r.Header.Get("Authorization")
		if fileAPIAuth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"header":{"resultMessage":"Authorization header is invalid"}}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(liveUploadFileBody))
	}))
	defer fileAPI.Close()

	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, fileAPI.URL+r.URL.Path, http.StatusTemporaryRedirect)
	}))
	defer api.Close()

	resp, err := NewWiki(api.URL).UploadPageFile("key", "w", "p", FileTypeGeneral, "sdk-inspect.txt", []byte("sdk live file\n"))
	if err != nil {
		t.Fatalf("UploadPageFile: %v", err)
	}
	if fileAPIAuth != "dooray-api key" {
		t.Errorf("file API Authorization: %q", fileAPIAuth)
	}
	if resp.Result.Name != "sdk-inspect.txt" {
		t.Errorf("result: %#v", resp.Result)
	}
}

func TestDownloadPageFile_LiveBytes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /wiki/v1/wikis/{wikiId}/pages/{pageId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("sdk live file\n"))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	dl, err := NewWiki(server.URL).DownloadPageFile("key", "w", "p", "f")
	if err != nil {
		t.Fatal(err)
	}
	if string(dl.Content) != "sdk live file\n" || !strings.Contains(dl.ContentType, "octet-stream") {
		t.Errorf("download: %#v", dl)
	}
}

func TestDeletePageFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /wiki/v1/wikis/{wikiId}/pages/{pageId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"resultCode":0,"resultMessage":"","isSuccessful":true},"result":null}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := NewWiki(server.URL).DeletePageFile("key", "w", "p", "f")
	if err != nil || !resp.Header.IsSuccessful {
		t.Fatalf("DeletePageFile: %v %#v", err, resp)
	}
}
