package business

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"time"
)

var (
	ExecLog = domain.Log{}

	ScopeHierarchy  = map[uint]*[]uint{}
	ScopedVariables = map[uint][]domain.Variable{}
	ScopedCookies   = map[uint][]domain.Cookie{}
)

type ExecContextService struct {
	ScenarioNodeRepo *repo.ScenarioNodeRepo `inject:""`
}

func NewExecContextService() *ExecContextService {
	return &ExecContextService{}
}

func (s *ExecContextService) InitScopeHierarchy(scenarioId uint) (variables []domain.Variable) {
	s.ScenarioNodeRepo.GetScopeHierarchy(scenarioId, &ScopeHierarchy)
	return
}

func (s *ExecContextService) ListVariable(scopeId uint) (variables []domain.Variable) {
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		variables = append(variables, ScopedVariables[id]...)
	}

	return
}

func (s *ExecContextService) GetVariable(scopeId uint, variableName string) (variable domain.Variable) {
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		for _, item := range ScopedVariables[id] {
			if item.Name == variableName {
				variable = item
				goto LABEL
			}
		}
	}

LABEL:

	return
}

func (s *ExecContextService) SetVariable(scopeId uint, variableName string, variableValue interface{}) (err error) {
	found := false

	newVariable := domain.Variable{
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

func (s *ExecContextService) ListCookie(scopeId uint) (cookies []domain.Cookie) {
	cookies = ScopedCookies[scopeId]
	return
}

func (s *ExecContextService) GetCookie(scopeId uint, cookieName string) (cookie domain.Cookie) {
	effectiveScopeIds := ScopeHierarchy[scopeId]

	for _, id := range *effectiveScopeIds {
		for _, item := range ScopedCookies[id] {
			if item.Name == cookieName {
				cookie = item

				goto LABEL
			}
		}
	}

LABEL:

	return
}

func (s *ExecContextService) SetCookie(scopeId uint, cookieName, cookieValue, domainName string, expireTime time.Time) (err error) {
	found := false

	newCookie := domain.Cookie{
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
