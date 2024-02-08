package ptdomain

import (
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/kataras/iris/v12"
)

type WsResp struct {
	Uuid            string                                `json:"uuid"`
	Category        ptconsts.MsgCategory                  `json:"category"`
	InstructionType ptconsts.MsgInstructionServerToRunner `json:"instructionType"`
	ResultType      ptconsts.MsgResultTypeToWsClient      `json:"resultType"`

	Msg  string      `json:"msg"`
	Info iris.Map    `json:"info,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type MqMsg struct {
	Namespace string `json:"namespace"`
	Room      string `json:"room"`
	Event     string `json:"event"`
	Content   string `json:"content"`
}
