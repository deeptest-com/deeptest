package httpHelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(reqUrl string, reqParams []domain.Param) (code int, ret []byte, err error) {
	client := &http.Client{}

	if _consts.Verbose {
		_logUtils.Info(reqUrl)
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	queryParams := url.Values{}
	for _, param := range reqParams {
		queryParams.Add(param.Name, param.Value)
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	code = resp.StatusCode

	ret, _ = ioutil.ReadAll(resp.Body)
	if _consts.Verbose {
		_logUtils.PrintUnicode(ret)
	}

	return
}

func Post(reqUrl string, reqParams []domain.Param, data interface{}) (ret []byte, err error) {
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

	req, reqErr := http.NewRequest("POST", reqUrl, bytes.NewReader(dataBytes))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return
	}

	queryParams := url.Values{}
	for _, param := range reqParams {
		queryParams.Add(param.Name, param.Value)
	}
	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, respErr := client.Do(req)
	defer resp.Body.Close()

	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return
	}

	ret, _ = ioutil.ReadAll(resp.Body)
	if _consts.Verbose {
		_logUtils.PrintUnicode(ret)
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
