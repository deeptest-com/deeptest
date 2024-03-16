package runnerExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/agent/performance/runner/metrics"
	"sync"
)

func ExecProgram(execCtx context.Context, execCancel context.CancelFunc, req *ptProto.PerformanceExecStartReq, sender metrics.MessageSender) {
	if req.Mode == ptconsts.Parallel.String() {
		var wgScenarios sync.WaitGroup

		for _, scenario := range req.Scenarios {
			wgScenarios.Add(1)

			go func() {
				defer wgScenarios.Done()

				ExecScenario(execCtx, ptconsts.ExecMode(req.Mode), scenario,
					req.Weight, req.EnvironmentId, req.ExecSceneRaw,
					req.Room, req.WebServerUrl, req.WebServerToken, req.ConductorGrpcAddress,
					req.RunnerId, req.RunnerName, sender)

				ptlog.Logf("scenario %s exec completed", scenario.Name)
			}()
		}

		wgScenarios.Wait()

		// parallel exec completed
		execCancel()

	} else {
		for _, scenario := range req.Scenarios {
			ExecScenario(execCtx, ptconsts.ExecMode(req.Mode), scenario,
				req.Weight, req.EnvironmentId, req.ExecSceneRaw,
				req.Room, req.WebServerUrl, req.WebServerToken, req.ConductorGrpcAddress,
				req.RunnerId, req.RunnerName, sender)
		}

		// sequential exec completed
		execCancel()
	}
}
