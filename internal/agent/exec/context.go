package agentExec

import (
	queryUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
)

func InitDebugExecContext(execUuid string) (variables []domain.ExecVariable) {
	SetScopedVariables(execUuid, map[uint][]domain.ExecVariable{})

	return
}

func InitScenarioExecContext(execObj *ScenarioExecObj) (variables []domain.ExecVariable) {
	execUuid := execObj.ExecUuid

	scopeHierarchy := map[uint]*[]uint{}
	ComputerScopeHierarchy(execObj.RootProcessor, &scopeHierarchy)
	SetScopeHierarchy(execUuid, scopeHierarchy)

	SetExecScene(execUuid, execObj.ExecScene)
	SetDatapoolCursor(execUuid, map[string]int{})

	SetScopedVariables(execUuid, map[uint][]domain.ExecVariable{})
	SetScopedCookies(execUuid, map[uint][]domain.ExecCookie{})

	return
}

func GetValidScopeIds(processorId uint, execUuid string) (ret *[]uint) {
	if processorId == 0 { // return an arr with single 0
		arr := []uint{processorId}
		ret = &arr
		return
	}

	scopeHierarchy := GetScopeHierarchy(execUuid)
	ret = scopeHierarchy[processorId]

	return
}

// like {name.prop}
func EvaluateVariablePropExpressionValue(variable domain.ExecVariable, propExpression string) (
	ret domain.ExecVariable, ok bool) {

	variableName, jsonPathExpression, isJsonPath := commUtils.IsJsonPathExpression(propExpression)

	if variable.Name == variableName {
		ret = variable
		ret.Name = propExpression // set name from item to item.a

		ok = true

		if isJsonPath {
			jsn, _ := commUtils.ConvertValueForStore(variable.Value)

			var err error
			ret.Value, ret.ValueType, err =
				queryUtils.JsonPath(jsn, jsonPathExpression)

			if err != nil {
				ok = false
			}
		}
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
