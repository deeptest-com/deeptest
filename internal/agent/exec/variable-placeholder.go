package agentExec

import (
	"fmt"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"strings"
)

func ReplaceVariableValue(value string) (ret string) {
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value)
	ret = value

	for _, placeholder := range variablePlaceholders {
		oldVal := fmt.Sprintf("${%s}", placeholder)
		newVal := getPlaceholderVariableValue(placeholder)

		ret = strings.ReplaceAll(ret, oldVal, newVal)
	}

	return
}

func getPlaceholderVariableValue(name string) (ret string) {
	typ := getPlaceholderType(name)

	if typ == consts.PlaceholderTypeVariable {
		variable, _ := GetVariable(CurrScenarioProcessorId, name)
		ret = valueUtils.InterfaceToStr(variable.Value)

	} else if typ == consts.PlaceholderTypeDatapool {
		ret = getDatapoolValue(name)
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
