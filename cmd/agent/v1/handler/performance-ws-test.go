package handler

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	conductorExec "github.com/aaronchen2k/deeptest/internal/agent/performance/conductor/exec"
	ptconsts "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/consts"
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	ptwebsocket "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

type PerformanceTestWebSocketCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`
}

func NewPerformanceTestWebSocketCtrl() *PerformanceTestWebSocketCtrl {
	inst := &PerformanceTestWebSocketCtrl{Namespace: consts.WsPerformanceTestNamespace}

	return inst
}

func (c *PerformanceTestWebSocketCtrl) OnNamespaceConnected(wsMsg websocket.Message) error {
	ptwebsocket.SetTestConn(c.Conn)

	_logUtils.Infof(_i118Utils.Sprintf("connect to test namespace %s, id=%s room=%s",
		consts.WsPerformanceTestNamespace, c.Conn.ID(), wsMsg.Room))

	return nil
}

func (c *PerformanceTestWebSocketCtrl) OnNamespaceDisconnect(wsMsg websocket.Message) error {
	_logUtils.Infof(_i118Utils.Sprintf("disconnect to test namespace %s, id=%s room=%s",
		consts.WsPerformanceTestNamespace, c.Conn.ID(), wsMsg.Room))

	// stop performance result msg
	conductorExec.SuspendWsMsg()

	resp := _domain.WsResp{Msg: "from agent: disconnected to test websocket"}
	bytes, _ := json.Marshal(resp)
	mqData := _domain.MqMsg{Namespace: wsMsg.Namespace, Room: wsMsg.Room, Event: wsMsg.Event, Content: string(bytes)}

	websocketHelper.PubMsg(mqData)

	return nil
}

func (c *PerformanceTestWebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)
	_logUtils.Infof("WebSocket test OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	req := agentDomain.WsReq{}
	err = json.Unmarshal(wsMsg.Body, &req)
	if err != nil {
		execUtils.SendErrorMsg(err, consts.Processor, &wsMsg)
		return
	}

	c.exec(req, wsMsg)

	return
}

func (c *PerformanceTestWebSocketCtrl) exec(req agentDomain.WsReq, wsMsg websocket.Message) (err error) {
	room := req.PerformanceTestExecReq.Room

	if req.Act == consts.StartPerformanceTest {
		service := conductorExec.CreatePerformanceTestService()
		conductorExec.SetTestService(room, service)

		err = service.ExecStart(req.PerformanceTestExecReq, &wsMsg)

	} else if req.Act == consts.StopPerformanceTest {
		service := conductorExec.GetTestService(room)
		if service == nil {
			ptlog.Logf("not found test service for room %s to stop", room)
			ptwebsocket.SendExecInstructionToClient("", "", ptconsts.MsgInstructionEnd, &wsMsg)
			return
		}

		err = service.ExecStop(&wsMsg)

		conductorExec.DeleteTestService(room)

	} else if req.Act == consts.JoinPerformanceTest {
		if room == "" { // to join exist room, may be reload page with no room field
			conductorTask := conductorExec.GetConductorTask()
			if conductorTask == nil {
				ptwebsocket.SendExecInstructionToClient("", nil, ptconsts.MsgInstructionJoinExist, &wsMsg)
				return
			}

			service := conductorExec.GetTestService(conductorTask.Room)
			if service == nil {
				ptwebsocket.SendExecInstructionToClient("", nil, ptconsts.MsgInstructionJoinExist, &wsMsg)
				return
			}

			ptwebsocket.SendExecInstructionToClient(conductorTask.Room, conductorTask.ConductorReq, ptconsts.MsgInstructionJoinExist, &wsMsg)
			conductorExec.ResumeWsMsg()

		} else { //  join exist room successfully, do nothing except sending a start msg
			service := conductorExec.GetTestService(room)
			if service != nil {
				conductorTask := conductorExec.GetConductorTask()
				ptwebsocket.SendExecInstructionToClient("performance test joined", conductorTask, ptconsts.MsgInstructionStart, &wsMsg)
			}
		}

	}

	return
}
