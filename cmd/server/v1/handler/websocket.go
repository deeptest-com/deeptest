package handler

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/websocket"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	_i118Utils "github.com/deeptest-com/deeptest/pkg/lib/i118"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

const (
	result = "result"
)

var (
	ch chan int
)

type WebSocketCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`

	ScenarioExecService *service.ScenarioExecService `inject:""`
}

func NewWebsocketCtrl() *WebSocketCtrl {
	inst := &WebSocketCtrl{Namespace: "TestNamespace"}
	return inst
}

func (c *WebSocketCtrl) OnNamespaceConnected(wsMsg websocket.Message) error {
	websocketHelper.SetConn(c.Conn)

	_logUtils.Infof(_i118Utils.Sprintf("ws_namespace_connected", c.Conn.ID(), wsMsg.Room))

	resp := _domain.WsResp{Msg: "from server: connected to websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
	websocketHelper.PubMsg(mqData)

	return nil
}

// OnNamespaceDisconnect
// This will call the "OnVisit" event on all clients, except the current one,
// it can't because it's left but for any case use this type of design.
func (c *WebSocketCtrl) OnNamespaceDisconnect(wsMsg websocket.Message) error {
	_logUtils.Infof(_i118Utils.Sprintf("ws_namespace_disconnected", c.Conn.ID()))

	resp := _domain.WsResp{Msg: "from server: disconnected to websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}
	websocketHelper.PubMsg(mqData)

	return nil
}

// OnChat This will call the "OnVisit" event on all clients, including the current one, with the 'newCount' variable.
func (c *WebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)

	_logUtils.Infof("WebSocket OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	resp := _domain.WsResp{Msg: fmt.Sprintf("from server: OnChat msg is \"%s\"", string(wsMsg.Body))}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return
}
