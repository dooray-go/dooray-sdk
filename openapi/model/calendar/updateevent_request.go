package calendar

import "github.com/dooray-go/dooray-sdk/utils"

// UpdateEventRequest is the body for PUT /calendar/v1/calendars/{calendar-id}/events/{event-id}.
// Omitted (nil/empty) fields are left unchanged by the Dooray API.
type UpdateEventRequest struct {
	Users            *Users            `json:"users,omitempty"`
	Subject          *string           `json:"subject,omitempty"`
	Body             *Body             `json:"body,omitempty"`
	StartedAt        *utils.JsonTime   `json:"startedAt,omitempty"`
	EndedAt          *utils.JsonTime   `json:"endedAt,omitempty"`
	WholeDayFlag     *bool             `json:"wholeDayFlag,omitempty"`
	Location         *string           `json:"location,omitempty"`
	RecurrenceRule   *RecurrenceRule   `json:"recurrenceRule,omitempty"`
	PersonalSettings *PersonalSettings `json:"personalSettings,omitempty"`
}
