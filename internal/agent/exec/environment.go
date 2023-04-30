package agentExec

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
)

func GenRequestUrl(req *v1.BaseRequest, interfaceId uint) {
	envId := InterfaceToEnvMap[interfaceId]
	mp := EnvToVariablesMap[envId]

	baseUrl := mp[consts.KEY_BASE_URL]["value"].(string)

	req.Url = _httpUtils.AddSepIfNeeded(baseUrl) + req.Url
}
