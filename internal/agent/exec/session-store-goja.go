package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func (s *ExecSession) GetGojaVariables() (ret *[]domain.ExecVariable) {
	ret = s._gojaVariables
	return
}
func (s *ExecSession) SetGojaVariables(val *[]domain.ExecVariable) {
	s._gojaVariables = val
	return
}
func (s *ExecSession) ResetGojaVariables() {
	s._gojaVariables = &[]domain.ExecVariable{}
	return
}
func (s *ExecSession) AppendGojaVariables(val domain.ExecVariable) {
	*s._gojaVariables = append(*s._gojaVariables, val)
	return
}

func (s *ExecSession) GetGojaLogs() (ret *[]string) {
	ret = s._gojaLogs
	return
}
func (s *ExecSession) SetGojaLogs(val *[]string) {
	s._gojaLogs = val
	return
}
func (s *ExecSession) ResetGojaLogs() {
	s._gojaLogs = &[]string{}
	return
}
func (s *ExecSession) AppendGojaLog(item string) {
	*s._gojaLogs = append(*s._gojaLogs, item)
	return
}
