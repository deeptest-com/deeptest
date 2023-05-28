package httpHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func genCookies(req domain.BaseRequest) (ret http.CookieJar) {
	ret, _ = cookiejar.New(nil)

	var cookies []*http.Cookie
	mp := map[string]bool{}
	for _, c := range req.Cookies {
		key := fmt.Sprintf("%s=%s", c.Name, c.Domain)
		if _, ok := mp[key]; ok {
			continue
		}

		cookies = append(cookies, &http.Cookie{
			Name:   c.Name,
			Value:  _stringUtils.InterfToStr(c.Value),
			Domain: c.Domain,
		})
		mp[key] = true
	}

	urlStr, _ := url.Parse(req.Url)
	ret.SetCookies(urlStr, cookies)

	return
}

func dealwithQueryParams(req domain.BaseRequest, httpReq *http.Request) {
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

func dealwithHeader(req domain.BaseRequest, httpReq *http.Request) {
	httpReq.Header.Set("User-Agent", consts.UserAgentChrome)
	httpReq.Header.Set("Origin", "DEEPTEST")

	for _, h := range req.Headers {
		if h.Name != "" && httpReq.Header.Get(h.Name) == "" {
			httpReq.Header.Set(h.Name, h.Value)
		}
	}

	addAuthorInfo(req, httpReq)
}

func DealwithCookie(req domain.BaseRequest, httpReq *http.Request) {
	httpReq.Header.Set("User-Agent", consts.UserAgentChrome)
	httpReq.Header.Set("Origin", "DEEPTEST")

	for _, h := range req.Headers {
		if h.Name != "" && httpReq.Header.Get(h.Name) == "" {
			httpReq.Header.Set(h.Name, h.Value)
		}
	}

	addAuthorInfo(req, httpReq)
}
