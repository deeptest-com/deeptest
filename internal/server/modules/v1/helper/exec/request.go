package execHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/kataras/iris/v12"
)

func ReplaceVariablesForInvocation(req *serverDomain.InvocationRequest, variables []domain.Variable) (err error) {
	variableArr := genVariableArr(variables)
	requestHelper.ReplaceAll(req, variableArr)

	return
}

func genVariableArr(variables []domain.Variable) (
	ret [][]string) {

	variableMap := iris.Map{}
	for _, item := range variables {
		variableMap[item.Name] = item.Value
	}

	for key, val := range variableMap {
		ret = append(ret, []string{fmt.Sprintf("${%s}", key), fmt.Sprintf("%v", val)})
	}

	return
}
