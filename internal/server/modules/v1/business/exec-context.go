package business

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

var (
	ExecLog = domain.ExecLog{}

	ScopeHierarchy  = map[uint]*[]uint{}
	ScopedVariables = map[uint][]domain.ExecVariable{}
	ScopedCookies   = map[uint][]domain.ExecCookie{}
)

type ExecContext struct {
	ScenarioNodeRepo *repo.ScenarioNodeRepo `inject:""`
}

func NewExecContextService() *ExecContext {
	return &ExecContext{}
}

func (s *ExecContext) InitScopeHierarchy(scenarioId uint) (variables []domain.ExecVariable) {
	s.ScenarioNodeRepo.GetScopeHierarchy(scenarioId, &ScopeHierarchy)
	return
}

func (s *ExecContext) ListVariable(scopeId uint) (variables []domain.ExecVariable) {
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		variables = append(variables, ScopedVariables[id]...)
	}

	return
}

func (s *ExecContext) GetVariable(scopeId uint, variableName string) (variable domain.ExecVariable, err error) {
	if variableName == "var1" {
		logUtils.Info("")
	}
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		for _, item := range ScopedVariables[id] {
			if item.Name == variableName {
				variable = item
				goto LABEL
			}
		}
	}

	if variable.Name == "" {
		err = errors.New(fmt.Sprintf("找不到变量%s", variableName))
	}

LABEL:

	return
}

func (s *ExecContext) SetVariable(scopeId uint, variableName string, variableValue interface{}) (err error) {
	found := false

	newVariable := domain.ExecVariable{
		Name:  variableName,
		Value: variableValue,
	}

	for i := 0; i < len(ScopedVariables[scopeId]); i++ {
		if ScopedVariables[scopeId][i].Name == variableName {
			ScopedVariables[scopeId][i] = newVariable

			found = true
			break
		}
	}

	if !found {
		ScopedVariables[scopeId] = append(ScopedVariables[scopeId], newVariable)
	}

	return
}

func (s *ExecContext) ClearVariable(scopeId uint, variableName string) (err error) {
	deleteIndex := -1
	for index, item := range ScopedVariables[scopeId] {
		if item.Name == variableName {
			deleteIndex = index
			break
		}
	}

	if deleteIndex > -1 {
		ScopedVariables[scopeId] = append(
			ScopedVariables[scopeId][:deleteIndex], ScopedVariables[scopeId][(deleteIndex+1):]...)
	}

	return
}

func (s *ExecContext) ListCookie(scopeId uint) (cookies []domain.ExecCookie) {
	cookies = ScopedCookies[scopeId]
	return
}

func (s *ExecContext) GetCookie(scopeId uint, cookieName, domain string) (cookie domain.ExecCookie) {
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		for _, item := range ScopedCookies[id] {
			if item.Name == cookieName && item.Domain == domain && item.ExpireTime.Unix() > time.Now().Unix() {
				cookie = item

				goto LABEL
			}
		}
	}

LABEL:

	return
}

func (s *ExecContext) SetCookie(scopeId uint, cookieName string, cookieValue interface{}, domainName string, expireTime *time.Time) (err error) {
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

func (s *ExecContext) ClearCookie(scopeId uint, cookieName string) (err error) {
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
