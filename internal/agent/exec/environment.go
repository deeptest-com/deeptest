package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"strings"
)

func GenRequestUrl(req *domain.BaseRequest, debugInterfaceId uint, baseUrl string) {
	envId := ExecScene.DebugInterfaceToEnvMap[debugInterfaceId]
	vars := ExecScene.EnvToVariables[envId]

	if baseUrl == "" {
		baseUrl = getValueFromList(consts.KEY_BASE_URL, vars)
	}

	req.Url = ReplacePathParams(req.Url, req.PathParams)

	if req.ProcessorInterfaceSrc != consts.DiagnoseDebug {
		req.Url = _httpUtils.CombineUrls(baseUrl, req.Url)
	}
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
