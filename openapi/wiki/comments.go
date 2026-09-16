package wiki

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/dooray-go/dooray-sdk/openapi/model/wiki"
)

func (c *Wiki) CreateComment(apikey, wikiID, pageID, content string) (*model.IDResponse, error) {
	return c.CreateCommentCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, content)
}
func (c *Wiki) CreateCommentContext(ctx context.Context, apikey, wikiID, pageID, content string) (*model.IDResponse, error) {
	return c.CreateCommentCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, content)
}
func (c *Wiki) CreateCommentCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, content string) (*model.IDResponse, error) {
	return c.CreateCommentCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, content)
}
func (c *Wiki) CreateCommentCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, content string) (*model.IDResponse, error) {
	var out model.IDResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/comments", wikiID, pageID), nil, model.CommentRequest{Body: model.CommentBody{Content: content}}, &out, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) GetComments(apikey, wikiID, pageID string, page, size int) (*model.ListCommentsResponse, error) {
	return c.GetCommentsCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, page, size)
}
func (c *Wiki) GetCommentsContext(ctx context.Context, apikey, wikiID, pageID string, page, size int) (*model.ListCommentsResponse, error) {
	return c.GetCommentsCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, page, size)
}
func (c *Wiki) GetCommentsCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string, page, size int) (*model.ListCommentsResponse, error) {
	return c.GetCommentsCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, page, size)
}
func (c *Wiki) GetCommentsCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string, page, size int) (*model.ListCommentsResponse, error) {
	var out model.ListCommentsResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/comments", wikiID, pageID), intQuery(page, size), nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) GetComment(apikey, wikiID, pageID, commentID string) (*model.GetCommentResponse, error) {
	return c.GetCommentCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, commentID)
}
func (c *Wiki) GetCommentContext(ctx context.Context, apikey, wikiID, pageID, commentID string) (*model.GetCommentResponse, error) {
	return c.GetCommentCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, commentID)
}
func (c *Wiki) GetCommentCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, commentID string) (*model.GetCommentResponse, error) {
	return c.GetCommentCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, commentID)
}
func (c *Wiki) GetCommentCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, commentID string) (*model.GetCommentResponse, error) {
	var out model.GetCommentResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/comments/%s", wikiID, pageID, commentID), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) UpdateComment(apikey, wikiID, pageID, commentID, content string) (*model.HeaderOnlyResponse, error) {
	return c.UpdateCommentCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, commentID, content)
}
func (c *Wiki) UpdateCommentContext(ctx context.Context, apikey, wikiID, pageID, commentID, content string) (*model.HeaderOnlyResponse, error) {
	return c.UpdateCommentCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, commentID, content)
}
func (c *Wiki) UpdateCommentCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, commentID, content string) (*model.HeaderOnlyResponse, error) {
	return c.UpdateCommentCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, commentID, content)
}
func (c *Wiki) UpdateCommentCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, commentID, content string) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodPut, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/comments/%s", wikiID, pageID, commentID), model.CommentRequest{Body: model.CommentBody{Content: content}})
}

func (c *Wiki) DeleteComment(apikey, wikiID, pageID, commentID string) (*model.HeaderOnlyResponse, error) {
	return c.DeleteCommentCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, commentID)
}
func (c *Wiki) DeleteCommentContext(ctx context.Context, apikey, wikiID, pageID, commentID string) (*model.HeaderOnlyResponse, error) {
	return c.DeleteCommentCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, commentID)
}
func (c *Wiki) DeleteCommentCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, commentID string) (*model.HeaderOnlyResponse, error) {
	return c.DeleteCommentCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, commentID)
}
func (c *Wiki) DeleteCommentCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, commentID string) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodDelete, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/comments/%s", wikiID, pageID, commentID), nil)
}

func (c *Wiki) GetSharedLinks(apikey, wikiID, pageID string, page, size int, valid *bool) (*model.ListSharedLinksResponse, error) {
	return c.GetSharedLinksCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, page, size, valid)
}
func (c *Wiki) GetSharedLinksContext(ctx context.Context, apikey, wikiID, pageID string, page, size int, valid *bool) (*model.ListSharedLinksResponse, error) {
	return c.GetSharedLinksCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, page, size, valid)
}
func (c *Wiki) GetSharedLinksCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string, page, size int, valid *bool) (*model.ListSharedLinksResponse, error) {
	return c.GetSharedLinksCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, page, size, valid)
}
func (c *Wiki) GetSharedLinksCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string, page, size int, valid *bool) (*model.ListSharedLinksResponse, error) {
	q := intQuery(page, size)
	if valid != nil {
		if q == nil {
			q = url.Values{}
		}
		q.Set("valid", strconv.FormatBool(*valid))
	}
	var out model.ListSharedLinksResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/shared-links", wikiID, pageID), q, nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}
