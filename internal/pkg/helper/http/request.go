package httpHelper

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"crypto/tls"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/request"
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/andybalholm/brotli"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func Get(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	return gets(req, consts.GET, true)
}

func Post(req v1.InvocationRequest) (
	ret v1.InvocationResponse, err error) {

	return posts(req, consts.POST, true)
}

func Put(req v1.InvocationRequest) (
	ret v1.InvocationResponse, err error) {

	return posts(req, consts.PUT, true)
}

func Patch(req v1.InvocationRequest) (
	ret v1.InvocationResponse, err error) {

	return posts(req, consts.PATCH, true)
}

func Delete(req v1.InvocationRequest) (
	ret v1.InvocationResponse, err error) {

	return posts(req, consts.DELETE, true)
}

func Head(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	return gets(req, consts.HEAD, false)
}

func Connect(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	return gets(req, consts.CONNECT, false)
}

func Options(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	return gets(req, consts.OPTIONS, false)
}

func Trace(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	return gets(req, consts.TRACE, false)
}

func gets(req v1.InvocationRequest, method consts.HttpMethod, readRespData bool) (
	ret v1.InvocationResponse, err error) {

	reqUrl := commUtils.RemoveLeftVariableSymbol(req.Url)
	reqParams := req.Params
	reqHeaders := req.Headers

	client := &http.Client{}

	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	request, err := http.NewRequest(method.String(), reqUrl, nil)
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	queryParams := url.Values{}
	for _, queryParam := range strings.Split(request.URL.RawQuery, "&") {
		arr := strings.Split(queryParam, "=")
		if len(arr) > 1 {
			queryParams.Add(arr[0], arr[1])
		}
	}
	for _, param := range reqParams {
		queryParams.Add(param.Name, param.Value)
	}
	request.URL.RawQuery = queryParams.Encode()

	for _, header := range reqHeaders {
		request.Header.Set(header.Name, header.Value)
	}

	request.Header.Set("User-Agent", consts.UserAgentChrome)
	request.Header.Set("Origin", "DEEPTEST")
	addAuthorInfo(req, request)

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(request)
	if err != nil {
		wrapperErrInResp(consts.ServiceUnavailable, "请求错误", err.Error(), &ret)
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

	ret.StatusCode = consts.HttpRespCode(resp.StatusCode)
	ret.StatusContent = resp.Status
	ret.ContentType = consts.HttpContentType(resp.Header.Get(consts.ContentType))
	ret.ContentLength = _stringUtils.ParseInt(resp.Header.Get(consts.ContentLength))
	ret.Headers = getHeaders(resp.Header)

	if !readRespData {
		return
	}
	reader := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	}

	unicodeContent, err := ioutil.ReadAll(reader)
	utf8Content, _ := _stringUtils.UnescapeUnicode(unicodeContent)

	if _consts.Verbose {
		_logUtils.Info(string(utf8Content))
	}

	ret.Content = string(utf8Content)

	return
}

func posts(req v1.InvocationRequest, method consts.HttpMethod, readRespData bool) (
	ret v1.InvocationResponse, err error) {

	reqUrl := commUtils.RemoveLeftVariableSymbol(req.Url)
	reqHeaders := req.Headers
	reqParams := req.Params
	reqData := req.Body
	bodyType := req.BodyType

	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar:     jar, // insert response cookies into request
		Timeout: 120 * time.Second,
	}
	//http2Client := &http.Client{
	//	Transport: &http2.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	},
	//	Timeout: 120 * time.Second,
	//}

	var dataBytes []byte

	if strings.HasPrefix(bodyType.String(), "application/x-www-form-urlencoded") { // post form data
		formData := make(url.Values)
		for _, param := range reqParams {
			formData.Add(param.Name, param.Value)
		}
		dataBytes = []byte(formData.Encode())
	} else {
		dataBytes = []byte(reqData)
	}

	if err != nil {
		_logUtils.Infof(color.RedString("marshal request failed, error: %s.", err.Error()))
		return
	}

	if _consts.Verbose {
		_logUtils.Infof(string(dataBytes))
	}

	request, reqErr := http.NewRequest(method.String(), reqUrl, bytes.NewReader(dataBytes))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return
	}

	queryParams := url.Values{}
	for _, queryParam := range strings.Split(request.URL.RawQuery, "&") {
		arr := strings.Split(queryParam, "=")
		if len(arr) > 1 {
			queryParams.Add(arr[0], arr[1])
		}
	}
	for _, param := range reqParams {
		queryParams.Add(param.Name, param.Value)
	}
	request.URL.RawQuery = queryParams.Encode()

	for _, header := range reqHeaders {
		request.Header.Set(header.Name, header.Value)
	}

	request.Header.Set(consts.ContentType, bodyType.String())
	addAuthorInfo(req, request)

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(request)
	if err != nil {
		wrapperErrInResp(consts.ServiceUnavailable, "请求错误", err.Error(), &ret)
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

	ret.StatusCode = consts.HttpRespCode(resp.StatusCode)
	ret.StatusContent = resp.Status

	ret.ContentType = consts.HttpContentType(resp.Header.Get(consts.ContentType))
	ret.ContentLength = _stringUtils.ParseInt(resp.Header.Get(consts.ContentLength))
	ret.Headers = getHeaders(resp.Header)

	if readRespData {
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
	}

	return
}

func addAuthorInfo(req v1.InvocationRequest, request *http.Request) {
	if req.AuthorizationType == consts.BasicAuth {
		str := fmt.Sprintf("%s:%s", req.BasicAuth.Username, req.BasicAuth.Password)
		str = fmt.Sprintf("Basic %s", requestHelper.Base64(str))

		request.Header.Set(consts.Authorization, str)

	} else if req.AuthorizationType == consts.BearerToken {
		str := fmt.Sprintf("Bearer %s", req.BearerToken.Token)
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

func wrapperErrInResp(code consts.HttpRespCode, statusContent string, content string, resp *v1.InvocationResponse) {
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
