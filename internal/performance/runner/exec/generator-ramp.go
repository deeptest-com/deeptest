package runnerExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"sync"
	"time"
)

type RampVuGenerator struct {
}

func (g RampVuGenerator) Run(execCtx context.Context) (err error) {
	execParams := getExecParamsInCtx(execCtx)

	if len(execParams.Stages) != 1 {
		return
	}

	vuNo := 0
	for i := 1; i <= len(execParams.Stages); i++ {
		stage := execParams.Stages[i]

		target := getVuNumbByWeight(int(stage.Target), execParams.Weight)
		stageDuration := int(stage.Duration)
		stageLoop := int(stage.Loop)

		startTime := time.Now().Unix()

		var wgVus sync.WaitGroup

		for index := 1; index <= target; index++ {
			childCtx := execCtx

			childTimeoutCtx, _ := context.WithTimeout(execCtx, time.Duration(stageDuration)*time.Second)

			// generate ExecParams for each stage
			execPramsOfStage := ptdomain.ExecParamsInCtx{}
			copier.CopyWithOption(&execPramsOfStage, execParams, copier.Option{DeepCopy: true})
			execPramsOfStage.Loop = stageLoop

			childCtx = genExecParamsCtx(&execPramsOfStage, childTimeoutCtx)

			wgVus.Add(1)

			result := ptProto.PerformanceExecResp{
				Timestamp: time.Now().UnixMilli(),
				RunnerId:  execParams.RunnerId,
				Room:      execParams.Room,

				VuCount: 1,
			}
			execParams.Sender.Send(result)

			go func() {
				defer wgVus.Done()

				execParams.VuNo = index
				ExecScenarioWithVu(childCtx)
			}()

			vuNo++

			// 尽量平均加载
			leftVus := target - index - 1
			leftTime := getLeftTime(startTime, stageDuration)
			waitTime(int64(leftVus), leftTime)

			select {
			case <-childCtx.Done():
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
	time.Sleep(time.Duration(leftTime/leftVus) * time.Second)

	return
}
