package business

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
)

var (
	executableContainerProcessors = []string{
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}

	noExecutableContainerProcessors = []string{
		consts.ProcessorRoot.ToString(),
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorGroup.ToString(),
	}

	actionProcessors = []string{
		consts.ProcessorTimer.ToString(),
		consts.ProcessorPrint.ToString(),
		consts.ProcessorVariable.ToString(),
		consts.ProcessorAssertion.ToString(),
		consts.ProcessorExtractor.ToString(),
		consts.ProcessorCookie.ToString(),
	}
)

type ExecComm struct {
}

func (s *ExecComm) IsLoopPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorCategory == consts.ProcessorLoop
}
func (s *ExecComm) IsLoopTimesPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopTime && containerLog.Output.Times > 0
}
func (s *ExecComm) IsLoopUntilPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopUntil && containerLog.Output.Expression != ""
}
func (s *ExecComm) IsLoopInPass(containerLog *domain.ExecLog) bool {
	return containerLog.ProcessorType == consts.ProcessorLoopIn && containerLog.Output.List != ""
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

func (s *ExecComm) IsNoExecutableContainerProcessor(processor *model.Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), noExecutableContainerProcessors)
}

func (s *ExecComm) IsExecutableContainerProcessor(processor *model.Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), executableContainerProcessors) &&
		processor.EntityType != consts.ProcessorLoopBreak
}

func (s *ExecComm) IsActionProcessor(processor *model.Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), actionProcessors) ||
		processor.EntityType == consts.ProcessorLoopBreak
}

func (s *ExecComm) IsInterfaceProcessor(processor *model.Processor) bool {
	return processor.EntityCategory == consts.ProcessorInterface
}
