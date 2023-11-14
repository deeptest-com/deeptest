package agentExec

import (
	"fmt"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"strings"
)

func ReplaceVariableValueInBody(value string, execUuid string) (ret string) {
	// add a plus to set field vale as a number
	// {"id": "${+dev_env_var1}"} => {"id": 2}

	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value)
	ret = value

	for _, placeholder := range variablePlaceholders {
		oldVal := fmt.Sprintf("${%s}", placeholder)
		if strings.Index(placeholder, "+") == 0 { // replace it with a number, if has prefix +
			oldVal = "\"" + oldVal + "\""
		}

		placeholderWithoutPlus := strings.TrimLeft(placeholder, "+")
		newVal := getPlaceholderVariableValue(placeholderWithoutPlus, execUuid)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func ReplaceVariableValue(value string, execUuid string) (ret string) {
	ret = value
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value)

	for _, placeholder := range variablePlaceholders {
		oldVal := fmt.Sprintf("${%s}", placeholder)

		placeholderWithoutPlus := strings.TrimLeft(placeholder, "+")
		newVal := getPlaceholderVariableValue(placeholderWithoutPlus, execUuid)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func getPlaceholderVariableValue(name string, execUuid string) (ret string) {
	typ := getPlaceholderType(name)

	if typ == consts.PlaceholderTypeVariable {
		variable, _ := GetVariable(GetCurrScenarioProcessorId(execUuid), name, execUuid)
		ret = valueUtils.InterfaceToStr(variable.Value)

	} else if typ == consts.PlaceholderTypeDatapool {
		ret = getDatapoolValue(name, execUuid)
	}
	//else if typ == consts.PlaceholderTypeFunction {
	//}

	return
}

func getPlaceholderType(placeholder string) (ret consts.PlaceholderType) {
	if strings.HasPrefix(placeholder, consts.PlaceholderPrefixDatapool.String()) {
		return consts.PlaceholderTypeDatapool
	} else if strings.HasPrefix(placeholder, consts.PlaceholderPrefixFunction.String()) {
		return consts.PlaceholderTypeFunction
	}

	return consts.PlaceholderTypeVariable
}
