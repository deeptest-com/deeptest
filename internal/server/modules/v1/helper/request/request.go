package requestHelper

import (
	"encoding/base64"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/kataras/iris/v12"
	"strings"
)

func ReplaceAll(req *serverDomain.InvocationRequest, variableArr [][]string) {
	replaceUrl(req, variableArr)
	replaceParams(req, variableArr)
	replaceHeaders(req, variableArr)
	replaceBody(req, variableArr)
	replaceAuthor(req, variableArr)
}

func replaceUrl(req *serverDomain.InvocationRequest, variableArr [][]string) {
	req.Url = ReplaceValue(req.Url, variableArr, 0)
}
func replaceParams(req *serverDomain.InvocationRequest, variableArr [][]string) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceValue(param.Value, variableArr, 0)
	}
}
func replaceHeaders(req *serverDomain.InvocationRequest, variableArr [][]string) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceValue(header.Value, variableArr, 0)
	}
}
func replaceBody(req *serverDomain.InvocationRequest, variableArr [][]string) {
	req.Body = ReplaceValue(req.Body, variableArr, 0)
}
func replaceAuthor(req *serverDomain.InvocationRequest, variableArr [][]string) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceValue(req.BasicAuth.Username, variableArr, 0)
		req.BasicAuth.Password = ReplaceValue(req.BasicAuth.Password, variableArr, 0)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceValue(req.BearerToken.Token, variableArr, 0)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceValue(req.OAuth20.Name, variableArr, 0)
		req.OAuth20.CallbackUrl = ReplaceValue(req.OAuth20.CallbackUrl, variableArr, 0)
		req.OAuth20.AuthURL = ReplaceValue(req.OAuth20.AuthURL, variableArr, 0)
		req.OAuth20.AccessTokenURL = ReplaceValue(req.OAuth20.AccessTokenURL, variableArr, 0)
		req.OAuth20.ClientID = ReplaceValue(req.OAuth20.ClientID, variableArr, 0)
		req.OAuth20.Scope = ReplaceValue(req.OAuth20.Scope, variableArr, 0)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceValue(req.ApiKey.Key, variableArr, 0)
		req.ApiKey.Value = ReplaceValue(req.ApiKey.Value, variableArr, 0)
		req.ApiKey.TransferMode = ReplaceValue(req.ApiKey.TransferMode, variableArr, 0)
	}
}

func MergeVariables(
	environmentVariables []model.EnvironmentVar,
	interfaceExtractorVariables []serverDomain.Variable,
	processorExecVariables []domain.ExecVariable) (
	ret [][]string) {

	variableMap := iris.Map{}
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
				ret = append(ret, []string{fmt.Sprintf("${%s.%s}", key, propKey), fmt.Sprintf("%v", v)})
			}

		} else {
			ret = append(ret, []string{fmt.Sprintf("${%s}", key), fmt.Sprintf("%v", val)})

		}
	}

	return
}

func ReplaceValue(value string, variableArr [][]string, index int) (ret string) {
	if len(variableArr) == 0 || len(variableArr) <= index || !strings.Contains(value, "${") {
		ret = value
		return
	}

	old := variableArr[index][0]
	new := variableArr[index][1]
	ret = strings.ReplaceAll(value, old, new)

	ret = ReplaceValue(ret, variableArr, index+1)

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
