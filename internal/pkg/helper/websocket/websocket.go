package websocketHelper

import (
	"encoding/json"
	"fmt"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"strings"
)

var (
	wsConn *neffos.Conn
)

func SendExecMsg(msg string, log interface{}, category consts.WsMsgCategory, wsMsg *websocket.Message) {
	msg = strings.TrimSpace(msg)
	resp := _domain.WsResp{Msg: msg, Category: category, Data: log}

	bytes, _ := json.Marshal(resp)

	msg = strings.ReplaceAll(strings.TrimSpace(msg), `%`, `%%`)
	if wsMsg != nil {
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, msg))

		mqData := _domain.MqMsg{
			Namespace: wsMsg.Namespace,
			Room:      wsMsg.Room,
			Event:     wsMsg.Event,
			Content:   string(bytes),
		}

		PubMsg(mqData)
	} else {
		logUtils.Infof(msg)
	}
}

func SendExecStatus(category consts.WsMsgCategory, wsMsg *websocket.Message) {
	resp := _domain.WsResp{Category: category}
	bytes, _ := json.Marshal(resp)

	if wsMsg != nil {
		mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, category))
		PubMsg(mqData)
	} else {
		logUtils.Infof(string(bytes))
	}
}

func SendExecResult(data interface{}, wsMsg *websocket.Message) {
	resp := _domain.WsResp{Category: consts.ProgressResult, Data: data}
	if data != nil {
		resp.Data = data
	}
	bytes, _ := json.Marshal(resp)

	if wsMsg != nil {
		mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, consts.ProgressResult))
		PubMsg(mqData)
	} else {
		logUtils.Infof(string(bytes))
	}
}

func SendStatInfo(data agentDomain.InterfaceStat, wsMsg *websocket.Message) {
	resp := _domain.WsResp{Category: consts.Statistic, Data: data}
	bytes, _ := json.Marshal(resp)

	if wsMsg != nil {
		mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
		PubMsg(mqData)
	} else {
		logUtils.Infof(string(bytes))
	}
}

func SendInitializeMsg(data interface{}, wsMsg *websocket.Message) {
	resp := _domain.WsResp{Category: consts.Initialize, Data: data}
	if data != nil {
		resp.Data = data
	}
	bytes, _ := json.Marshal(resp)

	if wsMsg != nil {
		mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, consts.ProgressResult))
		PubMsg(mqData)
	} else {
		logUtils.Infof(string(bytes))
	}
}

func Broadcast(namespace, room, event string, content string) {
	if wsConn == nil {
		return
	}

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
