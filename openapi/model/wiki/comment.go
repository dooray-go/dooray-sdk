package wiki

import "github.com/dooray-go/dooray-sdk/openapi/model"

type CommentBody struct {
	Content  string `json:"content"`
	MimeType string `json:"mimeType,omitempty"`
}

type CommentRequest struct {
	Body CommentBody `json:"body"`
}

type Comment struct {
	ID   string `json:"id"`
	Page struct {
		ID string `json:"id"`
	} `json:"page"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	Creator    Actor  `json:"creator"`
	Body       Body   `json:"body"`
}

type ListCommentsResponse struct {
	Header     model.ResponseHeader `json:"header"`
	Result     []Comment            `json:"result"`
	TotalCount int                  `json:"totalCount"`
	RawJSON    string               `json:"-"`
}

type GetCommentResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  Comment              `json:"result"`
	RawJSON string               `json:"-"`
}

type SharedLink struct {
	ID                 string `json:"id"`
	SharedLink         string `json:"sharedLink"`
	Scope              string `json:"scope"`
	IncludeDescendants bool   `json:"includeDescendants"`
}

type ListSharedLinksResponse struct {
	Header     model.ResponseHeader `json:"header"`
	Result     []SharedLink         `json:"result"`
	TotalCount int                  `json:"totalCount"`
	RawJSON    string               `json:"-"`
}

type UploadFileResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  File                 `json:"result"`
	RawJSON string               `json:"-"`
}

type Download struct {
	Content     []byte
	ContentType string
	StatusCode  int
}
