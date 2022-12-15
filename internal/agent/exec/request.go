package agentExec

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"regexp"
	"strings"
)

func Invoke(req v1.BaseRequest) (resp v1.InvocationResponse, err error) {
	req.Url, err = _httpUtils.AddDefaultUrlSchema(req.Url)
	if err != nil {
		return
	}

	if req.Method == consts.GET {
		resp, err = httpHelper.Get(req)
	} else if req.Method == consts.POST {
		resp, err = httpHelper.Post(req)
	} else if req.Method == consts.PUT {
		resp, err = httpHelper.Put(req)
	} else if req.Method == consts.DELETE {
		resp, err = httpHelper.Delete(req)
	} else if req.Method == consts.PATCH {
		resp, err = httpHelper.Patch(req)
	} else if req.Method == consts.HEAD {
		resp, err = httpHelper.Head(req)
	} else if req.Method == consts.CONNECT {
		resp, err = httpHelper.Connect(req)
	} else if req.Method == consts.OPTIONS {
		resp, err = httpHelper.Options(req)
	} else if req.Method == consts.TRACE {
		resp, err = httpHelper.Trace(req)
	}

	return
}

func GetRequestProps(req *v1.BaseRequest) {
	req.BodyLang = consts.LangPlainTEXT

	arr := strings.Split(string(req.BodyType), "/")
	if len(arr) == 1 {
		return
	}

	typeName := arr[1]
	if typeName == "text" || typeName == "plain" {
		typeName = consts.LangPlainTEXT.String()
	}

	req.BodyLang = consts.HttpRespLangType(typeName)
}

func GetContentProps(resp *v1.InvocationResponse) {
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

func ReplaceAll(req *v1.BaseRequest, variableMap map[string]interface{}) {
	replaceUrl(req, variableMap)
	replaceParams(req, variableMap)
	replaceHeaders(req, variableMap)
	replaceBody(req, variableMap)
	replaceAuthor(req, variableMap)
}

func replaceUrl(req *v1.BaseRequest, variableMap map[string]interface{}) {
	req.Url = ReplaceVariableValue(req.Url, variableMap)
}
func replaceParams(req *v1.BaseRequest, variableMap map[string]interface{}) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceVariableValue(param.Value, variableMap)
	}
}
func replaceHeaders(req *v1.BaseRequest, variableMap map[string]interface{}) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceVariableValue(header.Value, variableMap)
	}
}
func replaceBody(req *v1.BaseRequest, variableMap map[string]interface{}) {
	req.Body = ReplaceVariableValue(req.Body, variableMap)
}
func replaceAuthor(req *v1.BaseRequest, variableMap map[string]interface{}) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceVariableValue(req.BasicAuth.Username, variableMap)
		req.BasicAuth.Password = ReplaceVariableValue(req.BasicAuth.Password, variableMap)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceVariableValue(req.BearerToken.Token, variableMap)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceVariableValue(req.OAuth20.Name, variableMap)
		req.OAuth20.CallbackUrl = ReplaceVariableValue(req.OAuth20.CallbackUrl, variableMap)
		req.OAuth20.AuthURL = ReplaceVariableValue(req.OAuth20.AuthURL, variableMap)
		req.OAuth20.AccessTokenURL = ReplaceVariableValue(req.OAuth20.AccessTokenURL, variableMap)
		req.OAuth20.ClientID = ReplaceVariableValue(req.OAuth20.ClientID, variableMap)
		req.OAuth20.Scope = ReplaceVariableValue(req.OAuth20.Scope, variableMap)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceVariableValue(req.ApiKey.Key, variableMap)
		req.ApiKey.Value = ReplaceVariableValue(req.ApiKey.Value, variableMap)
		req.ApiKey.TransferMode = ReplaceVariableValue(req.ApiKey.TransferMode, variableMap)
	}
}

//func ReplaceExpressionAndVariableValue(value string, variableMap map[string]interface{},
//	expressionValueMapCache *map[string]interface{}) (ret string) {
//
//	ret = ReplaceExpressionValue(value, variableMap, expressionValueMapCache)
//	ret = ReplaceVariableValue(ret, variableMap)
//
//	return
//}

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

	regex := regexp.MustCompile(`(?Ui)\$expr\("(.*)"\)`) // $expr("uuid()")
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
