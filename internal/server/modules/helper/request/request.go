package requestHelper

import (
	"encoding/base64"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/expression"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"regexp"
	"strings"
)

func ReplaceAll(req *v1.InvocationRequest, variableMap map[string]interface{}) {
	expressionValueMap := map[string]interface{}{}

	replaceUrl(req, variableMap, &expressionValueMap)
	replaceParams(req, variableMap, &expressionValueMap)
	replaceHeaders(req, variableMap, &expressionValueMap)
	replaceBody(req, variableMap, &expressionValueMap)
	replaceAuthor(req, variableMap, &expressionValueMap)
}

func replaceUrl(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	req.Url = ReplaceVariableValue(req.Url, variableMap)
	req.Url = ReplaceExpressionValue(req.Url, variableMap, expressionValueMap)
}
func replaceParams(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceVariableValue(param.Value, variableMap)
		req.Params[idx].Value = ReplaceExpressionValue(req.Params[idx].Value, variableMap, expressionValueMap)
	}
}
func replaceHeaders(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceVariableValue(header.Value, variableMap)
		req.Headers[idx].Value = ReplaceExpressionValue(req.Headers[idx].Value, variableMap, expressionValueMap)
	}
}
func replaceBody(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	req.Body = ReplaceVariableValue(req.Body, variableMap)
	req.Body = ReplaceExpressionValue(req.Body, variableMap, expressionValueMap)
}
func replaceAuthor(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
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

func MergeVariables(environmentVariables []model.EnvironmentVar, interfaceExtractorVariables []v1.Variable,
	processorExecVariables []domain.ExecVariable) (
	ret map[string]interface{}) {

	ret = map[string]interface{}{}

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
				ret[fmt.Sprintf("${%s.%s}", key, propKey)] = v
			}

		} else {
			ret[fmt.Sprintf("${%s}", key)] = val

		}
	}

	return
}

func ReplaceVariableValue(value string, variableMap map[string]interface{}) (ret string) {
	variables := execHelper.GetVariablesInVariablePlaceholder(value)
	ret = value

	for _, item := range variables {
		old := fmt.Sprintf("${%s}", item)
		new := fmt.Sprintf("%v", variableMap[old])

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
			expressionValue, _ = expressionHelper.EvaluateGovaluateExpression(expressionWithoutSymbol, variableMap)
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
