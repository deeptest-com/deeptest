package httpHelper

import (
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"github.com/andybalholm/brotli"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func AddAuthorInfo(req domain.BaseRequest, request *http.Request) {
	if req.AuthorizationType == consts.BasicAuth {
		str := fmt.Sprintf("%s:%s", req.BasicAuth.Username, req.BasicAuth.Password)
		str = fmt.Sprintf("Basic %s", Base64(str))

		request.Header.Set(consts.Authorization, str)

	} else if req.AuthorizationType == consts.BearerToken {
		str := req.BearerToken.Token

		if !strings.HasPrefix(str, "Bearer ") {
			str = fmt.Sprintf("Bearer %s", req.BearerToken.Token)
		}

		request.Header.Set(consts.Authorization, str)

	} else if req.AuthorizationType == consts.OAuth2 {

	} else if req.AuthorizationType == consts.ApiKey {
		key := req.ApiKey.Key
		Value := req.ApiKey.Value

		if key != "" && Value != "" {
			request.Header.Set(key, Value)
		}
	}
}

func GetHeaders(header http.Header) (headers []domain.Header) {
	for key, val := range header {
		header := domain.Header{Name: key, Value: val[0]}
		headers = append(headers, header)
	}
	return
}
func GetCookies(cookies []*http.Cookie, jarCookies []*http.Cookie) (ret []domain.ExecCookie) {
	mp := map[string]bool{}

	for _, item := range cookies {
		cookie := domain.ExecCookie{
			Name:   item.Name,
			Value:  item.Value,
			Domain: item.Domain,
		}
		ret = append(ret, cookie)

		key := fmt.Sprintf("%s-%s-%s", item.Name, item.Value, item.Domain)
		mp[key] = true
	}

	for _, item := range jarCookies {
		key := fmt.Sprintf("%s-%s-%s", item.Name, item.Value, item.Domain)
		if _, ok := mp[key]; ok {
			continue
		}

		cookie := domain.ExecCookie{
			Name:   item.Name,
			Value:  item.Value,
			Domain: item.Domain,
		}
		ret = append(ret, cookie)
	}

	return
}

func GenUrl(server string, path string) string {
	server = UpdateUrl(server)
	url := fmt.Sprintf("%sapi/v1/%s", server, path)
	return url
}

func UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}

func WrapperErrInResp(code int, statusContent string, content string, resp *domain.DebugResponse) {
	resp.StatusCode = code
	resp.StatusContent = fmt.Sprintf("%d %s", code, statusContent)
	resp.Content, _ = url.QueryUnescape(content)
}

func DecodeResponseBody(resp *http.Response) (err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "br":
		resp.Body = io.NopCloser(brotli.NewReader(resp.Body))
	case "gzip":
		resp.Body, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		resp.ContentLength = -1 // set to unknown to avoid Content-Length mismatched
	case "deflate":
		resp.Body, err = zlib.NewReader(resp.Body)
		if err != nil {
			return err
		}
		resp.ContentLength = -1 // set to unknown to avoid Content-Length mismatched
	}
	return nil
}

func IsXmlContent(str string) bool {
	return strings.Contains(str, "xml")
}
func IsHtmlContent(str string) bool {
	return strings.Contains(str, "html")
}
func IsJsonContent(str string) bool {
	return strings.Contains(str, "json")
}
func IsImageContent(str string) bool {
	return strings.Contains(str, "image")
}

func Base64(str string) (ret string) {
	ret = base64.StdEncoding.EncodeToString([]byte(str))

	return
}

func IsStreamResponse(contentType consts.HttpContentType) bool {
	return strings.Index(contentType.String(), consts.ContentTypeStream.String()) > -1
}
