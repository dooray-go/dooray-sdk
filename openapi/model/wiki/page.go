package wiki

import "github.com/dooray-go/dooray-sdk/openapi/model"

type ListWikisResponse struct {
	Header     model.ResponseHeader `json:"header"`
	Result     []WikiInfo           `json:"result"`
	TotalCount int                  `json:"totalCount"`
	RawJSON    string               `json:"-"`
}

type GetPageResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  PageDetail           `json:"result"`
	RawJSON string               `json:"-"`
}

type ListPagesResponse struct {
	Header     model.ResponseHeader `json:"header"`
	Result     []PageSummary        `json:"result"`
	TotalCount int                  `json:"totalCount"`
	RawJSON    string               `json:"-"`
}

type CreatePageRequest struct {
	ParentPageID  string   `json:"parentPageId,omitempty"`
	Subject       string   `json:"subject"`
	Body          Body     `json:"body"`
	AttachFileIDs []string `json:"attachFileIds,omitempty"`
	Referrers     []Actor  `json:"referrers,omitempty"`
}

type CreatePageResult struct {
	ID           string `json:"id"`
	WikiID       string `json:"wikiId"`
	ParentPageID string `json:"parentPageId"`
	Version      int    `json:"version"`
}

type CreatePageResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  CreatePageResult     `json:"result"`
	RawJSON string               `json:"-"`
}

type UpdatePageRequest struct {
	Subject   string   `json:"subject"`
	Body      Body     `json:"body"`
	Referrers *[]Actor `json:"referrers,omitempty"`
}

type UpdateTitleRequest struct {
	Subject string `json:"subject"`
}

type UpdateContentRequest struct {
	Body Body `json:"body"`
}

type UpdateReferrersRequest struct {
	Referrers []Actor `json:"referrers"`
}

type MovePageRequest struct {
	TargetWikiID       string `json:"targetWikiId,omitempty"`
	TargetParentPageID string `json:"targetParentPageId"`
	WithChildren       *bool  `json:"withChildren,omitempty"`
	BeforePageID       string `json:"beforePageId,omitempty"`
}

type MovePageResult struct {
	WikiID       string `json:"wikiId"`
	PageID       string `json:"pageId"`
	ParentPageID string `json:"parentPageId"`
	Version      int    `json:"version"`
}

type MovePageResponse struct {
	Header  model.ResponseHeader `json:"header"`
	Result  MovePageResult       `json:"result"`
	RawJSON string               `json:"-"`
}
