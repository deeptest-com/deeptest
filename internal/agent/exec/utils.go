package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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

func IsExecutableContainerProcessor(processor *Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), executableContainerProcessors) &&
		processor.EntityType != consts.ProcessorLoopBreak
}

func IsNoExecutableContainerProcessor(processor *Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), noExecutableContainerProcessors)
}

func IsActionProcessor(processor *Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), actionProcessors) ||
		processor.EntityType == consts.ProcessorLoopBreak
}

func IsInterfaceProcessor(processor *Processor) bool {
	return processor.EntityCategory == consts.ProcessorInterface
}
