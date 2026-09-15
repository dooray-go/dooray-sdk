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

func (c *Calendar) UpdateEvent(apikey string, calendarID string, eventID string, event model.UpdateEventRequest) (*model.EventResponse, error) {
	return c.UpdateEventCustomHTTPContext(context.Background(), apikey, c.httpClient, calendarID, eventID, event)
}

func (c *Calendar) UpdateEventContext(ctx context.Context, apikey string, calendarID string, eventID string, event model.UpdateEventRequest) (*model.EventResponse, error) {
	return c.UpdateEventCustomHTTPContext(ctx, apikey, c.httpClient, calendarID, eventID, event)
}

func (c *Calendar) UpdateEventCustomHTTP(apikey string, httpClient *http.Client, calendarID string, eventID string, event model.UpdateEventRequest) (*model.EventResponse, error) {
	return c.UpdateEventCustomHTTPContext(context.Background(), apikey, httpClient, calendarID, eventID, event)
}

func (c *Calendar) UpdateEventCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, calendarID string, eventID string, event model.UpdateEventRequest) (*model.EventResponse, error) {
	url := fmt.Sprintf("%s/calendar/v1/calendars/%s/events/%s", c.endPoint, calendarID, eventID)

	payload, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")
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
		return nil, fmt.Errorf("failed to update event, status: %s, body: %s", resp.Status, string(body))
	}

	var eventResponse model.EventResponse
	if err := json.Unmarshal(body, &eventResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	eventResponse.RawJSON = string(body)

	return &eventResponse, nil
}
