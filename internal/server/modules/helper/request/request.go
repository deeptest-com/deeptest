package requestHelper

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
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
	req.Url = agentExec.ReplaceVariableValue(req.Url, variableMap)
	req.Url = agentExec.ReplaceExpressionValue(req.Url, variableMap, expressionValueMap)
}
func replaceParams(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, param := range req.Params {
		req.Params[idx].Value = agentExec.ReplaceVariableValue(param.Value, variableMap)
		req.Params[idx].Value = agentExec.ReplaceExpressionValue(req.Params[idx].Value, variableMap, expressionValueMap)
	}
}
func replaceHeaders(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = agentExec.ReplaceVariableValue(header.Value, variableMap)
		req.Headers[idx].Value = agentExec.ReplaceExpressionValue(req.Headers[idx].Value, variableMap, expressionValueMap)
	}
}
func replaceBody(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	req.Body = agentExec.ReplaceVariableValue(req.Body, variableMap)
	req.Body = agentExec.ReplaceExpressionValue(req.Body, variableMap, expressionValueMap)
}
func replaceAuthor(req *v1.InvocationRequest, variableMap map[string]interface{}, expressionValueMap *map[string]interface{}) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = agentExec.ReplaceVariableValue(req.BasicAuth.Username, variableMap)
		req.BasicAuth.Password = agentExec.ReplaceVariableValue(req.BasicAuth.Password, variableMap)

		req.BasicAuth.Username = agentExec.ReplaceExpressionValue(req.BasicAuth.Username, variableMap, expressionValueMap)
		req.BasicAuth.Password = agentExec.ReplaceExpressionValue(req.BasicAuth.Password, variableMap, expressionValueMap)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = agentExec.ReplaceVariableValue(req.BearerToken.Token, variableMap)
		req.BearerToken.Token = agentExec.ReplaceExpressionValue(req.BearerToken.Token, variableMap, expressionValueMap)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = agentExec.ReplaceVariableValue(req.OAuth20.Name, variableMap)
		req.OAuth20.CallbackUrl = agentExec.ReplaceVariableValue(req.OAuth20.CallbackUrl, variableMap)
		req.OAuth20.AuthURL = agentExec.ReplaceVariableValue(req.OAuth20.AuthURL, variableMap)
		req.OAuth20.AccessTokenURL = agentExec.ReplaceVariableValue(req.OAuth20.AccessTokenURL, variableMap)
		req.OAuth20.ClientID = agentExec.ReplaceVariableValue(req.OAuth20.ClientID, variableMap)
		req.OAuth20.Scope = agentExec.ReplaceVariableValue(req.OAuth20.Scope, variableMap)

		req.OAuth20.Name = agentExec.ReplaceExpressionValue(req.OAuth20.Name, variableMap, expressionValueMap)
		req.OAuth20.CallbackUrl = agentExec.ReplaceExpressionValue(req.OAuth20.CallbackUrl, variableMap, expressionValueMap)
		req.OAuth20.AuthURL = agentExec.ReplaceExpressionValue(req.OAuth20.AuthURL, variableMap, expressionValueMap)
		req.OAuth20.AccessTokenURL = agentExec.ReplaceExpressionValue(req.OAuth20.AccessTokenURL, variableMap, expressionValueMap)
		req.OAuth20.ClientID = agentExec.ReplaceExpressionValue(req.OAuth20.ClientID, variableMap, expressionValueMap)
		req.OAuth20.Scope = agentExec.ReplaceExpressionValue(req.OAuth20.Scope, variableMap, expressionValueMap)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = agentExec.ReplaceVariableValue(req.ApiKey.Key, variableMap)
		req.ApiKey.Value = agentExec.ReplaceVariableValue(req.ApiKey.Value, variableMap)
		req.ApiKey.TransferMode = agentExec.ReplaceVariableValue(req.ApiKey.TransferMode, variableMap)

		req.ApiKey.Key = agentExec.ReplaceExpressionValue(req.ApiKey.Key, variableMap, expressionValueMap)
		req.ApiKey.Value = agentExec.ReplaceExpressionValue(req.ApiKey.Value, variableMap, expressionValueMap)
		req.ApiKey.TransferMode = agentExec.ReplaceExpressionValue(req.ApiKey.TransferMode, variableMap, expressionValueMap)
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
