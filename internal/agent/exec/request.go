package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"regexp"
	"strings"
)

func Invoke(req domain.Request) (resp domain.Response, err error) {
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

func ReplaceAll(req *domain.Request, variableMap map[string]interface{}) {
	expressionValueMap := map[string]interface{}{}

	replaceUrl(req, variableMap, &expressionValueMap)
	replaceParams(req, variableMap, &expressionValueMap)
	replaceHeaders(req, variableMap, &expressionValueMap)
	replaceBody(req, variableMap, &expressionValueMap)
	replaceAuthor(req, variableMap, &expressionValueMap)
}

func replaceUrl(req *domain.Request, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	req.Url = ReplaceVariableValue(req.Url, variableMap)
	req.Url = ReplaceExpressionValue(req.Url, variableMap, expressionValueMap)
}
func replaceParams(req *domain.Request, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceVariableValue(param.Value, variableMap)
		req.Params[idx].Value = ReplaceExpressionValue(req.Params[idx].Value, variableMap, expressionValueMap)
	}
}
func replaceHeaders(req *domain.Request, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceVariableValue(header.Value, variableMap)
		req.Headers[idx].Value = ReplaceExpressionValue(req.Headers[idx].Value, variableMap, expressionValueMap)
	}
}
func replaceBody(req *domain.Request, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	req.Body = ReplaceVariableValue(req.Body, variableMap)
	req.Body = ReplaceExpressionValue(req.Body, variableMap, expressionValueMap)
}
func replaceAuthor(req *domain.Request, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceVariableValue(req.BasicAuth.Username, variableMap)
		req.BasicAuth.Password = ReplaceVariableValue(req.BasicAuth.Password, variableMap)

		req.BasicAuth.Username = ReplaceExpressionValue(req.BasicAuth.Username, variableMap, expressionValueMap)
		req.BasicAuth.Password = ReplaceExpressionValue(req.BasicAuth.Password, variableMap, expressionValueMap)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceVariableValue(req.BearerToken.Token, variableMap)
		req.BearerToken.Token = ReplaceExpressionValue(req.BearerToken.Token, variableMap, expressionValueMap)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceVariableValue(req.OAuth20.Name, variableMap)
		req.OAuth20.CallbackUrl = ReplaceVariableValue(req.OAuth20.CallbackUrl, variableMap)
		req.OAuth20.AuthURL = ReplaceVariableValue(req.OAuth20.AuthURL, variableMap)
		req.OAuth20.AccessTokenURL = ReplaceVariableValue(req.OAuth20.AccessTokenURL, variableMap)
		req.OAuth20.ClientID = ReplaceVariableValue(req.OAuth20.ClientID, variableMap)
		req.OAuth20.Scope = ReplaceVariableValue(req.OAuth20.Scope, variableMap)

		req.OAuth20.Name = ReplaceExpressionValue(req.OAuth20.Name, variableMap, expressionValueMap)
		req.OAuth20.CallbackUrl = ReplaceExpressionValue(req.OAuth20.CallbackUrl, variableMap, expressionValueMap)
		req.OAuth20.AuthURL = ReplaceExpressionValue(req.OAuth20.AuthURL, variableMap, expressionValueMap)
		req.OAuth20.AccessTokenURL = ReplaceExpressionValue(req.OAuth20.AccessTokenURL, variableMap, expressionValueMap)
		req.OAuth20.ClientID = ReplaceExpressionValue(req.OAuth20.ClientID, variableMap, expressionValueMap)
		req.OAuth20.Scope = ReplaceExpressionValue(req.OAuth20.Scope, variableMap, expressionValueMap)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceVariableValue(req.ApiKey.Key, variableMap)
		req.ApiKey.Value = ReplaceVariableValue(req.ApiKey.Value, variableMap)
		req.ApiKey.TransferMode = ReplaceVariableValue(req.ApiKey.TransferMode, variableMap)

		req.ApiKey.Key = ReplaceExpressionValue(req.ApiKey.Key, variableMap, expressionValueMap)
		req.ApiKey.Value = ReplaceExpressionValue(req.ApiKey.Value, variableMap, expressionValueMap)
		req.ApiKey.TransferMode = ReplaceExpressionValue(req.ApiKey.TransferMode, variableMap, expressionValueMap)
	}
}

func ReplaceVariableValue(value string, variableMap map[string]interface{}) (ret string) {
	variables := GetVariablesInVariablePlaceholder(value)
	ret = value

	for _, item := range variables {
		old := fmt.Sprintf("${%s}", item)
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
