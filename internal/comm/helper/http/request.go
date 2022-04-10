package httpHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/string"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get(reqUrl string, reqParams []domain.Param) (ret serverDomain.TestResponse, err error) {
	return gets(reqUrl, consts.GET, reqParams, true)
}

func Post(reqUrl string, reqParams []domain.Param, data interface{}, bodyType consts.HttpContentType) (
	ret serverDomain.TestResponse, err error) {

	return posts(reqUrl, consts.POST, reqParams, data, bodyType, true)
}

func Put(reqUrl string, reqParams []domain.Param, data interface{}, bodyType consts.HttpContentType) (
	ret serverDomain.TestResponse, err error) {

	return posts(reqUrl, consts.PUT, reqParams, data, bodyType, true)
}

func Patch(reqUrl string, reqParams []domain.Param, data interface{}, bodyType consts.HttpContentType) (
	ret serverDomain.TestResponse, err error) {

	return posts(reqUrl, consts.PATCH, reqParams, data, bodyType, true)
}

func Delete(reqUrl string, reqParams []domain.Param, data interface{}, bodyType consts.HttpContentType) (
	ret serverDomain.TestResponse, err error) {

	return posts(reqUrl, consts.DELETE, reqParams, data, bodyType, true)
}

func Head(reqUrl string, reqParams []domain.Param) (ret serverDomain.TestResponse, err error) {
	return gets(reqUrl, consts.HEAD, reqParams, false)
}

func Connect(reqUrl string, reqParams []domain.Param) (ret serverDomain.TestResponse, err error) {
	return gets(reqUrl, consts.CONNECT, reqParams, false)
}

func Options(reqUrl string, reqParams []domain.Param) (ret serverDomain.TestResponse, err error) {
	return gets(reqUrl, consts.OPTIONS, reqParams, false)
}

func Trace(reqUrl string, reqParams []domain.Param) (ret serverDomain.TestResponse, err error) {
	return gets(reqUrl, consts.TRACE, reqParams, false)
}

func gets(reqUrl string, method consts.HttpMethod, reqParams []domain.Param, readRespData bool) (
	ret serverDomain.TestResponse, err error) {

	client := &http.Client{}

	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	req, err := http.NewRequest(method.String(), reqUrl, nil)
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	queryParams := url.Values{}
	for _, param := range reqParams {
		queryParams.Add(param.Name, param.Value)
	}
	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set("Origin", "DEEPTEST")

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(req)
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

func posts(reqUrl string, method consts.HttpMethod, reqParams []domain.Param, data interface{},
	bodyType consts.HttpContentType, readRespData bool) (
	ret serverDomain.TestResponse, err error) {

	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	client := &http.Client{}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		_logUtils.Infof(color.RedString("marshal request failed, error: %s.", err.Error()))
		return
	}

	if _consts.Verbose {
		_logUtils.Infof(string(dataBytes))
	}

	req, reqErr := http.NewRequest(method.String(), reqUrl, bytes.NewReader(dataBytes))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return
	}

	queryParams := url.Values{}
	for _, param := range reqParams {
		queryParams.Add(param.Name, param.Value)
	}
	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set(consts.ContentType, bodyType.String())

	startTime := time.Now().UnixMilli()

	resp, err := client.Do(req)
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
