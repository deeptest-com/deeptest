package handler

import (
	"encoding/json"
	"fmt"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	execDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"runtime/debug"
)

var (
	ch chan int
)

type ExecByWebSocketCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`
}

func NewWebsocketCtrl() *ExecByWebSocketCtrl {
	inst := &ExecByWebSocketCtrl{Namespace: consts.WsDefaultNameSpace}
	return inst
}

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

// OnChat This will call the "OnVisit" event on all clients, including the current one, with the 'newCount' variable.
func (c *ExecByWebSocketCtrl) OnChat(wsMsg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)
	_logUtils.Infof("WebSocket OnChat: remote address=%s, room=%s, msg=%s", ctx.RemoteAddr(), wsMsg.Room, string(wsMsg.Body))

	req := agentDomain.WsReq{}
	err = json.Unmarshal(wsMsg.Body, &req)
	if err != nil {
		if req.ScenarioExecReq.ScenarioId > 0 {
			sendScenarioErr(err, &wsMsg)
		}
		return
	}

	act := req.Act

	// stop exec
	if act == consts.ExecStop {
		if ch != nil {
			if !execUtils.GetRunning() {
				ch = nil
			} else {
				ch <- 1
				ch = nil
			}
		}

		agentExec.ForceStopExec = true

		if req.ScenarioExecReq.ScenarioId > 0 {
			service.CancelAndSendMsg(req.ScenarioExecReq.ScenarioId, wsMsg)
		} else if req.CasesExecReq.ExecUUid != "" {
			//TODO:
		}

		return
	}

	// already running
	if execUtils.GetRunning() && (act == consts.ExecStart) {
		if req.ScenarioExecReq.ScenarioId > 0 {
			execUtils.SendAlreadyRunningMsg(req.ScenarioExecReq.ScenarioId, consts.Processor, wsMsg)
		}
		return
	}

	// exec task
	go func() {
		defer func(wsMsg websocket.Message) {
			if wsMsgErr := recover(); wsMsgErr != nil {
				s := string(debug.Stack())
				fmt.Printf("err=%v, stack=%s\n", wsMsgErr, s)

				if req.ScenarioExecReq.ScenarioId > 0 {
					sendScenarioErr(fmt.Errorf("%+v", wsMsgErr), &wsMsg)
				}
			}
		}(wsMsg)

		ch = make(chan int, 1)

		if act == consts.ExecScenario {
			service.RunScenario(&req.ScenarioExecReq, &wsMsg)

		} else if act == consts.ExecPlan {
			service.RunPlan(&req.PlanExecReq, &wsMsg)

		} else if act == consts.ExecCase {
			service.RunCases(&req.CasesExecReq, &wsMsg)

		} else if act == consts.ExecMessage {
			service.RunMessage(&req.MessageReq, &wsMsg)
		}
	}()

	return
}

func sendScenarioErr(err error, wsMsg *websocket.Message) {
	root := execDomain.ScenarioExecResult{
		ID:      -1,
		Name:    "执行失败",
		Summary: fmt.Sprintf("错误：%s", err.Error()),
	}
	execUtils.SendExecMsg(root, consts.Processor, wsMsg)

	result := execDomain.ScenarioExecResult{
		ID:       -2,
		ParentId: -1,
		Name:     "执行失败",
		Summary:  fmt.Sprintf("错误：%s", err.Error()),
	}
	execUtils.SendExecMsg(result, consts.Processor, wsMsg)
}
