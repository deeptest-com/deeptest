package ptwebsocket

import (
	"encoding/json"
	"fmt"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"strings"
)

var (
	wsConn *neffos.Conn
)

func SendExecInstructionToClient(msg string, data interface{}, instructionType ptconsts.MsgInstructionServerToRunner, uuid string, wsMsg *websocket.Message) {
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

		mqData := ptdomain.MqMsg{
			Namespace: wsMsg.Namespace,
			Room:      wsMsg.Room,
			Event:     wsMsg.Event,
			Content:   string(bytes),
		}

		PubWsMsg(mqData)
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
	bytes, _ := json.Marshal(resp)

	if wsMsg != nil {
		mqData := ptdomain.MqMsg{
			Namespace: wsMsg.Namespace,
			Room:      wsMsg.Room,
			Event:     wsMsg.Event,
			Content:   string(bytes),
		}
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, ptconsts.MsgCategoryResult))

		PubWsMsg(mqData)

	} else {
		logUtils.Infof(string(bytes))
	}
}

func Broadcast(namespace, room, event string, content string) {
	wsConn.Server().Broadcast(nil, websocket.Message{
		Namespace: namespace,
		Room:      room,
		Event:     event,
		Body:      []byte(content),
	})
}

func SetConn(conn *neffos.Conn) {
	wsConn = conn
}

type PrefixedLogger struct {
	Prefix string
}

func (s *PrefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.Prefix, msg)
}
