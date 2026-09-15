package calendar

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	model "github.com/dooray-go/dooray-sdk/openapi/model/calendar"
	"github.com/dooray-go/dooray-sdk/utils"
)

func TestUpdateEvent_SendsOnlySetFields(t *testing.T) {
	const (
		calendarID = "test-calendar-id"
		eventID    = "test-event-id"
		apiKey     = "test-api-key"
	)
	wantBody := `{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":""},"result":null}`

	var gotMethod, gotAuth, gotCalendarID, gotEventID string
	var rawPayload []byte

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /calendar/v1/calendars/{calendarId}/events/{eventId}", func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotAuth = r.Header.Get("Authorization")
		gotCalendarID = r.PathValue("calendarId")
		gotEventID = r.PathValue("eventId")
		var err error
		rawPayload, err = io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("failed to read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(wantBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	subject := "Updated Event"
	location := "Updated Location"
	startedAt := utils.NewJsonTime(time.Date(2026, 9, 15, 10, 0, 0, 0, time.UTC))
	response, err := NewCalendar(server.URL).UpdateEvent(apiKey, calendarID, eventID, model.UpdateEventRequest{
		Subject:   &subject,
		Location:  &location,
		StartedAt: &startedAt,
	})
	if err != nil {
		t.Fatalf("UpdateEvent returned an error: %v", err)
	}

	if gotMethod != http.MethodPut {
		t.Errorf("method: want %s, got %s", http.MethodPut, gotMethod)
	}
	if gotAuth != "dooray-api "+apiKey {
		t.Errorf("Authorization: want %q, got %q", "dooray-api "+apiKey, gotAuth)
	}
	if gotCalendarID != calendarID {
		t.Errorf("calendarId: want %q, got %q", calendarID, gotCalendarID)
	}
	if gotEventID != eventID {
		t.Errorf("eventId: want %q, got %q", eventID, gotEventID)
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(rawPayload, &payload); err != nil {
		t.Fatalf("payload JSON: %v", err)
	}
	if len(payload) != 3 {
		t.Errorf("expected only set fields, got keys %v from %s", keys(payload), rawPayload)
	}
	for _, key := range []string{"users", "body", "endedAt", "wholeDayFlag", "recurrenceRule", "personalSettings"} {
		if _, ok := payload[key]; ok {
			t.Errorf("omitted field %q was sent: %s", key, rawPayload)
		}
	}
	if response.RawJSON != wantBody {
		t.Errorf("RawJSON: want %s, got %s", wantBody, response.RawJSON)
	}
	if !response.Header.IsSuccessful {
		t.Error("header.isSuccessful: want true, got false")
	}
}

func TestUpdateEvent_SendsWholeDayFlagFalse(t *testing.T) {
	var rawPayload []byte
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /calendar/v1/calendars/{calendarId}/events/{eventId}", func(w http.ResponseWriter, r *http.Request) {
		rawPayload, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":""},"result":null}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	flag := false
	if _, err := NewCalendar(server.URL).UpdateEvent("key", "cal", "evt", model.UpdateEventRequest{WholeDayFlag: &flag}); err != nil {
		t.Fatalf("UpdateEvent: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(rawPayload, &payload); err != nil {
		t.Fatalf("payload JSON: %v", err)
	}
	if payload["wholeDayFlag"] != false {
		t.Errorf("wholeDayFlag false must be sent, got %s", rawPayload)
	}
}

func TestUpdateEvent_NonOKStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /calendar/v1/calendars/{calendarId}/events/{eventId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":false,"resultCode":403,"resultMessage":"forbidden"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	subject := "x"
	_, err := NewCalendar(server.URL).UpdateEvent("key", "cal", "evt", model.UpdateEventRequest{Subject: &subject})
	if err == nil {
		t.Fatal("expected error for non-OK status, got nil")
	}
}

func keys(m map[string]json.RawMessage) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
