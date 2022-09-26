package requestHelper

import (
	"encoding/base64"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	expressionHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/expression"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"regexp"
	"strings"
)

func ReplaceAll(req *serverDomain.InvocationRequest, variableArr [][]interface{}) {
	expressionValueMap := map[string]interface{}{}

	replaceUrl(req, variableArr, &expressionValueMap)
	replaceParams(req, variableArr, &expressionValueMap)
	replaceHeaders(req, variableArr, &expressionValueMap)
	replaceBody(req, variableArr, &expressionValueMap)
	replaceAuthor(req, variableArr, &expressionValueMap)
}

func replaceUrl(req *serverDomain.InvocationRequest, variableArr [][]interface{}, expressionValueMap *map[string]interface{}) {
	req.Url = ReplaceVariableValue(req.Url, variableArr, 0)
	req.Url = ReplaceExpressionValue(req.Url, variableArr, expressionValueMap)
}
func replaceParams(req *serverDomain.InvocationRequest, variableArr [][]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceVariableValue(param.Value, variableArr, 0)
		req.Params[idx].Value = ReplaceExpressionValue(req.Params[idx].Value, variableArr, expressionValueMap)
	}
}
func replaceHeaders(req *serverDomain.InvocationRequest, variableArr [][]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceVariableValue(header.Value, variableArr, 0)
		req.Headers[idx].Value = ReplaceExpressionValue(req.Headers[idx].Value, variableArr, expressionValueMap)
	}
}
func replaceBody(req *serverDomain.InvocationRequest, variableArr [][]interface{}, expressionValueMap *map[string]interface{}) {
	req.Body = ReplaceVariableValue(req.Body, variableArr, 0)
	req.Body = ReplaceExpressionValue(req.Body, variableArr, expressionValueMap)
}
func replaceAuthor(req *serverDomain.InvocationRequest, variableArr [][]interface{}, expressionValueMap *map[string]interface{}) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceVariableValue(req.BasicAuth.Username, variableArr, 0)
		req.BasicAuth.Password = ReplaceVariableValue(req.BasicAuth.Password, variableArr, 0)

		req.BasicAuth.Username = ReplaceExpressionValue(req.BasicAuth.Username, variableArr, expressionValueMap)
		req.BasicAuth.Password = ReplaceExpressionValue(req.BasicAuth.Password, variableArr, expressionValueMap)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceVariableValue(req.BearerToken.Token, variableArr, 0)
		req.BearerToken.Token = ReplaceExpressionValue(req.BearerToken.Token, variableArr, expressionValueMap)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceVariableValue(req.OAuth20.Name, variableArr, 0)
		req.OAuth20.CallbackUrl = ReplaceVariableValue(req.OAuth20.CallbackUrl, variableArr, 0)
		req.OAuth20.AuthURL = ReplaceVariableValue(req.OAuth20.AuthURL, variableArr, 0)
		req.OAuth20.AccessTokenURL = ReplaceVariableValue(req.OAuth20.AccessTokenURL, variableArr, 0)
		req.OAuth20.ClientID = ReplaceVariableValue(req.OAuth20.ClientID, variableArr, 0)
		req.OAuth20.Scope = ReplaceVariableValue(req.OAuth20.Scope, variableArr, 0)

		req.OAuth20.Name = ReplaceExpressionValue(req.OAuth20.Name, variableArr, expressionValueMap)
		req.OAuth20.CallbackUrl = ReplaceExpressionValue(req.OAuth20.CallbackUrl, variableArr, expressionValueMap)
		req.OAuth20.AuthURL = ReplaceExpressionValue(req.OAuth20.AuthURL, variableArr, expressionValueMap)
		req.OAuth20.AccessTokenURL = ReplaceExpressionValue(req.OAuth20.AccessTokenURL, variableArr, expressionValueMap)
		req.OAuth20.ClientID = ReplaceExpressionValue(req.OAuth20.ClientID, variableArr, expressionValueMap)
		req.OAuth20.Scope = ReplaceExpressionValue(req.OAuth20.Scope, variableArr, expressionValueMap)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceVariableValue(req.ApiKey.Key, variableArr, 0)
		req.ApiKey.Value = ReplaceVariableValue(req.ApiKey.Value, variableArr, 0)
		req.ApiKey.TransferMode = ReplaceVariableValue(req.ApiKey.TransferMode, variableArr, 0)

		req.ApiKey.Key = ReplaceExpressionValue(req.ApiKey.Key, variableArr, expressionValueMap)
		req.ApiKey.Value = ReplaceExpressionValue(req.ApiKey.Value, variableArr, expressionValueMap)
		req.ApiKey.TransferMode = ReplaceExpressionValue(req.ApiKey.TransferMode, variableArr, expressionValueMap)
	}
}

func MergeVariables(environmentVariables []model.EnvironmentVar, interfaceExtractorVariables []serverDomain.Variable,
	processorExecVariables []domain.ExecVariable) (
	ret [][]interface{}) {

	variableMap := map[string]interface{}{}
	for _, item := range environmentVariables {
		variableMap[item.Name] = item.Value
	}
	for _, item := range interfaceExtractorVariables { // overwrite previous ones
		variableMap[item.Name] = item.Value
	}
	for _, item := range processorExecVariables { // overwrite previous ones
		variableMap[item.Name] = item.Value
	}

	for key, val := range variableMap {
		valMp, isMap := val.(map[string]interface{})

		if isMap {
			for propKey, v := range valMp {
				ret = append(ret, []interface{}{fmt.Sprintf("${%s.%s}", key, propKey), v})
			}

		} else {
			ret = append(ret, []interface{}{fmt.Sprintf("${%s}", key), val})

		}
	}

	return
}

func ReplaceVariableValue(value string, variableArr [][]interface{}, index int) (ret string) {
	if len(variableArr) == 0 || len(variableArr) <= index || !strings.Contains(value, "${") {
		ret = value
		return
	}

	old := fmt.Sprintf("%v", variableArr[index][0])
	new := fmt.Sprintf("%v", variableArr[index][1])
	ret = strings.ReplaceAll(value, old, new)

	ret = ReplaceVariableValue(ret, variableArr, index+1)

	return
}

func ReplaceExpressionValue(value string, variableArr [][]interface{}, expressionValueMap *map[string]interface{}) (
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
			expressionValue, _ = expressionHelper.EvaluateGovaluateExpression(expressionWithoutSymbol, variableArr)
			(*expressionValueMap)[expressionWithoutSymbol] = expressionValue
		}

		ret = strings.ReplaceAll(ret, expressionWithSymbol, fmt.Sprintf("%v", expressionValue))
	}

	return
}

func Base64(str string) (ret string) {
	ret = base64.StdEncoding.EncodeToString([]byte(str))

	return
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
