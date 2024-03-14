package runnerExec

import (
	"context"
	"encoding/json"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

func ExecRendezvousProcessor(processor *agentExec.Processor, timeoutCtx context.Context, room string, runnerId int32, vuNo int, serverAddress string, index int) {
	name := processor.Name

	entity := agentExec.ProcessorPerformanceRendezvous{}
	json.Unmarshal(processor.EntityRaw, &entity)

	newArrivedVal := performanceUtils.IncreaseRendezvousArrived(room, name, serverAddress)
	ptlog.Logf("====== VU %d: rendezvous Arrived Value added, value is %d", vuNo, newArrivedVal)

	// wait condition util rendezvous ready
	var value int
	var ready bool
	var newPassedVal int

	for !ready {
		value, ready = performanceUtils.IsRendezvousReady(room, name, serverAddress, entity.Target)

		ptlog.Logf("------ VU %d: rendezvous wait, Arrived Value is %d", vuNo, value)
		time.Sleep(1 * time.Second)

		select {
		case <-timeoutCtx.Done():
			_logUtils.Debugf("****** VU %d: rendezvous processor exit scenario processors by signal", vuNo)
			goto Label_END_RENDEZVOUS_PROCESSOR

		default:
		}
	}

	ptlog.Logf("------ VU %d: rendezvous passed, Arrived Value is %d", vuNo, value)

	// reset if all vus passed
	newPassedVal = performanceUtils.IncreaseRendezvousPassed(room, name, serverAddress)
	if newPassedVal >= entity.Target {
		ptlog.Logf("------ VU %d: before rendezvous reset, Arrived Value is %d", vuNo, value)

		value = performanceUtils.ResetRendezvous(room, name, serverAddress)

		ptlog.Logf("====== VU %d: after rendezvous reset, Arrived Value is %d", vuNo, value)

		ptlog.Logf("\n")
	}

Label_END_RENDEZVOUS_PROCESSOR:

	return
}
