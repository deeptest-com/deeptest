package agentExec

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Logger *zap.Logger

func Get(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.GET, true, "", nil)
}
func Post(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return posts(req, consts.POST, true, "", nil)
}
func Put(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return posts(req, consts.PUT, true, "", nil)
}
func Patch(req domain.BaseRequest) (
	ret domain.DebugResponse, err error) {
	return posts(req, consts.PATCH, true, "", nil)
}
func Delete(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return posts(req, consts.DELETE, true, "", nil)
}
func Head(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.HEAD, false, "", nil)
}
func Connect(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.CONNECT, false, "", nil)
}
func Options(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.OPTIONS, false, "", nil)
}
func Trace(req domain.BaseRequest) (ret domain.DebugResponse, err error) {
	return gets(req, consts.TRACE, false, "", nil)
}

func GetBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return gets(req, consts.GET, true, key, wsMsg)
}
func PostBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return posts(req, consts.POST, true, key, wsMsg)
}
func PutBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return posts(req, consts.PUT, true, key, wsMsg)
}
func PatchBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (
	ret domain.DebugResponse, err error) {
	return posts(req, consts.PATCH, true, key, wsMsg)
}
func DeleteBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return posts(req, consts.DELETE, true, key, wsMsg)
}
func HeadBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return gets(req, consts.HEAD, false, key, wsMsg)
}
func ConnectBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return gets(req, consts.CONNECT, false, key, wsMsg)
}
func OptionsBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return gets(req, consts.OPTIONS, false, key, wsMsg)
}
func TraceBy(req domain.BaseRequest, key string, wsMsg *websocket.Message) (ret domain.DebugResponse, err error) {
	return gets(req, consts.TRACE, false, key, wsMsg)
}

func gets(req domain.BaseRequest, method consts.HttpMethod, readRespData bool,
	key string, wsMsg *websocket.Message) (
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
		httpHelper.WrapperErrInResp(consts.ServiceUnavailable.Int(), "请求错误", err.Error(), &ret)
		return
	}

	// decode response body in br/gzip/deflate formats
	err = httpHelper.DecodeResponseBody(resp)
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

	ret.Headers = httpHelper.GetHeaders(resp.Header)
	ret.Cookies = httpHelper.GetCookies(resp.Cookies(), nil)

	if readRespData {
		generateRespBody(resp, &ret, key, wsMsg)
	}

	return
}

func posts(req domain.BaseRequest, method consts.HttpMethod, readRespData bool,
	key string, wsMsg *websocket.Message) (
	ret domain.DebugResponse, err error) {

	reqUrl := commUtils.RemoveLeftVariableSymbol(req.Url)
	if _consts.Verbose {
		Logger.Info(reqUrl)
	}

	jar := genCookies(req)

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
	if httpHelper.IsFormBody(bodyType) {
		bodyFormData := GenBodyFormData(req)

		formDataWriter, _ := httpHelper.MultipartEncoder(bodyFormData)
		formDataContentType = httpHelper.MultipartContentType(formDataWriter)

		dataBytes = formDataWriter.Payload.Bytes()

	} else if httpHelper.IsFormUrlencodedBody(bodyType) {
		bodyFormUrlencoded := GenBodyFormUrlencoded(req)
		dataBytes = []byte(bodyFormUrlencoded)

	} else if httpHelper.IsJsonBody(bodyType) {
		// post json
		reqBody := commUtils.CompressedJson(req.Body)
		dataBytes = []byte(reqBody)
		if err != nil {
			return
		}
	} else {
		// post text
		dataBytes = []byte(req.Body)
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
		httpHelper.WrapperErrInResp(consts.ServiceUnavailable.Int(), "请求错误", err.Error(), &ret)
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

	ret.Headers = httpHelper.GetHeaders(resp.Header)
	ret.Cookies = httpHelper.GetCookies(resp.Cookies(), nil)

	if readRespData {
		generateRespBody(resp, &ret, key, wsMsg)
	}

	return
}

func generateRespBody(resp *http.Response, result *domain.DebugResponse, key string, wsMsg *websocket.Message) (
	err error) {

	if wsMsg != nil && httpHelper.IsStreamResponse(result.ContentType) {
		result := iris.Map{
			"source":   "execInterface",
			"response": result,
		}
		execUtils.SendExecMsg(result, consts.ProgressResult, wsMsg)

		r := bufio.NewReader(resp.Body)
		for {
			bytes, err1 := r.ReadBytes('\n')
			str := string(bytes)

			if err1 != nil && err1 != io.EOF {
				err = err1
				break
			}
			if err1 == io.EOF {
				break
			}

			fmt.Println("\n>>> stream response item: " + str + "\n")

			result := iris.Map{
				"source":     "execInterface",
				"streamItem": str,
			}
			execUtils.SendExecMsg(result, consts.ProgressResult, wsMsg)

			if IsExecCtxCancel(key) {
				break
			}
		}

	} else {
		reader := resp.Body
		if resp.Header.Get("Content-Encoding") == "gzip" {
			reader, _ = gzip.NewReader(resp.Body)
		}

		unicodeContent, _ := ioutil.ReadAll(reader)
		if httpHelper.IsImageContent(result.ContentType.String()) {
			imgBase64Str := base64.StdEncoding.EncodeToString(unicodeContent)
			result.Content = imgBase64Str
			return
		}

		utf8Content, _ := _stringUtils.UnescapeUnicode(unicodeContent)
		if _consts.Verbose {
			_logUtils.Info(string(utf8Content))
		}

		result.Content = string(utf8Content)
		result.Content = strings.ReplaceAll(result.Content, "\u0000", "")
	}

	return
}
