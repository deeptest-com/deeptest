package execHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/kataras/iris/v12"
)

func ReplaceExecVariablesForInvocation(req *serverDomain.InvocationRequest, variables []domain.ExecVariable) (err error) {
	variableArr := genComplexVariableArr(variables)
	requestHelper.ReplaceAll(req, variableArr)

	return
}

// use variables from testdata processor like ${user.email}, its value is a map.
func genComplexVariableArr(variables []domain.ExecVariable) (ret [][]string) {

	variableMap := iris.Map{}
	for _, item := range variables {
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
