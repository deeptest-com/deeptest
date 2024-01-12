package httpHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
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
	if req.Cookies != nil {
		for _, c := range *req.Cookies {
			key := fmt.Sprintf("%s=%s", c.Name, c.Domain)
			if _, ok := mp[key]; ok { // skip duplicate one
				continue
			}

			//domain := strings.TrimSpace(c.Domain)
			//if domain == "127.0.0.1" {
			//	domain = "localhost"
			//}

			coo := http.Cookie{
				Name:  c.Name,
				Value: _stringUtils.InterfToStr(c.Value),
				//Domain:  domain,
				Path: c.Path,
			}
			if c.ExpireTime != nil {
				coo.Expires = *c.ExpireTime
			}

			cookies = append(cookies, &coo)

			mp[key] = true
		}
	}

	urlStr, _ := url.Parse(req.Url)
	ret.SetCookies(urlStr, cookies)

	return
}

func genBodyFormData(req domain.BaseRequest) (formData []domain.BodyFormDataItem) {
	mp := map[string]bool{}

	if req.BodyFormData != nil {
		for _, item := range *req.BodyFormData {
			key := item.Name
			if _, ok := mp[key]; ok { // skip duplicate one
				continue
			}

			formData = append(formData, item)
			mp[key] = true
		}
	}

	return
}
func genBodyFormUrlencoded(req domain.BaseRequest) (ret string) {
	mp := map[string]bool{}
	formData := make(url.Values)

	if req.BodyFormUrlencoded != nil {
		for _, item := range *req.BodyFormUrlencoded {
			key := item.Name
			if _, ok := mp[key]; ok { // skip duplicate one
				continue
			}

			formData.Add(item.Name, item.Value)
			mp[key] = true
		}
	}

	ret = formData.Encode()

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

	if req.QueryParams != nil {
		for _, p := range *req.QueryParams {
			name := p.Name

			if !p.Disabled && name != "" && queryParams.Get(name) == "" {
				queryParams.Add(name, p.Value)
			}
		}
	}

	httpReq.URL.RawQuery = queryParams.Encode()
}

func dealwithHeader(req domain.BaseRequest, httpReq *http.Request) {
	httpReq.Header.Set("User-Agent", consts.UserAgentChrome)
	httpReq.Header.Set("Origin", "DEEPTEST")

	if req.Headers != nil {
		for _, h := range *req.Headers {
			if !h.Disabled && h.Name != "" && httpReq.Header.Get(h.Name) == "" {
				httpReq.Header.Set(h.Name, h.Value)
			}
		}
	}

	addAuthorInfo(req, httpReq)

	logUtils.Debugf("httpReq-header:%s", commonUtils.JsonEncode(httpReq.Header))
}

//func dealwithCookie(req domain.BaseRequest, httpReq *http.Request) {
//	httpReq.Header.Set("User-Agent", consts.UserAgentChrome)
//	httpReq.Header.Set("Origin", "DEEPTEST")
//
//	for _, h := range req.Headers {
//		if !h.Disabled && h.Name != "" && httpReq.Header.Get(h.Name) == "" {
//			httpReq.Header.Set(h.Name, h.Value)
//		}
//	}
//}
