package wiki

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	model "github.com/dooray-go/dooray-sdk/openapi/model/wiki"
)

func (c *Wiki) GetWikis(apikey string) (*model.ListWikisResponse, error) {
	return c.GetWikisCustomHTTPContext(context.Background(), apikey, c.httpClient, 0, 0)
}
func (c *Wiki) GetWikisContext(ctx context.Context, apikey string) (*model.ListWikisResponse, error) {
	return c.GetWikisCustomHTTPContext(ctx, apikey, c.httpClient, 0, 0)
}
func (c *Wiki) GetWikisCustomHTTP(apikey string, httpClient *http.Client) (*model.ListWikisResponse, error) {
	return c.GetWikisCustomHTTPContext(context.Background(), apikey, httpClient, 0, 0)
}
func (c *Wiki) GetWikisWithOptions(apikey string, page, size int) (*model.ListWikisResponse, error) {
	return c.GetWikisCustomHTTPContext(context.Background(), apikey, c.httpClient, page, size)
}
func (c *Wiki) GetWikisCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, page, size int) (*model.ListWikisResponse, error) {
	var out model.ListWikisResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, "/wiki/v1/wikis", intQuery(page, size), nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) GetPage(apikey, pageID string) (*model.GetPageResponse, error) {
	return c.GetPageCustomHTTPContext(context.Background(), apikey, c.httpClient, pageID)
}
func (c *Wiki) GetPageContext(ctx context.Context, apikey, pageID string) (*model.GetPageResponse, error) {
	return c.GetPageCustomHTTPContext(ctx, apikey, c.httpClient, pageID)
}
func (c *Wiki) GetPageCustomHTTP(apikey string, httpClient *http.Client, pageID string) (*model.GetPageResponse, error) {
	return c.GetPageCustomHTTPContext(context.Background(), apikey, httpClient, pageID)
}
func (c *Wiki) GetPageCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, pageID string) (*model.GetPageResponse, error) {
	var out model.GetPageResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, "/wiki/v1/pages/"+pageID, nil, nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) CreatePage(apikey, wikiID string, req model.CreatePageRequest) (*model.CreatePageResponse, error) {
	return c.CreatePageCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, req)
}
func (c *Wiki) CreatePageContext(ctx context.Context, apikey, wikiID string, req model.CreatePageRequest) (*model.CreatePageResponse, error) {
	return c.CreatePageCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, req)
}
func (c *Wiki) CreatePageCustomHTTP(apikey string, httpClient *http.Client, wikiID string, req model.CreatePageRequest) (*model.CreatePageResponse, error) {
	return c.CreatePageCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, req)
}
func (c *Wiki) CreatePageCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID string, req model.CreatePageRequest) (*model.CreatePageResponse, error) {
	var out model.CreatePageResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/wiki/v1/wikis/%s/pages", wikiID), nil, req, &out, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) GetPages(apikey, wikiID, parentPageID string) (*model.ListPagesResponse, error) {
	return c.GetPagesCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, parentPageID)
}
func (c *Wiki) GetPagesContext(ctx context.Context, apikey, wikiID, parentPageID string) (*model.ListPagesResponse, error) {
	return c.GetPagesCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, parentPageID)
}
func (c *Wiki) GetPagesCustomHTTP(apikey string, httpClient *http.Client, wikiID, parentPageID string) (*model.ListPagesResponse, error) {
	return c.GetPagesCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, parentPageID)
}
func (c *Wiki) GetPagesCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, parentPageID string) (*model.ListPagesResponse, error) {
	q := url.Values{}
	if parentPageID != "" {
		q.Set("parentPageId", parentPageID)
	}
	var out model.ListPagesResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, fmt.Sprintf("/wiki/v1/wikis/%s/pages", wikiID), q, nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) GetWikiPage(apikey, wikiID, pageID string) (*model.GetPageResponse, error) {
	return c.GetWikiPageCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID)
}
func (c *Wiki) GetWikiPageContext(ctx context.Context, apikey, wikiID, pageID string) (*model.GetPageResponse, error) {
	return c.GetWikiPageCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID)
}
func (c *Wiki) GetWikiPageCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string) (*model.GetPageResponse, error) {
	return c.GetWikiPageCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID)
}
func (c *Wiki) GetWikiPageCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string) (*model.GetPageResponse, error) {
	var out model.GetPageResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodGet, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s", wikiID, pageID), nil, nil, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) UpdatePage(apikey, wikiID, pageID string, req model.UpdatePageRequest) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, req)
}
func (c *Wiki) UpdatePageContext(ctx context.Context, apikey, wikiID, pageID string, req model.UpdatePageRequest) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, req)
}
func (c *Wiki) UpdatePageCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string, req model.UpdatePageRequest) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, req)
}
func (c *Wiki) UpdatePageCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string, req model.UpdatePageRequest) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodPut, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s", wikiID, pageID), req)
}

