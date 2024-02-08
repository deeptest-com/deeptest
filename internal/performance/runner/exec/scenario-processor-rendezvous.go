package exec

import (
	"context"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

func ExecRendezvousProcessor(timeoutCtx context.Context, processor *ptProto.Processor, vuNo, index int, room, serverAddress string) {
	name := processor.Name

	newArrivedVal := IncreaseRendezvousArrived(room, name, serverAddress)
	ptlog.Logf("====== VU %d: rendezvous Arrived Value added, value is %d", vuNo, newArrivedVal)

	// wait condition util rendezvous ready
	var value int
	var ready bool
	var newPassedVal int

	for !ready {
		value, ready = IsRendezvousReady(room, name, serverAddress, int(processor.RendezvousTarget))

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
	newPassedVal = IncreaseRendezvousPassed(room, name, serverAddress)
	if newPassedVal >= int(processor.RendezvousTarget) {
		ptlog.Logf("------ VU %d: before rendezvous reset, Arrived Value is %d", vuNo, value)

		value = ResetRendezvous(room, name, serverAddress)

		ptlog.Logf("====== VU %d: after rendezvous reset, Arrived Value is %d", vuNo, value)

		ptlog.Logf("\n")
	}

Label_END_RENDEZVOUS_PROCESSOR:

	return
}
