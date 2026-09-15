package calendar

// DeleteType is the recurrence scope for POST .../events/{event-id}/delete.
type DeleteType string

const (
	DeleteTypeThis          DeleteType = "this"
	DeleteTypeWholeFromThis DeleteType = "wholeFromThis"
	DeleteTypeWhole         DeleteType = "whole"
)

// DeleteEventRequest is the body for deleting a calendar event.
// An empty DeleteType is sent as "this". For a recurring series, set DeleteType
// and use the occurrence id from GetEvents (it may include a timestamp suffix).
type DeleteEventRequest struct {
	DeleteType DeleteType `json:"deleteType,omitempty"`
}
