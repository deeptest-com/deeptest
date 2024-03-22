package handler

import (
	conductorExec "github.com/aaronchen2k/deeptest/internal/agent/performance/conductor/exec"
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"strings"
)

type PerformanceCtrl struct {
}

func (c *PerformanceCtrl) GetState(ctx iris.Context) {
	runningTests := conductorExec.GetTestItems()

	if runningTests == nil || *runningTests == nil {
		ret := iris.Map{
			"isBusy":   false,
			"grpcPort": getGrpcPort(),
		}

		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
		return
	}

	var conductorItems, runnerItems []*ptdomain.TestItem

	for _, item := range *runningTests {
		if item.Role == ptconsts.Conductor {
			simpleItem := ptdomain.TestItem{}
			copier.CopyWithOption(&simpleItem, item, copier.Option{DeepCopy: true})

			removeConductorRawData(&simpleItem)

			conductorItems = append(conductorItems, &simpleItem)

		} else if item.Role == ptconsts.Runner {
			simpleItem := ptdomain.TestItem{}
			copier.CopyWithOption(&simpleItem, item, copier.Option{DeepCopy: true})

			removeRunnerRawData(&simpleItem)

			runnerItems = append(runnerItems, &simpleItem)
		}
	}

	conductorExec.DestroyPerformanceLogService("")

	service := conductorExec.GetTestService("room")
	if service != nil {
		service.ExecStop(nil)
		conductorExec.DeleteTestService("room")
	}

	ret := iris.Map{
		"conductorTests": conductorItems,
		"runnerTests":    runnerItems,
		"isBusy":         len(conductorItems) > 0 || len(runnerItems) > 0,

		"grpcPort": getGrpcPort(),
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

func (c *PerformanceCtrl) ForceStop(ctx iris.Context) {
	// stop all log services
	conductorExec.DestroyAllPerformanceLogServices()

	// stop conductor
	conductorTask := conductorExec.GetConductorTask()
	if conductorTask != nil {
		service := conductorExec.GetTestService(conductorTask.Room)
		if service != nil {
			service.ExecStop(nil)
			conductorExec.DeleteTestService(conductorTask.Room)
		}
	}

	// stop runner
	runnerTask := conductorExec.GetRunnerTask()
	if runnerTask != nil {
		// send a grpc stop instruction to itself

		client := conductorExec.GetGrpcClient(runnerTask.RunnerReq.ConductorGrpcAddress)

		conductorExec.CallRunnerExecStopByGrpc(client, runnerTask.Room)
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func removeConductorRawData(item *ptdomain.TestItem) {
	if item.ConductorReq == nil {
		return
	}

	item.ConductorReq.Token = ""
}

func removeRunnerRawData(item *ptdomain.TestItem) {
	if item.RunnerReq == nil {
		return
	}

	for index, _ := range item.RunnerReq.Scenarios {
		item.RunnerReq.Scenarios[index].ProcessorRaw = nil
	}

	item.RunnerReq.ExecSceneRaw = nil
}

func getGrpcPort() (ret string) {
	ret = "9528"

	arr1 := strings.Split(config.CONFIG.System.GrpcAddress, ":")
	if len(arr1) > 1 {
		ret = arr1[1]
	}

	return
}
