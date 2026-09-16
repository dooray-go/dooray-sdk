package wiki

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Wiki) doJSON(ctx context.Context, httpClient *http.Client, apikey, method, path string, query url.Values, payload any, dest any, okStatuses ...int) ([]byte, error) {
	if httpClient == nil {
		httpClient = c.httpClient
	}
	var body io.Reader
	if payload != nil {
		raw, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
		body = bytes.NewBuffer(raw)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.endPoint+path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "dooray-api "+apikey)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	ok := okStatuses
	if len(ok) == 0 {
		ok = []int{http.StatusOK}
	}
	if !statusOK(resp.StatusCode, ok) {
		return nil, fmt.Errorf("request failed, status: %s, body: %s", resp.Status, string(resBody))
	}

	if dest != nil && len(resBody) > 0 {
		if err := json.Unmarshal(resBody, dest); err != nil {
			return resBody, fmt.Errorf("failed to unmarshal response body: %w", err)
		}
	}
	return resBody, nil
}

func clientWithAuthRedirect(base *http.Client) *http.Client {
	if base == nil {
		base = http.DefaultClient
	}
	c := *base
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 5 {
			return fmt.Errorf("stopped after 5 redirects")
		}
		if auth := via[0].Header.Get("Authorization"); auth != "" {
			req.Header.Set("Authorization", auth)
		}
		return nil
	}
	return &c
}

func statusOK(code int, allowed []int) bool {
	for _, a := range allowed {
		if code == a {
			return true
		}
	}
	return false
}

func intQuery(page, size int) url.Values {
	q := url.Values{}
	if page > 0 {
		q.Set("page", fmt.Sprintf("%d", page))
	}
	if size > 0 {
		q.Set("size", fmt.Sprintf("%d", size))
	}
	return q
}
