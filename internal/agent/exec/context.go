package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"strings"
	"time"
)

var (
	ScopeHierarchy  = map[uint]*[]uint{}
	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies   = map[uint][]domain.ExecCookie{}
)

func InitScopeHierarchy(processor *Processor) (variables []domain.ExecVariable) {
	GetScopeHierarchy(processor, &ScopeHierarchy)

	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies = map[uint][]domain.ExecCookie{}

	return
}

func ListCachedVariable(processorId uint) (variables []domain.ExecVariable) {
	effectiveScopeIds := ScopeHierarchy[processorId]

	if effectiveScopeIds == nil {
		return
	}

	for _, id := range *effectiveScopeIds {
		for _, vari := range ScopedVariables[id] {
			if vari.Scope == consts.Global || vari.Scope == consts.Local ||
				(vari.Scope == consts.Private && id == processorId) {

				variables = append(variables, vari)
			}
		}
	}

	return
}
func GetVariableMap(processorId uint) (ret map[string]interface{}) {
	ret = map[string]interface{}{}

	variables := ListCachedVariable(processorId)

	for _, item := range variables {
		valMap, isMap := item.Value.(map[string]interface{})

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

func GetVariable(processorId uint, variablePath string) (variable domain.ExecVariable, err error) {
	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *allValidIds {
			for _, item := range ScopedVariables[id] {
				var ok bool
				if variable, ok = EvaluateVariableExpressionValue(item, variablePath); ok {
					goto LABEL
				}
			}
		}
	}

	if variable.Name == "" { // not found
		err = errors.New(fmt.Sprintf("找不到变量\"%s\"", variablePath))
	}

LABEL:

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
			ret.Value = variable.Value.(map[string]interface{})[variableProp]
		}

		ok = true

	}

	return
}

func ImportVariables(processorId uint, variables []domain.Variable, scope consts.ExtractorScope) (err error) {
	for _, item := range variables {
		newVariable := domain.ExecVariable{
			Name:  item.Name,
			Value: item.Value,
			Scope: scope,
		}

		found := false
		for i := 0; i < len(ScopedVariables[processorId]); i++ {
			if ScopedVariables[processorId][i].Name == item.Name {
				ScopedVariables[processorId][i] = newVariable

				found = true
				break
			}
		}

		if !found {
			ScopedVariables[processorId] = append(ScopedVariables[processorId], newVariable)
		}
	}

	return
}

func SetVariable(processorId uint, variableName string, variableValue interface{}, scope consts.ExtractorScope) (
	err error) {

	found := false

	newVariable := domain.ExecVariable{
		Name:  variableName,
		Value: variableValue,
		Scope: scope,
	}

	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *allValidIds {
			for i := 0; i < len(ScopedVariables[id]); i++ {
				if ScopedVariables[id][i].Name == variableName {
					ScopedVariables[id][i] = newVariable

					found = true
					break
				}
			}
		}
	}

	if !found {
		ScopedVariables[processorId] = append(ScopedVariables[processorId], newVariable)
	}

	return
}

func ClearVariable(processorId uint, variableName string) (err error) {
	deleteIndex := -1

	targetScopeId := uint(0)

	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[processorId] {
			for index, item := range ScopedVariables[id] {
				if item.Name == variableName {
					deleteIndex = index
					targetScopeId = id
					break
				}
			}
		}
	}

	if deleteIndex > -1 {
		if len(ScopedVariables[targetScopeId]) == deleteIndex+1 {
			ScopedVariables[targetScopeId] = make([]domain.ExecVariable, 0)
		} else {
			ScopedVariables[targetScopeId] = append(
				ScopedVariables[targetScopeId][:deleteIndex], ScopedVariables[targetScopeId][(deleteIndex+1):]...)
		}
	}

	return
}

func ListCookie(processorId uint) (cookies []domain.ExecCookie) {
	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[processorId] {
			cookies = append(cookies, ScopedCookies[id]...)
		}
	}

	return
}

func GetCookie(processorId uint, cookieName, domain string) (cookie domain.ExecCookie) {
	allValidIds := ScopeHierarchy[processorId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[processorId] {
			for _, item := range ScopedCookies[id] {
				if item.Name == cookieName && (item.Domain == "" || domain == "" || item.Domain == domain) &&
					(item.ExpireTime == nil || item.ExpireTime.Unix() > time.Now().Unix()) {
					cookie = item

					goto LABEL
				}
			}
		}
	}

LABEL:

	return
}

func SetCookie(processorId uint, cookieName string, cookieValue interface{}, domainName string, expireTime *time.Time) (err error) {
	found := false

	newCookie := domain.ExecCookie{
		Name:  cookieName,
		Value: cookieValue,

		Domain:     domainName,
		ExpireTime: expireTime,
	}

	for i := 0; i < len(ScopedCookies[processorId]); i++ {
		if ScopedCookies[processorId][i].Name == cookieName {
			ScopedCookies[processorId][i] = newCookie

			found = true
			break
		}
	}

	if !found {
		ScopedCookies[processorId] = append(ScopedCookies[processorId], newCookie)
	}

	return
}

func ClearCookie(processorId uint, cookieName string) (err error) {
	deleteIndex := -1
	for index, item := range ScopedCookies[processorId] {
		if item.Name == cookieName {
			deleteIndex = index
			break
		}
	}

	if deleteIndex > -1 {
		ScopedCookies[processorId] = append(
			ScopedCookies[processorId][:deleteIndex], ScopedCookies[processorId][(deleteIndex+1):]...)
	}

	return
}

func GetScopeHierarchy(processor *Processor, scopeHierarchyMap *map[uint]*[]uint) {
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
