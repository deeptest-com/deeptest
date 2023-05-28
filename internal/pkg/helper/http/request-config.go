package httpHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"net/http"
	"net/url"
	"strings"
)

func DealwithQueryParams(req domain.BaseRequest, httpReq *http.Request) {
	queryParams := url.Values{}

	for _, pair := range strings.Split(httpReq.URL.RawQuery, "&") {
		arr := strings.Split(pair, "=")
		if len(arr) > 1 {
			queryParams.Add(arr[0], arr[1])
		}
	}

	for _, p := range req.QueryParams {
		name := strings.ToUpper(p.Name)

		if name != "" && queryParams.Get(name) == "" {
			queryParams.Add(name, p.Value)
		}
	}

	httpReq.URL.RawQuery = queryParams.Encode()
}

func DealwithHeader(req domain.BaseRequest, httpReq *http.Request) {
	httpReq.Header.Set("User-Agent", consts.UserAgentChrome)
	httpReq.Header.Set("Origin", "DEEPTEST")

	for _, h := range req.Headers {
		if h.Name != "" && httpReq.Header.Get(h.Name) == "" {
			httpReq.Header.Set(h.Name, h.Value)
		}
	}

	addAuthorInfo(req, httpReq)
}
