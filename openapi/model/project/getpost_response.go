package project

import "github.com/dooray-go/dooray-sdk/openapi/model"

// PostFile is an attachment on a post detail response.
type PostFile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// PostDetail is a single post including body and files.
type PostDetail struct {
	PostInfo
	Body       PostBody   `json:"body"`
	Files      []PostFile `json:"files"`
	FileIDList []string   `json:"fileIdList"`
}

// GetPostResponse is GET /project/v1/projects/{project-id}/posts/{post-id}.
type GetPostResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  PostDetail           `json:"result"`
	RawJSON string               `json:"-"`
}
