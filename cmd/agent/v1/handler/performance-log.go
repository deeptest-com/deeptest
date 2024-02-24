package handler

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
	controllerExec "github.com/aaronchen2k/deeptest/internal/performance/conductor/exec"
	conductorService "github.com/aaronchen2k/deeptest/internal/performance/conductor/service"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
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
	websocketHelper.SetConn(c.Conn)
	_logUtils.Infof(_i118Utils.Sprintf("connect to namespace %s, id=%s room=%s",
		consts.WsPerformanceLogNamespace, c.Conn.ID(), wsMsg.Room))

	return nil
}

func (c *PerformanceLogWebSocketCtrl) OnNamespaceDisconnect(wsMsg websocket.Message) error {
	_logUtils.Infof(_i118Utils.Sprintf("disconnect to namespace %s, id=%s room=%s",
		consts.WsPerformanceLogNamespace, c.Conn.ID(), wsMsg.Room))

	// stop log schedule job
	req := agentDomain.WsReq{
		Act: consts.StopPerformanceLog,
		PerformanceTestExecReq: ptdomain.PerformanceTestReq{
			BaseExecReqOfRunner: ptdomain.BaseExecReqOfRunner{
				Room: controllerExec.GetRunningRoom(),
			},
		},
	}
	service.StartExec(req, &wsMsg)

	resp := _domain.WsResp{Msg: "from agent: disconnected to websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return nil
}

func (c *PerformanceLogWebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)
	_logUtils.Infof("WebSocket OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	req := agentDomain.WsReq{}
	err = json.Unmarshal(wsMsg.Body, &req)
	if err != nil {
		execUtils.SendErrorMsg(err, consts.Processor, &wsMsg)
		return
	}

	if req.Act == "init" {
		return
	}

	err = conductorService.RunPerformanceLog(req.Act, req.PerformanceTestExecReq, &wsMsg)

	return
}
