package handler

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	conductorExec "github.com/aaronchen2k/deeptest/internal/performance/conductor/exec"
	ptlog "github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	ptwebsocket "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

type PerformanceLogWebSocketCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`
}

func NewPerformanceLogWebSocketCtrl() *PerformanceLogWebSocketCtrl {
	inst := &PerformanceLogWebSocketCtrl{Namespace: consts.WsPerformanceLogNamespace}

	return inst
}

func (c *PerformanceLogWebSocketCtrl) OnNamespaceConnected(wsMsg websocket.Message) error {
	ptwebsocket.SetLogConn(c.Conn)

	_logUtils.Infof(_i118Utils.Sprintf("connect to log namespace %s, id=%s room=%s",
		consts.WsPerformanceLogNamespace, c.Conn.ID(), wsMsg.Room))

	return nil
}

func (c *PerformanceLogWebSocketCtrl) OnNamespaceDisconnect(wsMsg websocket.Message) error {
	_logUtils.Infof(_i118Utils.Sprintf("disconnect to log namespace %s, id=%s room=%s",
		consts.WsPerformanceLogNamespace, c.Conn.ID(), wsMsg.Room))

	// stop performance log msg
	testItem := conductorExec.GetCurrItem()
	if testItem != nil {
		service := conductorExec.GetLogService(testItem.Room)
		if service != nil {
			service.StopSendLog()
		}
	}

	resp := _domain.WsResp{Msg: "disconnected to log websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return nil
}

func (c *PerformanceLogWebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)
	_logUtils.Infof("WebSocket log OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	req := agentDomain.WsReq{}
	err = json.Unmarshal(wsMsg.Body, &req)
	if err != nil {
		execUtils.SendErrorMsg(err, consts.Processor, &wsMsg)
		return
	}

	c.exec(req, wsMsg)

	return
}

func (c *PerformanceLogWebSocketCtrl) exec(req agentDomain.WsReq, wsMsg websocket.Message) (err error) {
	room := req.PerformanceLogExecReq.Room

	if req.Act == consts.StartPerformanceLog { // log
		service := conductorExec.GetLogService(room)
		if service == nil {
			service = conductorExec.CreatePerformanceLogService()
			conductorExec.SetLogService(room, service)
		}

		err = service.StartSendLog(req.PerformanceLogExecReq, &wsMsg)

	} else if req.Act == consts.StopPerformanceLog {
		service := conductorExec.GetLogService(room)
		if service == nil {
			ptlog.Logf("not found test service for room %s to stop log", room)
			return
		}

		err = service.StopSendLog()

	}

	return
}
