package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func (s *ExecSession) GetGojaVariables() (ret *[]domain.ExecVariable) {
	ret = s.GojaVariables
	return
}
func (s *ExecSession) SetGojaVariables(val *[]domain.ExecVariable) {
	s.GojaVariables = val
	return
}
func (s *ExecSession) ResetGojaVariables() {
	s.GojaVariables = &[]domain.ExecVariable{}
	return
}
func (s *ExecSession) AppendGojaVariables(val domain.ExecVariable) {
	*s.GojaVariables = append(*s.GojaVariables, val)
	return
}

func (s *ExecSession) GetGojaLogs() (ret *[]string) {
	ret = s.GojaLogs
	return
}
func (s *ExecSession) SetGojaLogs(val *[]string) {
	s.GojaLogs = val
	return
}
func (s *ExecSession) ResetGojaLogs() {
	s.GojaLogs = &[]string{}
	return
}
func (s *ExecSession) AppendGojaLog(item string) {
	*s.GojaLogs = append(*s.GojaLogs, item)
	return
}
