package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
	"time"
)

var (
	ScopeHierarchy  = map[uint]*[]uint{}
	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies   = map[uint][]domain.ExecCookie{}
)

func InitScopeHierarchy(processors []*Processor) (variables []domain.ExecVariable) {
	GetScopeHierarchy(processors, &ScopeHierarchy)

	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies = map[uint][]domain.ExecCookie{}

	return
}

func ListCachedVariable(scopeId uint) (variables []domain.ExecVariable) {
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		for _, vari := range ScopedVariables[id] {
			if !vari.IsShare && id != scopeId {
				continue
			}

			variables = append(variables, vari)
		}
	}

	return
}

func GetVariable(scopeId uint, variablePath string) (variable domain.ExecVariable, err error) {
	if variablePath == "var1" {
		logUtils.Info("")
	}

	allValidIds := ScopeHierarchy[scopeId]
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

func SetVariable(scopeId uint, variableName string, variableValue interface{}, isShare bool) (
	err error) {

	found := false

	newVariable := domain.ExecVariable{
		Name:    variableName,
		Value:   variableValue,
		IsShare: isShare,
	}

	allValidIds := ScopeHierarchy[scopeId]
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
		ScopedVariables[scopeId] = append(ScopedVariables[scopeId], newVariable)
	}

	return
}

func ClearVariable(scopeId uint, variableName string) (err error) {
	deleteIndex := -1

	targetScopeId := uint(0)

	allValidIds := ScopeHierarchy[scopeId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[scopeId] {
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
		ScopedVariables[scopeId] = append(
			ScopedVariables[targetScopeId][:deleteIndex], ScopedVariables[scopeId][(deleteIndex+1):]...)
	}

	return
}

func ListCookie(scopeId uint) (cookies []domain.ExecCookie) {
	allValidIds := ScopeHierarchy[scopeId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[scopeId] {
			cookies = append(cookies, ScopedCookies[id]...)
		}
	}

	return
}

func GetCookie(scopeId uint, cookieName, domain string) (cookie domain.ExecCookie) {
	allValidIds := ScopeHierarchy[scopeId]
	if allValidIds != nil {
		for _, id := range *ScopeHierarchy[scopeId] {
			for _, item := range ScopedCookies[id] {
				if item.Name == cookieName && item.Domain == domain && item.ExpireTime.Unix() > time.Now().Unix() {
					cookie = item

					goto LABEL
				}
			}
		}
	}

LABEL:

	return
}

func SetCookie(scopeId uint, cookieName string, cookieValue interface{}, domainName string, expireTime *time.Time) (err error) {
	found := false

	newCookie := domain.ExecCookie{
		Name:  cookieName,
		Value: cookieValue,

		Domain:     domainName,
		ExpireTime: expireTime,
	}

	for i := 0; i < len(ScopedCookies[scopeId]); i++ {
		if ScopedCookies[scopeId][i].Name == cookieName {
			ScopedCookies[scopeId][i] = newCookie

			found = true
			break
		}
	}

	if !found {
		ScopedCookies[scopeId] = append(ScopedCookies[scopeId], newCookie)
	}

	return
}

func ClearCookie(scopeId uint, cookieName string) (err error) {
	deleteIndex := -1
	for index, item := range ScopedCookies[scopeId] {
		if item.Name == cookieName {
			deleteIndex = index
			break
		}
	}

	if deleteIndex > -1 {
		ScopedCookies[scopeId] = append(
			ScopedCookies[scopeId][:deleteIndex], ScopedCookies[scopeId][(deleteIndex+1):]...)
	}

	return
}

func GetScopeHierarchy(processors []*Processor, scopeHierarchyMap *map[uint]*[]uint) {
	childToParentIdMap := map[uint]uint{}
	for _, processor := range processors {
		childToParentIdMap[processor.ID] = processor.ParentId
	}

	for childId, parentId := range childToParentIdMap {
		if (*scopeHierarchyMap)[childId] == nil {
			arr := []uint{childId}
			(*scopeHierarchyMap)[childId] = &arr
		}
		*(*scopeHierarchyMap)[childId] = append(*(*scopeHierarchyMap)[childId], parentId)

		addSuperParent(childId, parentId, childToParentIdMap, scopeHierarchyMap)
	}
}

func addSuperParent(id, parentId uint, childToParentIdMap map[uint]uint, scopeHierarchyMap *map[uint]*[]uint) {
	superId, ok := childToParentIdMap[parentId]
	if ok {
		*(*scopeHierarchyMap)[id] = append(*(*scopeHierarchyMap)[id], superId)

		addSuperParent(id, superId, childToParentIdMap, scopeHierarchyMap)
	}
}
