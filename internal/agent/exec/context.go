package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"strings"
	"time"
)

var (
	CurrProcessorId            = uint(0)
	CachedVariablesByProcessor map[uint]domain.VarKeyValuePair

	ScopeHierarchy  = map[uint]*[]uint{}               // only for scenario
	ScopedVariables = map[uint][]domain.ExecVariable{} // only for scenario
	ScopedCookies   = map[uint][]domain.ExecCookie{}   // only for scenario

	// global variables and params
	GlobalEnvVars   []domain.GlobalEnvVar
	GlobalParamVars []domain.GlobalParamVar

	// datapool
	DatapoolData   = domain.Datapools{}
	DatapoolCursor = map[string]int{} // only for scenario

	// env variables
	InterfaceToEnvMap map[uint]uint
	EnvToVariablesMap map[uint]map[string]domain.VarKeyValuePair
)

func InitExecContext(execObj *ScenarioExecObj) (variables []domain.ExecVariable) {
	GetScopeHierarchy(execObj.RootProcessor, &ScopeHierarchy)
	DatapoolData = execObj.Datapools

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
