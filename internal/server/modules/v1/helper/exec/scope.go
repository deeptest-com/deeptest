package execHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

var (
	ScopeHierarchyMap = map[uint]*[]uint{}
	ScopeVariableMap  = map[uint][]domain.Variable{}
)

func GetScopedVariable(scopeId uint) (variables []domain.Variable) {

	return
}

func SetScopedVariable(scopeId uint, variableId uint) (err error) {

	return
}
