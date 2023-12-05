package httpHelper

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/andybalholm/brotli"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var Logger *zap.Logger

func Get(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.GET, true)
}

func Post(req domain.BaseRequest) (
	ret domain.DebugResponse, err error) {

	return posts(req, consts.POST, true)
}

func Put(req domain.BaseRequest) (
	ret domain.DebugResponse, err error) {

	return posts(req, consts.PUT, true)
}

func Patch(req domain.BaseRequest) (
	ret domain.DebugResponse, err error) {

	return posts(req, consts.PATCH, true)
}

func Delete(req domain.BaseRequest) (
	ret domain.DebugResponse, err error) {

	return posts(req, consts.DELETE, true)
}

func Head(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.HEAD, false)
}

func Connect(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.CONNECT, false)
}

func Options(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.OPTIONS, false)
}

func Trace(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.TRACE, false)
}

func gets(req domain.BaseRequest, method consts.HttpMethod, readRespData bool) (
	ret domain.DebugResponse, err error) {

	reqUrl := commUtils.RemoveLeftVariableSymbol(req.Url)
	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	jar := genCookies(req)

	client := &http.Client{
		Jar:     jar,
		Timeout: consts.HttpRequestTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	httpReq, err := http.NewRequest(method.String(), reqUrl, nil)
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	dealwithQueryParams(req, httpReq)
	dealwithHeader(req, httpReq)
	//dealwithCookie(req, httpReq)

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(httpReq)
	if err != nil {
		wrapperErrInResp(consts.ServiceUnavailable.Int(), "请求错误", err.Error(), &ret)
		_logUtils.Error(err.Error())
		return
	}

	// decode response body in br/gzip/deflate formats
	err = decodeResponseBody(resp)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	endTime := time.Now().UnixMilli()
	ret.Time = endTime - startTime

	ret.StatusCode = resp.StatusCode
	ret.StatusContent = resp.Status
	ret.ContentType = consts.HttpContentType(resp.Header.Get(consts.ContentType))
	ret.ContentLength = _stringUtils.ParseInt(resp.Header.Get(consts.ContentLength))

	//u, _ := url.Parse(req.Url)
	ret.Headers = getHeaders(resp.Header)
	ret.Cookies = getCookies(resp.Cookies(), nil)
	//ret.Cookies = getCookies(resp.Cookies(), jar.Cookies(u))

	if !readRespData {
		return
	}
	reader := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	}

	unicodeContent, err := ioutil.ReadAll(reader)
	if IsImageContent(ret.ContentType.String()) {
		imgBase64Str := base64.StdEncoding.EncodeToString(unicodeContent)
		ret.Content = imgBase64Str
		return
	}

	utf8Content, _ := _stringUtils.UnescapeUnicode(unicodeContent)
	if _consts.Verbose {
		_logUtils.Info(string(utf8Content))
	}

	ret.Content = string(utf8Content)
	ret.Content = strings.ReplaceAll(ret.Content, "\u0000", "")

	return
}

func posts(req domain.BaseRequest, method consts.HttpMethod, readRespData bool) (
	ret domain.DebugResponse, err error) {

	reqUrl := commUtils.RemoveLeftVariableSymbol(req.Url)
	if _consts.Verbose {
		Logger.Info(reqUrl)
	}

	jar := genCookies(req)

	reqBody := req.Body
	bodyType := req.BodyType

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar:     jar,
		Timeout: consts.HttpRequestTimeout,
	}

	var dataBytes []byte

	formDataContentType := ""
	if IsFormBody(bodyType) {
		bodyFormData := genBodyFormData(req)

		formDataWriter, _ := MultipartEncoder(bodyFormData)
		formDataContentType = MultipartContentType(formDataWriter)

		dataBytes = formDataWriter.Payload.Bytes()

	} else if IsFormUrlencodedBody(bodyType) {
		bodyFormUrlencoded := genBodyFormUrlencoded(req)
		dataBytes = []byte(bodyFormUrlencoded)

	} else if IsJsonBody(bodyType) {
		// post json
		dataBytes = []byte(reqBody)
		if err != nil {
			return
		}
	}

	if err != nil {
		_logUtils.Infof(color.RedString("marshal httpReq failed, error: %s.", err.Error()))
		return
	}

	if _consts.Verbose {
		_logUtils.Infof(string(dataBytes))
	}

	httpReq, reqErr := http.NewRequest(method.String(), reqUrl, bytes.NewReader(dataBytes))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return
	}

	dealwithQueryParams(req, httpReq)
	dealwithHeader(req, httpReq)
	//dealwithCookie(req, httpReq)

	// body type
	if strings.HasPrefix(bodyType.String(), consts.ContentTypeJSON.String()) {
		httpReq.Header.Set(consts.ContentType, fmt.Sprintf("%s; charset=utf-8", bodyType))
	} else if strings.HasPrefix(bodyType.String(), consts.ContentTypeFormData.String()) {
		httpReq.Header.Set(consts.ContentType, formDataContentType)
	} else {
		httpReq.Header.Set(consts.ContentType, bodyType.String())
	}

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(httpReq)
	if err != nil {
		wrapperErrInResp(consts.ServiceUnavailable.Int(), "请求错误", err.Error(), &ret)
		_logUtils.Error(err.Error())
		return
	}

	defer resp.Body.Close()

	endTime := time.Now().UnixMilli()
	ret.Time = endTime - startTime

	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	ret.StatusCode = resp.StatusCode
	ret.StatusContent = resp.Status

	ret.ContentType = consts.HttpContentType(resp.Header.Get(consts.ContentType))
	ret.ContentLength = _stringUtils.ParseInt(resp.Header.Get(consts.ContentLength))

	//u, _ := url.Parse(req.Url)
	//ret.Cookies = getCookies(resp.Cookies(), jar.Cookies(u))
	ret.Headers = getHeaders(resp.Header)
	ret.Cookies = getCookies(resp.Cookies(), nil)

	if !readRespData {
		return
	}

	reader := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	}

	unicodeContent, _ := ioutil.ReadAll(reader)
	utf8Content, _ := _stringUtils.UnescapeUnicode(unicodeContent)

	if _consts.Verbose {
		_logUtils.Info(string(utf8Content))
	}

	ret.Content = string(utf8Content)

	return
}

func addAuthorInfo(req domain.BaseRequest, request *http.Request) {
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

func getHeaders(header http.Header) (headers []domain.Header) {
	for key, val := range header {
		header := domain.Header{Name: key, Value: val[0]}
		headers = append(headers, header)
	}
	return
}
func getCookies(cookies []*http.Cookie, jarCookies []*http.Cookie) (ret []domain.ExecCookie) {
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

func wrapperErrInResp(code int, statusContent string, content string, resp *domain.DebugResponse) {
	resp.StatusCode = code
	resp.StatusContent = fmt.Sprintf("%d %s", code, statusContent)
	resp.Content, _ = url.QueryUnescape(content)
}

func decodeResponseBody(resp *http.Response) (err error) {
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
