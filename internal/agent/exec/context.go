package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"strings"
)

var (
	ServerUrl     = ""
	ServerApiPath = "api/v1"
	ServerToken   = ""

	CurrProcessorId = uint(0)
	CurrInterfaceId = uint(0)

	CachedShareVarByProcessorForRead map[uint]domain.VarKeyValuePair

	ScopeHierarchy  = map[uint]*[]uint{}               // only for scenario
	ScopedVariables = map[uint][]domain.ExecVariable{} // only for scenario
	ScopedCookies   = map[uint][]domain.ExecCookie{}   // only for scenario

	ExecScene      = domain.ExecScene{}
	DatapoolCursor = map[string]int{}
)

func InitExecContext(execObj *ScenarioExecObj) (variables []domain.ExecVariable) {
	ComputerScopeHierarchy(execObj.RootProcessor, &ScopeHierarchy)

	ExecScene = execObj.ExecScene

	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies = map[uint][]domain.ExecCookie{}

	return
}

func GetCachedVariableMapInContext(processorId uint) (ret domain.VarKeyValuePair) {
	ret = domain.VarKeyValuePair{}

	variables := listCachedVariable(processorId)

	for _, item := range variables {
		valMap, isMap := item.Value.(domain.VarKeyValuePair)

		if isMap {
			for propKey, v := range valMap {
				ret[fmt.Sprintf("%s.%s", item.Name, propKey)] = v
			}
		} else {
			ret[item.Name] = item.Value
		}
	}

	return
}
func listCachedVariable(processorId uint) (variables []domain.ExecVariable) {
	effectiveScopeIds := ScopeHierarchy[processorId]

	if effectiveScopeIds == nil {
		return
	}

	for _, id := range *effectiveScopeIds {
		for _, vari := range ScopedVariables[id] {
			if vari.Scope == consts.Public || (vari.Scope == consts.Private && id == processorId) {

				variables = append(variables, vari)
			}
		}
	}

	return
}

func EvaluateVariableExpressionValue(variable domain.ExecVariable, variablePath string) (
	ret domain.ExecVariable, ok bool) {
	arr := strings.Split(variablePath, ".")
	variableName := arr[0]

	if variable.Name == variableName {
		ret = variable

		if len(arr) > 1 {
			variableProp := arr[1]
			ret.Value = variable.Value.(domain.VarKeyValuePair)[variableProp]
		}

		ok = true

	}

	return
}

func ComputerScopeHierarchy(processor *Processor, scopeHierarchyMap *map[uint]*[]uint) {
	processors := make([]*Processor, 0)
	GetProcessorList(processor, &processors)

	childToParentIdMap := map[uint]uint{}
	for _, processor := range processors {
		childToParentIdMap[processor.ID] = processor.ParentId
	}

	for childId, parentId := range childToParentIdMap {
		if (*scopeHierarchyMap)[childId] == nil {
			arr := []uint{childId}
			(*scopeHierarchyMap)[childId] = &arr
		}

		if !_intUtils.FindUintInArr(parentId, *(*scopeHierarchyMap)[childId]) {
			*(*scopeHierarchyMap)[childId] = append(*(*scopeHierarchyMap)[childId], parentId)
		}

		addSuperParent(childId, parentId, childToParentIdMap, scopeHierarchyMap)
	}
}

func GetProcessorList(processor *Processor, list *[]*Processor) {
	*list = append(*list, processor)

	for _, child := range processor.Children {
		GetProcessorList(child, list)
	}

	return
}

func addSuperParent(id, parentId uint, childToParentIdMap map[uint]uint, scopeHierarchyMap *map[uint]*[]uint) {
	superId, ok := childToParentIdMap[parentId]
	if ok {
		if !_intUtils.FindUintInArr(superId, *(*scopeHierarchyMap)[id]) {
			*(*scopeHierarchyMap)[id] = append(*(*scopeHierarchyMap)[id], superId)
		}

		addSuperParent(id, superId, childToParentIdMap, scopeHierarchyMap)
	}
}
