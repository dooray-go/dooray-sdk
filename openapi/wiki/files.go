package wiki

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	model "github.com/dooray-go/dooray-sdk/openapi/model/wiki"
)

const (
	FileTypeGeneral     = "general"
	FileTypeInlineImage = "inline_image"
)

func (c *Wiki) DownloadAttachFile(apikey, wikiID, attachFileID string) (*model.Download, error) {
	return c.DownloadAttachFileCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, attachFileID)
}
func (c *Wiki) DownloadAttachFileContext(ctx context.Context, apikey, wikiID, attachFileID string) (*model.Download, error) {
	return c.DownloadAttachFileCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, attachFileID)
}
func (c *Wiki) DownloadAttachFileCustomHTTP(apikey string, httpClient *http.Client, wikiID, attachFileID string) (*model.Download, error) {
	return c.DownloadAttachFileCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, attachFileID)
}
func (c *Wiki) DownloadAttachFileCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, attachFileID string) (*model.Download, error) {
	return c.download(ctx, httpClient, apikey, fmt.Sprintf("/wiki/v1/wikis/%s/attachFiles/%s", wikiID, attachFileID))
}

func (c *Wiki) DownloadPageFile(apikey, wikiID, pageID, fileID string) (*model.Download, error) {
	return c.DownloadPageFileCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, fileID)
}
func (c *Wiki) DownloadPageFileContext(ctx context.Context, apikey, wikiID, pageID, fileID string) (*model.Download, error) {
	return c.DownloadPageFileCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, fileID)
}
func (c *Wiki) DownloadPageFileCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, fileID string) (*model.Download, error) {
	return c.DownloadPageFileCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, fileID)
}
func (c *Wiki) DownloadPageFileCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, fileID string) (*model.Download, error) {
	return c.download(ctx, httpClient, apikey, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/files/%s", wikiID, pageID, fileID))
}

func (c *Wiki) DeletePageFile(apikey, wikiID, pageID, fileID string) (*model.HeaderOnlyResponse, error) {
	return c.DeletePageFileCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, fileID)
}
func (c *Wiki) DeletePageFileContext(ctx context.Context, apikey, wikiID, pageID, fileID string) (*model.HeaderOnlyResponse, error) {
	return c.DeletePageFileCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, fileID)
}
func (c *Wiki) DeletePageFileCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, fileID string) (*model.HeaderOnlyResponse, error) {
	return c.DeletePageFileCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, fileID)
}
func (c *Wiki) DeletePageFileCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, fileID string) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodDelete, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/files/%s", wikiID, pageID, fileID), nil)
}

func (c *Wiki) UploadPageFile(apikey, wikiID, pageID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.UploadPageFileCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, fileType, filename, content)
}
func (c *Wiki) UploadPageFileContext(ctx context.Context, apikey, wikiID, pageID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.UploadPageFileCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, fileType, filename, content)
}
func (c *Wiki) UploadPageFileCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.UploadPageFileCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, fileType, filename, content)
}
func (c *Wiki) UploadPageFileCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.upload(ctx, httpClient, apikey, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/files", wikiID, pageID), fileType, filename, content)
}

func (c *Wiki) UploadWikiFile(apikey, wikiID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.UploadWikiFileCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, fileType, filename, content)
}
func (c *Wiki) UploadWikiFileContext(ctx context.Context, apikey, wikiID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.UploadWikiFileCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, fileType, filename, content)
}
func (c *Wiki) UploadWikiFileCustomHTTP(apikey string, httpClient *http.Client, wikiID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	return c.UploadWikiFileCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, fileType, filename, content)
}
func (c *Wiki) UploadWikiFileCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	if fileType == "" {
		fileType = FileTypeGeneral
	}
	return c.upload(ctx, httpClient, apikey, fmt.Sprintf("/wiki/v1/wikis/%s/files", wikiID), fileType, filename, content)
}

func (c *Wiki) download(ctx context.Context, httpClient *http.Client, apikey, path string) (*model.Download, error) {
	if httpClient == nil {
		httpClient = c.httpClient
	}
	httpClient = clientWithAuthRedirect(httpClient)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endPoint+path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "dooray-api "+apikey)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTemporaryRedirect {
		return nil, fmt.Errorf("download failed, status: %s, body: %s", resp.Status, string(raw))
	}
	return &model.Download{Content: raw, ContentType: resp.Header.Get("Content-Type"), StatusCode: resp.StatusCode}, nil
}

func (c *Wiki) upload(ctx context.Context, httpClient *http.Client, apikey, path, fileType, filename string, content []byte) (*model.UploadFileResponse, error) {
	if httpClient == nil {
		httpClient = c.httpClient
	}
	httpClient = clientWithAuthRedirect(httpClient)
	if fileType == "" {
		fileType = FileTypeGeneral
	}
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	// Dooray requires the type field before the file field.
	if err := w.WriteField("type", fileType); err != nil {
		return nil, fmt.Errorf("failed to write type field: %w", err)
	}
	part, err := w.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create file field: %w", err)
	}
	if _, err := part.Write(content); err != nil {
		return nil, fmt.Errorf("failed to write file content: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endPoint+path, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "dooray-api "+apikey)
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusTemporaryRedirect {
		return nil, fmt.Errorf("upload failed, status: %s, body: %s", resp.Status, string(raw))
	}
	var out model.UploadFileResponse
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &out); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
	}
	out.RawJSON = string(raw)
	return &out, nil
}
