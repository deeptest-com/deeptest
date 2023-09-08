package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"strings"
)

var (
	ForceStopExec = false

	ServerUrl     = ""
	ServerApiPath = "api/v1"
	ServerToken   = ""

	CurrRequest domain.BaseRequest

	CurrScenarioProcessor   *Processor
	CurrScenarioProcessorId = uint(0)
	CurrDebugInterfaceId    = uint(0)

	//CachedShareVarByProcessorForRead map[uint]domain.VarKeyValuePair

	ScopedVariables = map[uint][]domain.ExecVariable{} // for scenario and debug
	ScopedCookies   = map[uint][]domain.ExecCookie{}   // only for scenario
	ScopeHierarchy  = map[uint]*[]uint{}               // only for scenario (processId -> ancestorProcessIds)

	ExecScene      = domain.ExecScene{} // for scenario and debug, from server
	DatapoolCursor = map[string]int{}   // only for scenario
)

func InitDebugExecContext() (variables []domain.ExecVariable) {
	ScopedVariables = map[uint][]domain.ExecVariable{}

	return
}

func InitScenarioExecContext(execObj *ScenarioExecObj) (variables []domain.ExecVariable) {
	ComputerScopeHierarchy(execObj.RootProcessor, &ScopeHierarchy)

	ExecScene = execObj.ExecScene
	DatapoolCursor = map[string]int{}

	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies = map[uint][]domain.ExecCookie{}

	return
}

func GetValidScopeIds(id uint) (ret *[]uint) {
	ret = &[]uint{}

	if id == 0 {
		*ret = append(*ret, id)
		return
	}

	ret = ScopeHierarchy[id]

	return
}

// like {name.prop}
func EvaluateVariablePropExpressionValue(variable domain.ExecVariable, propExpression string) (
	ret domain.ExecVariable, ok bool) {
	arr := strings.Split(propExpression, ".")
	variableName := arr[0]

	if variable.Name == variableName {
		ret = variable
		ret.Name = propExpression // set name from item to item.a

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
