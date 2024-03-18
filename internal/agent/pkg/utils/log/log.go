package agentLog

import (
	"fmt"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func Logf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)

	if ptlog.Logger != nil {
		ptlog.Logf(msg)
	} else {
		_logUtils.Infof(msg)
	}
}

func Errf(str string, args ...interface{}) {
	msg := fmt.Sprintf(str, args...)

	if ptlog.Logger != nil {
		ptlog.Logf("!!! ERROR: " + msg)
	} else {
		_logUtils.Errorf(msg)
	}
}
