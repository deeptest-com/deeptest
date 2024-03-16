package ptwebsocket

import (
	"encoding/json"
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"strings"
)

var (
	wsConnTest *neffos.Conn
)

func SendExecInstructionToClient(msg string, data interface{}, instructionType ptconsts.MsgInstructionServerToRunner, wsMsg *websocket.Message) {
	obj := ptdomain.WsResp{
		Category:        ptconsts.MsgCategoryInstruction,
		InstructionType: instructionType,
		Msg:             strings.TrimSpace(msg),
		Data:            data,
	}
	bytes, _ := json.Marshal(obj)

	msg = strings.ReplaceAll(strings.TrimSpace(msg), `%`, `%%`)
	if wsMsg != nil {
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, msg))

		BroadcastTest(wsMsg.Namespace, wsMsg.Room, wsMsg.Event, string(bytes))

	} else {
		logUtils.Infof(msg)
	}
}

func SendExecResultToClient(data interface{}, resultType ptconsts.MsgResultTypeToWsClient, execUUid string, wsMsg *websocket.Message) {
	resp := ptdomain.WsResp{
		Uuid:       execUUid,
		Category:   ptconsts.MsgCategoryResult,
		ResultType: resultType,
		Data:       data,
	}
	if data != nil {
		resp.Data = data
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		ptlog.Logf("SendExecResultToClient err: %s", err.Error())
	}

	if wsMsg != nil {
		mqData := _domain.MqMsg{
			Namespace: wsMsg.Namespace,
			Room:      wsMsg.Room,
			Event:     wsMsg.Event,
			Content:   string(bytes),
		}
		ptlog.Logf(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, ptconsts.MsgCategoryResult))

		PubTestMsg(mqData)

	} else {
		logUtils.Infof(string(bytes))
	}
}

func BroadcastTest(namespace, room, event string, content string) {
	if wsConnTest == nil {
		return
	}

	wsConnTest.Server().Broadcast(nil, websocket.Message{
		Namespace: namespace,
		Room:      room,
		Event:     event,
		Body:      []byte(content),
	})
}

func SetTestConn(conn *neffos.Conn) {
	wsConnTest = conn
}
