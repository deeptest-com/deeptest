package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
)

func GenRequestUrl(req *domain.BaseRequest, interfaceId uint, baseUrl string) {
	envId := ExecScene.InterfaceToEnvMap[interfaceId]
	vars := ExecScene.EnvToVariables[envId]

	if baseUrl == "" {
		baseUrl = getValueFromList(consts.KEY_BASE_URL, vars)
	}

	req.Url = _httpUtils.AddSepIfNeeded(baseUrl) + req.Url
}
