package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
)

func InitDebugExecContext(execUuid string) {
	SetScopedVariables(execUuid, map[uint][]domain.ExecVariable{})

	return
}

func InitScenarioExecContext(execObj *ScenarioExecObj) {
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
