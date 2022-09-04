package business

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

type ExecComm struct {
}

func NewExecHelperService() *ExecComm {
	return &ExecComm{}
}

func (s *ExecComm) IsLoop(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorCategory == consts.ProcessorLoop
}
func (s *ExecComm) IsLoopTimesPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopTime && containerLog.Output.Times > 0
}
func (s *ExecComm) IsLoopUntilPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopUntil && containerLog.Output.Expression != ""
}
func (s *ExecComm) IsLoopInPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopRange && containerLog.Output.List != ""
}
func (s *ExecComm) IsLoopRangePass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopRange && containerLog.Output.Range != ""
}
func (s *ExecComm) IsLoopLoopBreak(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopBreak
}

func (s *ExecComm) IsLogicPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorCategory == consts.ProcessorLogic && containerLog.Output.Pass
}

func (s *ExecComm) IsDataPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorCategory == consts.ProcessorData && containerLog.Output.Url != ""
}

func (s *ExecComm) IsWrapperProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		consts.ProcessorRoot.ToString(),
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ExecComm) IsExecutableWrapperProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ExecComm) IsActionProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		consts.ProcessorTimer.ToString(),
		consts.ProcessorVariable.ToString(),
		consts.ProcessorAssertion.ToString(),
		consts.ProcessorExtractor.ToString(),
		consts.ProcessorCookie.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}
