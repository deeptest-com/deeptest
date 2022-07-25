package websocketHelper

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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

func SendOutputMsg(msg string, data interface{}, wsMsg *websocket.Message) {
	logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room,
		strings.ReplaceAll(strings.TrimSpace(msg), `%`, `%%`)))

	msg = strings.Trim(msg, "\n")
	resp := _domain.WsResp{Msg: msg, Data: data}

	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
	PubMsg(mqData)
}

func SendExecMsg(msg string, log domain.Log, wsMsg *websocket.Message) {
	logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room,
		strings.ReplaceAll(strings.TrimSpace(msg), `%`, `%%`)))

	msg = strings.TrimSpace(msg)
	resp := _domain.WsResp{Msg: msg, Data: log}

	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
	PubMsg(mqData)
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
