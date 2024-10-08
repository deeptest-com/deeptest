package handler

import (
	"encoding/json"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	execUtils "github.com/deeptest-com/deeptest/internal/agent/exec/utils/exec"
	agentService "github.com/deeptest-com/deeptest/internal/agent/service"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	websocketHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/websocket"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	_i118Utils "github.com/deeptest-com/deeptest/pkg/lib/i118"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

var (
	ch chan int
)

type ExecByWebSocketCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`
}

func NewWebsocketCtrl() *ExecByWebSocketCtrl {
	inst := &ExecByWebSocketCtrl{Namespace: consts.WsDefaultNamespace}
	return inst
}

// OnNamespaceConnected
func (c *ExecByWebSocketCtrl) OnNamespaceConnected(wsMsg websocket.Message) error {
	websocketHelper.SetConn(c.Conn)
	_logUtils.Infof(_i118Utils.Sprintf("ws_namespace_connected :id=%v room=%v", c.Conn.ID(), wsMsg.Room))

	resp := _domain.WsResp{Msg: "from agent: connected to websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return nil
}

// OnNamespaceDisconnect
// This will call the "OnVisit" event on all clients, except the current one,
// it can't because it's left but for any case use this type of design.
func (c *ExecByWebSocketCtrl) OnNamespaceDisconnect(wsMsg websocket.Message) error {
	_logUtils.Infof(_i118Utils.Sprintf("ws_namespace_disconnected :id=%v room=%v", c.Conn.ID(), wsMsg.Room))

	resp := _domain.WsResp{Msg: "from agent: disconnected to websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return nil
}

// OnChat This will call the "OnVisit" event on all clients,
// including the current one, with the 'newCount' variable.
func (c *ExecByWebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)
	_logUtils.Infof("WebSocket OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	req := agentExec.WsReq{}
	err = json.Unmarshal(wsMsg.Body, &req)
	if err != nil {
		execUtils.SendErrorMsg(err, consts.Processor, &wsMsg)

		return
	}

	if req.Act == "init" {
		return
	}

	err = agentService.StartExec(req, &wsMsg)

	return
}
