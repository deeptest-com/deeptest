package gojaUtils

import (
	"encoding/json"
	valueUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/value"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	httpHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/http"
	"github.com/dop251/goja"
	"log"
	"reflect"
	"strings"
)

func GenRequest(data goja.Value, vm *goja.Runtime) (req domain.BaseRequest) {
	dataObj := data.Export()
	dataType := reflect.TypeOf(dataObj).Name()

	if dataType == "string" { // simple get
		log.Println("url is ", data)

		req = domain.BaseRequest{
			Method: consts.GET,
			Url:    dataObj.(string),
		}

		return

	}
	log.Println("url is ", data.ToObject(vm).Get("url"))

	dataMap := dataObj.(map[string]interface{})

	paramsObj := dataMap["params"]
	if paramsObj == nil {
		paramsObj = map[string]interface{}{}
	}
	paramsMap := paramsObj.(map[string]interface{})

	headersObj := dataMap["headers"]
	if headersObj == nil {
		headersObj = map[string]interface{}{}
	}
	headersMap := headersObj.(map[string]interface{})

	var authType consts.AuthorType
	basicAuthObj := dataMap["basicAuth"]
	basicAuthMap := map[string]interface{}{}
	if basicAuthObj != nil {
		authType = consts.BasicAuth
		basicAuthMap = basicAuthObj.(map[string]interface{})
	}

	bearerTokenObj := dataMap["bearerToken"]
	bearerTokenMap := map[string]interface{}{}
	if bearerTokenObj != nil {
		authType = consts.BearerToken
		bearerTokenMap = bearerTokenObj.(map[string]interface{})
	}

	contentType := consts.HttpContentType(valueUtils.InterfaceToStr(headersMap["Content-Type"]))
	if contentType == "" {
		contentType = consts.ContentTypeJSON
	}

	var bodyBytes []byte
	var bodyFormData []domain.BodyFormDataItem
	var bodyFormUrlEncodedData []domain.BodyFormUrlEncodedItem

	if httpHelper.IsJsonBody(contentType) {
		bodyBytes, _ = json.Marshal(dataMap["body"])

	} else if httpHelper.IsFormBody(contentType) || httpHelper.IsFormUrlencodedBody(contentType) {
		formItemsObject := dataMap["formItems"]
		if formItemsObject != nil {
			formItemsMap := formItemsObject.(map[string]interface{})

			for key, val := range formItemsMap {
				if httpHelper.IsFormBody(contentType) {
					bodyFormData = append(bodyFormData, domain.BodyFormDataItem{
						Name:  key,
						Value: valueUtils.InterfaceToStr(val),
						Type:  consts.FormDataTypeText,
					})
				} else if httpHelper.IsFormUrlencodedBody(contentType) {
					bodyFormUrlEncodedData = append(bodyFormUrlEncodedData, domain.BodyFormUrlEncodedItem{
						Name:  key,
						Value: valueUtils.InterfaceToStr(val),
					})
				}
			}
		}
	}

	// create request
	req = domain.BaseRequest{
		Method: consts.HttpMethod(strings.ToUpper(valueUtils.InterfaceToStr(dataMap["method"]))),
		Url:    valueUtils.InterfaceToStr(dataMap["url"]),

		BodyType: contentType,
		Body:     string(bodyBytes),
	}

	if paramsObj != nil {
		for key, val := range paramsMap {
			param := domain.Param{
				Name:  key,
				Value: valueUtils.InterfaceToStr(val),
			}

			if req.QueryParams == nil {
				req.QueryParams = &[]domain.Param{}
			}
			*req.QueryParams = append(*req.QueryParams, param)
		}
	}

	if headersObj != nil {
		for key, val := range headersMap {
			header := domain.Header{
				Name:  key,
				Value: valueUtils.InterfaceToStr(val),
			}

			if req.Headers == nil {
				req.Headers = &[]domain.Header{}
			}
			*req.Headers = append(*req.Headers, header)
		}
	}

	req.AuthorizationType = authType
	if basicAuthObj != nil {
		req.BasicAuth = domain.BasicAuth{
			Username: valueUtils.InterfaceToStr(basicAuthMap["username"]),
			Password: valueUtils.InterfaceToStr(basicAuthMap["password"]),
		}
	}
	if bearerTokenObj != nil {
		req.BearerToken = domain.BearerToken{
			Token: valueUtils.InterfaceToStr(bearerTokenMap["token"]),
		}
	}

	if bodyFormData != nil || bodyFormUrlEncodedData != nil {
		if httpHelper.IsFormBody(contentType) {
			req.BodyFormData = &bodyFormData
		} else if httpHelper.IsFormUrlencodedBody(contentType) {
			req.BodyFormUrlencoded = &bodyFormUrlEncodedData
		}
	}

	return
}
