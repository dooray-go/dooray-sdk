package project

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "github.com/dooray-go/dooray-sdk/openapi/model/project"
)

func (c *Project) GetPost(apikey string, projectID string, postID string) (*model.GetPostResponse, error) {
	return c.GetPostCustomHTTPContext(context.Background(), apikey, c.httpClient, projectID, postID)
}

func (c *Project) GetPostContext(ctx context.Context, apikey string, projectID string, postID string) (*model.GetPostResponse, error) {
	return c.GetPostCustomHTTPContext(ctx, apikey, c.httpClient, projectID, postID)
}

func (c *Project) GetPostCustomHTTP(apikey string, httpClient *http.Client, projectID string, postID string) (*model.GetPostResponse, error) {
	return c.GetPostCustomHTTPContext(context.Background(), apikey, httpClient, projectID, postID)
}

func (c *Project) GetPostCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, projectID string, postID string) (*model.GetPostResponse, error) {
	url := fmt.Sprintf("%s/project/v1/projects/%s/posts/%s", c.endPoint, projectID, postID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "dooray-api "+apikey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get post, status: %s, body: %s", resp.Status, string(body))
	}

	var postResponse model.GetPostResponse
	if err := json.Unmarshal(body, &postResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	postResponse.RawJSON = string(body)

	return &postResponse, nil
}
