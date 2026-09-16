package wiki

import (
	"net/http"

	"github.com/dooray-go/dooray-sdk/utils"
)

type Wiki struct {
	endPoint   string
	httpClient *http.Client
}

func NewDefaultWiki() *Wiki {
	return &Wiki{
		endPoint:   "https://api.dooray.com",
		httpClient: utils.NewDefaultHTTPClient(),
	}
}

func NewWiki(endPoint string) *Wiki {
	return &Wiki{
		endPoint:   endPoint,
		httpClient: utils.NewDefaultHTTPClient(),
	}
}

func NewWikiWithClient(endPoint string, httpClient *http.Client) *Wiki {
	return &Wiki{
		endPoint:   endPoint,
		httpClient: httpClient,
	}
}
