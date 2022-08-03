package execHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/kataras/iris/v12/websocket"
)

//func ExecThreadGroup(processor *model.ProcessorThreadGroup, log *domain.Log, msg websocket.Message) (
//	result string, err error) {
//
//	return
//}

func ExecLogic(processor *model.ProcessorLogic, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecLoop(loop *model.ProcessorLoop, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	if loop.ID == 0 {
		output.Text = "执行前请先配置处理器。"
		return
	}

	typ := loop.ProcessorType
	if typ == consts.ProcessorLoopTime {
		output.Times = loop.Times
		output.Text = fmt.Sprintf("迭代执行%d次。", output.Times)
		return
	}

	return
}

func ExecData(processor *model.ProcessorData, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecTimer(processor *model.ProcessorTimer, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecVariable(processor *model.ProcessorVariable, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecAssertion(processor *model.ProcessorAssertion, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecExtractor(processor *model.ProcessorExtractor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecCookie(processor *model.ProcessorCookie, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}
