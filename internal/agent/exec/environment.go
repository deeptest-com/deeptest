package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"strings"
)

func GenRequestUrl(req *domain.BaseRequest, debugInterfaceId uint, baseUrl string) {
	envId := ExecScene.DebugInterfaceToEnvMap[debugInterfaceId]
	vars := ExecScene.EnvToVariables[envId]

	if baseUrl == "" {
		baseUrl = getValueFromList(consts.KEY_BASE_URL, vars)
	}

	uri := ReplacePathParams(req.Url, req.PathParams)

	//req.Url = _httpUtils.AddSepIfNeeded(baseUrl) + uri
	req.Url = baseUrl + uri
}

func ReplacePathParams(uri string, pathParams []domain.Param) string {

	for _, param := range pathParams {
		if param.ParamIn != consts.ParamInPath {
			continue
		}

		vari := fmt.Sprintf("{%v}", param.Name)

		uri = strings.ReplaceAll(uri, vari, param.Value)
	}

	return uri
}
