package agentUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/string"
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

func IsExecutableContainerProcessor(entityType consts.ProcessorType) bool {
	return _stringUtils.FindInArr(entityType.ToString(), executableContainerProcessors) &&
		entityType != consts.ProcessorLoopBreak
}

func IsNoExecutableContainerProcessor(entityType consts.ProcessorType) bool {
	return _stringUtils.FindInArr(entityType.ToString(), noExecutableContainerProcessors)
}

func IsActionProcessor(entityType consts.ProcessorType) bool {
	return _stringUtils.FindInArr(entityType.ToString(), actionProcessors) ||
		entityType == consts.ProcessorLoopBreak
}

func IsInterfaceProcessor(category consts.ProcessorCategory) bool {
	return category == consts.ProcessorInterface
}
