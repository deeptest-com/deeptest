package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"regexp"
	"strings"
)

func Invoke(req domain.Request) (resp domain.Response, err error) {
	req.Url, err = httpHelper.AddDefaultUrlSchema(req.Url)
	if err != nil {
		return
	}

	if req.Method == consts.GET {
		resp, err = utils.Get(req)
	} else if req.Method == consts.POST {
		resp, err = utils.Post(req)
	} else if req.Method == consts.PUT {
		resp, err = utils.Put(req)
	} else if req.Method == consts.DELETE {
		resp, err = utils.Delete(req)
	} else if req.Method == consts.PATCH {
		resp, err = utils.Patch(req)
	} else if req.Method == consts.HEAD {
		resp, err = utils.Head(req)
	} else if req.Method == consts.CONNECT {
		resp, err = utils.Connect(req)
	} else if req.Method == consts.OPTIONS {
		resp, err = utils.Options(req)
	} else if req.Method == consts.TRACE {
		resp, err = utils.Trace(req)
	}

	return
}

func GetRequestProps(req *domain.Request) {
	req.BodyLang = consts.LangPlainTEXT

	arr := strings.Split(string(req.BodyType), "/")
	if len(arr) == 1 {
		return
	}

	typeName := arr[1]
	if typeName == "text" || typeName == "plain" {
		typeName = "plaintext"
	}

	req.BodyLang = consts.HttpRespLangType(typeName)
}

func GetContentProps(resp *domain.Response) {
	resp.ContentLang = consts.LangPlainTEXT

	if resp.ContentLang == "" {
		return
	}

	arr := strings.Split(string(resp.ContentType), ";")

	arr1 := strings.Split(arr[0], "/")
	if len(arr1) == 1 {
		return
	}

	typeName := arr1[1]
	if typeName == "text" || typeName == "plain" {
		typeName = "plaintext"
	}
	resp.ContentLang = consts.HttpRespLangType(typeName)

	if len(arr) > 1 {
		arr2 := strings.Split(arr[1], "=")
		if len(arr2) > 1 {
			resp.ContentCharset = consts.HttpRespCharset(arr2[1])
		}
	}

	//ret.Content = mockHelper.FormatXml(ret.Content)

	return
}

func ReplaceAll(req *domain.BaseRequest, variableMap map[string]interface{}) {
	expressionValueMapCache := map[string]interface{}{}

	replaceUrl(req, variableMap, &expressionValueMapCache)
	replaceParams(req, variableMap, &expressionValueMapCache)
	replaceHeaders(req, variableMap, &expressionValueMapCache)
	replaceBody(req, variableMap, &expressionValueMapCache)
	replaceAuthor(req, variableMap, &expressionValueMapCache)
}

func replaceUrl(req *domain.BaseRequest, variableMap map[string]interface{}, expressionValueMapCache *map[string]interface{}) {
	req.Url = ReplaceExpressionAndVariableValue(req.Url, variableMap, expressionValueMapCache)
}
func replaceParams(req *domain.BaseRequest, variableMap map[string]interface{}, expressionValueMapCache *map[string]interface{}) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceExpressionAndVariableValue(param.Value, variableMap, expressionValueMapCache)
	}
}
func replaceHeaders(req *domain.BaseRequest, variableMap map[string]interface{}, expressionValueMapCache *map[string]interface{}) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceExpressionAndVariableValue(header.Value, variableMap, expressionValueMapCache)
	}
}
func replaceBody(req *domain.BaseRequest, variableMap map[string]interface{}, expressionValueMapCache *map[string]interface{}) {
	req.Body = ReplaceExpressionAndVariableValue(req.Body, variableMap, expressionValueMapCache)
}
func replaceAuthor(req *domain.BaseRequest, variableMap map[string]interface{}, expressionValueMapCache *map[string]interface{}) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceExpressionAndVariableValue(req.BasicAuth.Username, variableMap, expressionValueMapCache)
		req.BasicAuth.Password = ReplaceExpressionAndVariableValue(req.BasicAuth.Password, variableMap, expressionValueMapCache)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceExpressionAndVariableValue(req.BearerToken.Token, variableMap, expressionValueMapCache)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceExpressionAndVariableValue(req.OAuth20.Name, variableMap, expressionValueMapCache)
		req.OAuth20.CallbackUrl = ReplaceExpressionAndVariableValue(req.OAuth20.CallbackUrl, variableMap, expressionValueMapCache)
		req.OAuth20.AuthURL = ReplaceExpressionAndVariableValue(req.OAuth20.AuthURL, variableMap, expressionValueMapCache)
		req.OAuth20.AccessTokenURL = ReplaceExpressionAndVariableValue(req.OAuth20.AccessTokenURL, variableMap, expressionValueMapCache)
		req.OAuth20.ClientID = ReplaceExpressionAndVariableValue(req.OAuth20.ClientID, variableMap, expressionValueMapCache)
		req.OAuth20.Scope = ReplaceExpressionAndVariableValue(req.OAuth20.Scope, variableMap, expressionValueMapCache)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceExpressionAndVariableValue(req.ApiKey.Key, variableMap, expressionValueMapCache)
		req.ApiKey.Value = ReplaceExpressionAndVariableValue(req.ApiKey.Value, variableMap, expressionValueMapCache)
		req.ApiKey.TransferMode = ReplaceExpressionAndVariableValue(req.ApiKey.TransferMode, variableMap, expressionValueMapCache)
	}
}

func ReplaceExpressionAndVariableValue(value string, variableMap map[string]interface{},
	expressionValueMapCache *map[string]interface{}) (ret string) {

	ret = ReplaceExpressionValue(value, variableMap, expressionValueMapCache)
	ret = ReplaceVariableValue(ret, variableMap)

	return
}

func ReplaceVariableValue(value string, variableMap map[string]interface{}) (ret string) {
	variables := GetVariablesInVariablePlaceholder(value)
	ret = value

	for _, item := range variables {
		variablePlaceholde := fmt.Sprintf("${%s}", item)
		old := variablePlaceholde
		new := fmt.Sprintf("%v", variableMap[item])

		ret = strings.ReplaceAll(ret, old, new)
	}

	return
}

func ReplaceExpressionValue(value string, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) (
	ret string) {

	ret = value

	regex := regexp.MustCompile(`(?Ui)\$expr\('(.*)'\)`) // $expr('uuid()')
	arrOfArr := regex.FindAllStringSubmatch(ret, -1)

	if len(arrOfArr) == 0 {
		return
	}

	for _, arr := range arrOfArr {
		expressionWithSymbol := arr[0]
		expressionWithoutSymbol := arr[1]

		expressionValue, ok := (*expressionValueMap)[expressionWithoutSymbol]
		if !ok {
			expressionValue, _ = EvaluateGovaluateExpressionWithVariables(expressionWithoutSymbol, variableMap)
			(*expressionValueMap)[expressionWithoutSymbol] = expressionValue
		}

		ret = strings.ReplaceAll(ret, expressionWithSymbol, fmt.Sprintf("%v", expressionValue))
	}

	return
}
