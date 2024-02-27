package handler

import (
	conductorExec "github.com/aaronchen2k/deeptest/internal/performance/conductor/exec"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type PerformanceStateCtrl struct {
}

func (c *PerformanceStateCtrl) Get(ctx iris.Context) {
	runningTests := conductorExec.GetTestItems()

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: runningTests})
}
