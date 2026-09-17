package messenger

import "github.com/dooray-go/dooray-sdk/openapi/model"

type MemberRef struct {
	Type   string `json:"type"`
	Member struct {
		OrganizationMemberID string `json:"organizationMemberId"`
	} `json:"member"`
}

type ChannelMe struct {
	MemberRef
	Role string `json:"role"`
}

type Channel struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Organization struct {
		ID string `json:"id"`
	} `json:"organization"`
	Type  string `json:"type"`
	Users struct {
		Participants []MemberRef `json:"participants"`
	} `json:"users"`
	Me         ChannelMe `json:"me"`
	Capacity   int       `json:"capacity"`
	Status     string    `json:"status"`
	CreatedAt  string    `json:"createdAt"`
	UpdatedAt  string    `json:"updatedAt"`
	Displayed  bool      `json:"displayed"`
	Role       string    `json:"role"`
	ArchivedAt *string   `json:"archivedAt"`
}

type ListChannelsResponse struct {
	Header     model.ResponseHeader `json:"header"`
	Result     []Channel            `json:"result"`
	TotalCount int                  `json:"totalCount"`
	RawJSON    string               `json:"-"`
}

type CreateChannelRequest struct {
	Type      string   `json:"type"`
	Capacity  string   `json:"capacity,omitempty"`
	MemberIDs []string `json:"memberIds"`
	Title     string   `json:"title,omitempty"`
}

type IDResult struct {
	ID string `json:"id"`
}

type IDResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  IDResult             `json:"result"`
	RawJSON string               `json:"-"`
}

type MemberIDsRequest struct {
	MemberIDs []string `json:"memberIds"`
}

type TextRequest struct {
	Text string `json:"text"`
}

type HeaderOnlyResponse struct {
	Header  model.ResponseHeader `json:"header"`
	RawJSON string               `json:"-"`
}

type CreateAndSendThreadRequest struct {
	Text       string `json:"text"`
	ThreadText string `json:"threadText,omitempty"`
}
