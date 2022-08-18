package business

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

type ExecHelperService struct {
}

func NewExecHelperService() *ExecHelperService {
	return &ExecHelperService{}
}

func (s *ExecHelperService) IsLoopTimes(containerLog *domain.Log) bool {
	return containerLog.Output.Times > 0
}

func (s *ExecHelperService) IsLoopRange(containerLog *domain.Log) bool {
	return containerLog.Output.Range != ""
}

func (s *ExecHelperService) IsWrapperProcessor(category consts.ProcessorCategory) bool {
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

func (s *ExecHelperService) IsExecutableWrapperProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}

func (s *ExecHelperService) IsActionProcessor(category consts.ProcessorCategory) bool {
	arr := []string{
		consts.ProcessorTimer.ToString(),
		consts.ProcessorVariable.ToString(),
		consts.ProcessorAssertion.ToString(),
		consts.ProcessorExtractor.ToString(),
		consts.ProcessorCookie.ToString(),
	}
	return _stringUtils.FindInArr(category.ToString(), arr)
}
