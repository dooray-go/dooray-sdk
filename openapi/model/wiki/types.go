package wiki

import (
	"encoding/json"

	"github.com/dooray-go/dooray-sdk/openapi/model"
)

type Body struct {
	MimeType string `json:"mimeType,omitempty"`
	Content  string `json:"content"`
}

type Member struct {
	OrganizationMemberID string `json:"organizationMemberId,omitempty"`
	Name                 string `json:"name,omitempty"`
}

type Actor struct {
	Type   string `json:"type"`
	Member Member `json:"member"`
}

type File struct {
	ID           string `json:"id"`
	PageFileID   string `json:"pageFileId,omitempty"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	AttachFileID string `json:"attachFileId,omitempty"`
	MimeType     string `json:"mimeType,omitempty"`
	Type         string `json:"type,omitempty"`
	Extension    string `json:"extension,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
}

type WikiInfo struct {
	ID      string `json:"id"`
	Project struct {
		ID string `json:"id"`
	} `json:"project"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Scope string `json:"scope"`
	Home  struct {
		PageID string `json:"pageId"`
	} `json:"home"`
}

type PageSummary struct {
	ID           string      `json:"id"`
	WikiID       string      `json:"wikiId"`
	Version      json.Number `json:"version"`
	ParentPageID string      `json:"parentPageId"`
	Subject      string      `json:"subject"`
	Root         bool        `json:"root"`
	Creator      Actor       `json:"creator"`
}

type PageDetail struct {
	PageSummary
	Body      Body    `json:"body"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	Referrers []Actor `json:"referrers"`
	Files     []File  `json:"files"`
	Images    []File  `json:"images"`
}

type HeaderOnlyResponse struct {
	Header  model.ResponseHeader `json:"header"`
	RawJSON string               `json:"-"`
}

type IDResult struct {
	ID string `json:"id"`
}

type IDResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  IDResult             `json:"result"`
	RawJSON string               `json:"-"`
}
