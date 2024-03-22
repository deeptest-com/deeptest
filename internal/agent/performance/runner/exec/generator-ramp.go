package runnerExec

import (
	"context"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	ptProto "github.com/aaronchen2k/deeptest/internal/agent/performance/proto"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"sync"
	"time"
)

type RampVuGenerator struct {
}

func (g RampVuGenerator) Run(execCtx context.Context) (err error) {
	execParams := performanceUtils.GetExecParamsInCtx(execCtx)

	if len(execParams.Stages) == 0 {
		return
	}

	vuNo := 0
	for i := 1; i < len(execParams.Stages); i++ {
		stage := execParams.Stages[i]

		target := performanceUtils.GetVuNumbByWeight(int(stage.Target), execParams.Weight)

		startTime := time.Now().Unix()

		var wgVus sync.WaitGroup

		for index := 1; index <= target; index++ {
			vuCtx, _ := context.WithCancel(execCtx)
			if execParams.GoalDuration > 0 { // control exec time
				vuCtx, _ = context.WithTimeout(execCtx, time.Duration(execParams.GoalDuration)*time.Second)
			}

			wgVus.Add(1)

			result := ptProto.PerformanceExecResp{
				Timestamp:  time.Now().UnixMilli(),
				RunnerId:   execParams.RunnerId,
				RunnerName: execParams.RunnerName,
				Room:       execParams.Room,

				VuCount: 1,
			}
			execParams.Sender.Send(result)

			go func() {
				defer wgVus.Done()

				//execParams.VuNo = index
				ExecScenarioWithVu(vuCtx, index)
			}()

			vuNo++

			// 尽量平均加载
			leftVus := target - index - 1
			leftTime := getLeftTime(startTime, int(stage.Duration))

			if leftTime > 0 {

			}
			waitTime(int64(leftVus), leftTime)

			select {
			case <-vuCtx.Done():
				_logUtils.Debug("<<<<<<< stop stage targets")
				goto Label_END_STAGES

			default:
			}
		}

		// wait all vus completed
		wgVus.Wait()

		select {
		case <-execCtx.Done():
			_logUtils.Debug("<<<<<<< stop stages")
			goto Label_END_STAGES

		default:
		}
	}

Label_END_STAGES:

	return
}

func getLeftTime(startTime int64, dur int) (leftTime int64) {
	currTime := time.Now().Unix()
	leftTime = int64(dur) - (currTime - startTime)

	if leftTime < 0 {
		leftTime = 0
	}

	return
}

func waitTime(leftVus, leftTime int64) (err error) {
	if leftTime > 0 && leftVus > 0 {
		time.Sleep(time.Duration(leftTime/leftVus) * time.Second)
	}

	return
}
