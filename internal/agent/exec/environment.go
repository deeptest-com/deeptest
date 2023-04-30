package agentExec

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
)

func GenRequestUrl(req *v1.BaseRequest, interfaceId uint) {
	envId := InterfaceToEnvMap[interfaceId]
	vars := EnvToVariablesMap[envId]

	baseUrl := getValueFromList(consts.KEY_BASE_URL, vars)

	req.Url = _httpUtils.AddSepIfNeeded(baseUrl) + req.Url
}
