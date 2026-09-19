package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "github.com/dooray-go/dooray-sdk/openapi/model/account"
	"github.com/dooray-go/dooray-sdk/utils"
)

// GetMe returns the member associated with the supplied personal authentication token.
func (a *Account) GetMe(apikey string) (*model.GetMemberResponse, error) {
	return a.GetMeCustomHTTPContext(context.Background(), apikey, a.httpClient)
}

// GetMeContext returns the current member using the supplied context.
func (a *Account) GetMeContext(ctx context.Context, apikey string) (*model.GetMemberResponse, error) {
	return a.GetMeCustomHTTPContext(ctx, apikey, a.httpClient)
}

// GetMeCustomHTTP returns the current member using the supplied HTTP client.
func (a *Account) GetMeCustomHTTP(apikey string, httpClient *http.Client) (*model.GetMemberResponse, error) {
	return a.GetMeCustomHTTPContext(context.Background(), apikey, httpClient)
}

// GetMeCustomHTTPContext returns the current member using the supplied context and HTTP client.
func (a *Account) GetMeCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client) (*model.GetMemberResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.endPoint+"/common/v1/members/me", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "dooray-api "+apikey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get current member: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, utils.StatusCodeError{Code: resp.StatusCode, Status: resp.Status}
	}

	var memberResponse model.GetMemberResponse
	if err := json.Unmarshal(body, &memberResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	memberResponse.RawJSON = string(body)

	return &memberResponse, nil
}