func (c *Wiki) UpdatePageTitle(apikey, wikiID, pageID, subject string) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageTitleCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, subject)
}
func (c *Wiki) UpdatePageTitleContext(ctx context.Context, apikey, wikiID, pageID, subject string) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageTitleCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, subject)
}
func (c *Wiki) UpdatePageTitleCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID, subject string) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageTitleCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, subject)
}
func (c *Wiki) UpdatePageTitleCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID, subject string) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodPut, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/title", wikiID, pageID), model.UpdateTitleRequest{Subject: subject})
}

func (c *Wiki) UpdatePageContent(apikey, wikiID, pageID string, body model.Body) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageContentCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, body)
}
func (c *Wiki) UpdatePageContentContext(ctx context.Context, apikey, wikiID, pageID string, body model.Body) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageContentCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, body)
}
func (c *Wiki) UpdatePageContentCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string, body model.Body) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageContentCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, body)
}
func (c *Wiki) UpdatePageContentCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string, body model.Body) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodPut, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/content", wikiID, pageID), model.UpdateContentRequest{Body: body})
}

func (c *Wiki) UpdatePageReferrers(apikey, wikiID, pageID string, referrers []model.Actor) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageReferrersCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, referrers)
}
func (c *Wiki) UpdatePageReferrersContext(ctx context.Context, apikey, wikiID, pageID string, referrers []model.Actor) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageReferrersCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, referrers)
}
func (c *Wiki) UpdatePageReferrersCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string, referrers []model.Actor) (*model.HeaderOnlyResponse, error) {
	return c.UpdatePageReferrersCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, referrers)
}
func (c *Wiki) UpdatePageReferrersCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string, referrers []model.Actor) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodPut, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/referrers", wikiID, pageID), model.UpdateReferrersRequest{Referrers: referrers})
}

func (c *Wiki) MovePage(apikey, wikiID, pageID string, req model.MovePageRequest) (*model.MovePageResponse, error) {
	return c.MovePageCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID, req)
}
func (c *Wiki) MovePageContext(ctx context.Context, apikey, wikiID, pageID string, req model.MovePageRequest) (*model.MovePageResponse, error) {
	return c.MovePageCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID, req)
}
func (c *Wiki) MovePageCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string, req model.MovePageRequest) (*model.MovePageResponse, error) {
	return c.MovePageCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID, req)
}
func (c *Wiki) MovePageCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string, req model.MovePageRequest) (*model.MovePageResponse, error) {
	var out model.MovePageResponse
	body, err := c.doJSON(ctx, httpClient, apikey, http.MethodPost, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s/move", wikiID, pageID), nil, req, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}

func (c *Wiki) DeletePage(apikey, wikiID, pageID string) (*model.HeaderOnlyResponse, error) {
	return c.DeletePageCustomHTTPContext(context.Background(), apikey, c.httpClient, wikiID, pageID)
}
func (c *Wiki) DeletePageContext(ctx context.Context, apikey, wikiID, pageID string) (*model.HeaderOnlyResponse, error) {
	return c.DeletePageCustomHTTPContext(ctx, apikey, c.httpClient, wikiID, pageID)
}
func (c *Wiki) DeletePageCustomHTTP(apikey string, httpClient *http.Client, wikiID, pageID string) (*model.HeaderOnlyResponse, error) {
	return c.DeletePageCustomHTTPContext(context.Background(), apikey, httpClient, wikiID, pageID)
}
func (c *Wiki) DeletePageCustomHTTPContext(ctx context.Context, apikey string, httpClient *http.Client, wikiID, pageID string) (*model.HeaderOnlyResponse, error) {
	return c.headerCall(ctx, httpClient, apikey, http.MethodDelete, fmt.Sprintf("/wiki/v1/wikis/%s/pages/%s", wikiID, pageID), nil)
}

func (c *Wiki) headerCall(ctx context.Context, httpClient *http.Client, apikey, method, path string, payload any) (*model.HeaderOnlyResponse, error) {
	var out model.HeaderOnlyResponse
	body, err := c.doJSON(ctx, httpClient, apikey, method, path, nil, payload, &out)
	if err != nil {
		return nil, err
	}
	out.RawJSON = string(body)
	return &out, nil
}
