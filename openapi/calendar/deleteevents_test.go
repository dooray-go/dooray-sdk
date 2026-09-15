package calendar

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dooray-go/dooray-sdk/openapi/model/calendar"
)

const deleteSuccessBody = `{"header":{"isSuccessful":true,"resultCode":0,"resultMessage":""},"result":null}`

func TestDeleteEvent_SingleEventDefaultsToThis(t *testing.T) {
	const (
		calendarID = "3216797327178181418"
		eventID    = "3748560866668749754"
		apiKey     = "test-api-key"
	)

	var gotMethod, gotCalendarID, gotEventID string
	var gotPayload model.DeleteEventRequest

	mux := http.NewServeMux()
	mux.HandleFunc("POST /calendar/v1/calendars/{calendarId}/events/{eventId}/delete", func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotCalendarID = r.PathValue("calendarId")
		gotEventID = r.PathValue("eventId")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("failed to read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(body, &gotPayload); err != nil {
			t.Errorf("failed to unmarshal request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(deleteSuccessBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	response, err := NewCalendar(server.URL).DeleteEvent(apiKey, calendarID, eventID, model.DeleteEventRequest{})
	if err != nil {
		t.Fatalf("DeleteEvent returned an error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("method: want %s, got %s", http.MethodPost, gotMethod)
	}
	if gotCalendarID != calendarID {
		t.Errorf("calendarId: want %q, got %q", calendarID, gotCalendarID)
	}
	if gotEventID != eventID {
		t.Errorf("eventId: want %q, got %q", eventID, gotEventID)
	}
	if gotPayload.DeleteType != model.DeleteTypeThis {
		t.Errorf("deleteType: want %q, got %q", model.DeleteTypeThis, gotPayload.DeleteType)
	}
	if response.RawJSON != deleteSuccessBody {
		t.Errorf("RawJSON: want %s, got %s", deleteSuccessBody, response.RawJSON)
	}
	if !response.Header.IsSuccessful {
		t.Error("header.isSuccessful: want true, got false")
	}
}

func TestDeleteEvent_RecurringOccurrence(t *testing.T) {
	const (
		calendarID = "3216797327178181418"
		eventID    = "3748560866668749754-20240228T013000Z"
	)

	var gotEventID string
	var gotPayload model.DeleteEventRequest

	mux := http.NewServeMux()
	mux.HandleFunc("POST /calendar/v1/calendars/{calendarId}/events/{eventId}/delete", func(w http.ResponseWriter, r *http.Request) {
		gotEventID = r.PathValue("eventId")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("failed to read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(body, &gotPayload); err != nil {
			t.Errorf("failed to unmarshal request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(deleteSuccessBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewCalendar(server.URL).DeleteEvent("key", calendarID, eventID, model.DeleteEventRequest{
		DeleteType: model.DeleteTypeWholeFromThis,
	})
	if err != nil {
		t.Fatalf("DeleteEvent: %v", err)
	}
	if gotEventID != eventID {
		t.Errorf("eventId: want %q, got %q", eventID, gotEventID)
	}
	if gotPayload.DeleteType != model.DeleteTypeWholeFromThis {
		t.Errorf("deleteType: want %q, got %q", model.DeleteTypeWholeFromThis, gotPayload.DeleteType)
	}
}

func TestDeleteEvent_NonOKStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /calendar/v1/calendars/{calendarId}/events/{eventId}/delete", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"header":{"isSuccessful":false,"resultCode":403,"resultMessage":"forbidden"}}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewCalendar(server.URL).DeleteEvent("key", "cal", "evt", model.DeleteEventRequest{DeleteType: model.DeleteTypeThis})
	if err == nil {
		t.Fatal("expected error for non-OK status, got nil")
	}
}
