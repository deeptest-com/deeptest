package httpHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
	return gets(req, consts.GET, true)
}

func Post(req serverDomain.InvocationRequest) (
	ret serverDomain.InvocationResponse, err error) {

	return posts(req, consts.POST, true)
}

func Put(req serverDomain.InvocationRequest) (
	ret serverDomain.InvocationResponse, err error) {

	return posts(req, consts.PUT, true)
}

func Patch(req serverDomain.InvocationRequest) (
	ret serverDomain.InvocationResponse, err error) {

	return posts(req, consts.PATCH, true)
}

func Delete(req serverDomain.InvocationRequest) (
	ret serverDomain.InvocationResponse, err error) {

	return posts(req, consts.DELETE, true)
}

func Head(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
	return gets(req, consts.HEAD, false)
}

func Connect(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
	return gets(req, consts.CONNECT, false)
}

func Options(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
	return gets(req, consts.OPTIONS, false)
}

func Trace(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
	return gets(req, consts.TRACE, false)
}

func gets(req serverDomain.InvocationRequest, method consts.HttpMethod, readRespData bool) (
	ret serverDomain.InvocationResponse, err error) {

	reqUrl := req.Url
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
	ret.ContentLength = stringUtils.ParseInt(resp.Header.Get(consts.ContentLength))
	ret.Headers = getHeaders(resp.Header)

	if readRespData {
		content, _ := ioutil.ReadAll(resp.Body)
		if _consts.Verbose {
			_logUtils.PrintUnicode(content)
		}

		ret.Content = string(content)
	}

	return
}

func posts(req serverDomain.InvocationRequest, method consts.HttpMethod, readRespData bool) (
	ret serverDomain.InvocationResponse, err error) {

	reqUrl := req.Url
	reqParams := req.Params
	reqData := req.Body
	bodyType := req.BodyType

	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	client := &http.Client{}

	dataBytes, err := json.Marshal(reqData)
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

	request.Header.Set(consts.ContentType, bodyType.String())
	addAuthorInfo(req, request)

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(request)
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
	ret.ContentLength = stringUtils.ParseInt(resp.Header.Get(consts.ContentLength))
	ret.Headers = getHeaders(resp.Header)

	if readRespData {
		content, _ := ioutil.ReadAll(resp.Body)
		if _consts.Verbose {
			_logUtils.PrintUnicode(content)
		}

		ret.Content = string(content)
	}

	return
}

func addAuthorInfo(req serverDomain.InvocationRequest, request *http.Request) {
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
