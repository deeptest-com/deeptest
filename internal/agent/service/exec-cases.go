package service

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
)

func RunCases(req *agentExec.CasesExecReq, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run cases %s on env %d", req.ExecUUid, req.EnvironmentId)

	return
}
