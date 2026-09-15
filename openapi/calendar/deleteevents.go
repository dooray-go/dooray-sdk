package calendar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "github.com/dooray-go/dooray-sdk/openapi/model/calendar"
)

func (c *Calendar) DeleteEvent(apikey string, calendarID string, eventID string, req model.DeleteEventRequest) (*model.EventResponse, error) {
	return c.DeleteEventCustomHTTPContext(context.Background(), apikey, c.httpClient, calendarID, eventID, req)
}

func (c *Calendar) DeleteEventContext(ctx context.Context, apikey string, calendarID string, eventID string, req model.DeleteEventRequest) (*model.EventResponse, error) {
	return c.DeleteEventCustomHTTPContext(ctx, apikey, c.httpClient, calendarID, eventID, req)
}

func (c *Calendar) DeleteEventCustomHTTP(apikey string, httpClient *http.Client, calendarID string, eventID string, req model.DeleteEventRequest) (*model.EventResponse, error) {
	return c.DeleteEventCustomHTTPContext(context.Background(), apikey, httpClient, calendarID, eventID, req)
}

func (c *Calendar) DeleteEventCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, calendarID string, eventID string, req model.DeleteEventRequest) (*model.EventResponse, error) {
	url := fmt.Sprintf("%s/calendar/v1/calendars/%s/events/%s/delete", c.endPoint, calendarID, eventID)

	if req.DeleteType == "" {
		req.DeleteType = model.DeleteTypeThis
	}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal delete event request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json;charset=utf-8")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Accept-Charset", "utf-8")
	httpReq.Header.Set("Authorization", "dooray-api "+apikey)

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to delete event, status: %s, body: %s", resp.Status, string(body))
	}

	var eventResponse model.EventResponse
	if err := json.Unmarshal(body, &eventResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	eventResponse.RawJSON = string(body)

	return &eventResponse, nil
}
