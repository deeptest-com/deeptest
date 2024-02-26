package handler

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/performance/conductor/exec"
	conductorService "github.com/aaronchen2k/deeptest/internal/performance/conductor/service"
	ptwebsocket "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12/websocket"
	"github.com/sirupsen/logrus"
)

type PerformanceTestWebSocketCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`

	performanceTestService *conductorService.PerformanceTestService
}

func NewPerformanceTestWebSocketCtrl() *PerformanceTestWebSocketCtrl {
	inst := &PerformanceTestWebSocketCtrl{Namespace: consts.WsPerformanceTestNamespace}

	return inst
}

func (c *PerformanceTestWebSocketCtrl) getService() (ret *conductorService.PerformanceTestService) {
	if c.performanceTestService != nil {
		return c.performanceTestService
	}

	insts := &conductorService.PerformanceTestService{}

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: insts},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	c.performanceTestService = insts

	return c.getService()
}

func (c *PerformanceTestWebSocketCtrl) OnNamespaceConnected(wsMsg websocket.Message) error {
	ptwebsocket.SetConn(c.Conn)

	_logUtils.Infof(_i118Utils.Sprintf("connect to namespace %s, id=%s room=%s",
		consts.WsPerformanceTestNamespace, c.Conn.ID(), wsMsg.Room))

	return nil
}

func (c *PerformanceTestWebSocketCtrl) OnNamespaceDisconnect(wsMsg websocket.Message) error {
	_logUtils.Infof(_i118Utils.Sprintf("disconnect to namespace %s, id=%s room=%s",
		consts.WsPerformanceTestNamespace, c.Conn.ID(), wsMsg.Room))

	// stop performance msg and log schedule job
	conductorExec.SuspendWsMsg()
	c.getService().StopSendLog()

	resp := _domain.WsResp{Msg: "from agent: disconnected to websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return nil
}

func (c *PerformanceTestWebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)
	_logUtils.Infof("WebSocket OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	req := agentDomain.WsReq{}
	err = json.Unmarshal(wsMsg.Body, &req)
	if err != nil {
		execUtils.SendErrorMsg(err, consts.Processor, &wsMsg)
		return
	}

	if req.Act == consts.JoinPerformanceTest { // test
		err = c.getService().ExecJoin(req.PerformanceTestExecReq.Room, &wsMsg)

	} else if req.Act == consts.StartPerformanceTest {
		err = c.getService().ExecStart(req.PerformanceTestExecReq, &wsMsg)

	} else if req.Act == consts.StopPerformanceTest {
		err = c.getService().ExecStop(&wsMsg)

	} else if req.Act == consts.StartPerformanceLog { // log
		err = c.getService().StartSendLog(req.PerformanceTestExecReq, &wsMsg)

	} else if req.Act == consts.StopPerformanceLog {
		err = c.getService().StopSendLog()

	}

	return
}
