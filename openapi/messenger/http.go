package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	model "github.com/dooray-go/dooray-sdk/openapi/model/messenger"
)

func (m Messenger) doJSON(ctx context.Context, httpClient *http.Client, apikey, method, path string, query url.Values, payload any, dest any) ([]byte, error) {
	if httpClient == nil {
		httpClient = m.httpClient
	}
	var body io.Reader
	if payload != nil {
		raw, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("marshal failed: %w", err)
		}
		body = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, method, m.endPoint+path, body)
	if err != nil {
		return nil, fmt.Errorf("failed new request: %w", err)
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("dooray-api %s", apikey))
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	if err := checkStatusCode(resp); err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if dest != nil && len(resBody) > 0 {
		if err := json.Unmarshal(resBody, dest); err != nil {
			return resBody, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	return resBody, nil
}

func (m Messenger) headerCall(ctx context.Context, httpClient *http.Client, apikey, method, path string, payload any) (*model.HeaderOnlyResponse, error) {
	var out model.HeaderOnlyResponse
	body, err := m.doJSON(ctx, httpClient, apikey, method, path, nil, payload, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}
