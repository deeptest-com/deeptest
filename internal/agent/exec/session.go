package agentExec

import (
	"crypto/tls"
	agentExecDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/kataras/iris/v12/websocket"
	"golang.org/x/net/http2"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type ExecSession struct {
	Name          string
	VuNo          int
	ExecUuid      string
	EnvironmentId uint
	ProjectId     uint
	TenantId      consts.TenantId

	// exec status
	ExecScene  domain.ExecScene
	_isRunning bool
	_forceStop bool

	InterfaceStat agentExecDomain.InterfaceStat

	// communication
	ServerUrl   string
	ServerToken string

	// goja engine
	GojaRuntime      *goja.Runtime
	GojaRequire      *require.RequireModule
	GojaSetValueFunc func(name string, value interface{}) // call this as call js _setData method in goja

	_gojaLogs      *[]string
	_gojaVariables *[]domain.ExecVariable

	InterfaceDebug *InterfaceDebugSession
	ScenarioDebug  *ScenarioDebugSession
}

type InterfaceDebugSession struct {
	_debugInterfaceId uint

	AllVariables map[uint][]domain.ExecVariable

	// used to exchange request and response data between goja and go
	_currRequest  domain.BaseRequest
	_currResponse domain.DebugResponse
}

type ScenarioDebugSession struct {
	ScenarioId uint

	ScopedVariables map[uint][]domain.ExecVariable
	ScopedCookies   map[uint][]domain.ExecCookie
	ScopeHierarchy  map[uint]*[]uint // processId -> ancestorProcessIds
	DatapoolCursor  map[string]int

	_currProcessorId uint // for interface, pass an empty param to variable opt methods
	_currProcessor   *Processor
	RootProcessor    *Processor
	Report           *Report

	WsMsg *websocket.Message

	HttpClient  *http.Client
	Http2Client *http.Client
}

func NewInterfaceExecSession(call domain.InterfaceCall) (session *ExecSession) {
	session = &ExecSession{
		Name:     call.Data.Name,
		VuNo:     0,
		ExecUuid: call.ExecUuid,

		EnvironmentId: call.Data.EnvironmentId,
		ProjectId:     call.Data.ProjectId,
		TenantId:      call.TenantId,

		ExecScene:     call.ExecScene,
		InterfaceStat: agentExecDomain.InterfaceStat{},

		ServerUrl:   call.ServerUrl,
		ServerToken: call.Token,

		InterfaceDebug: &InterfaceDebugSession{
			AllVariables: map[uint][]domain.ExecVariable{},

			_debugInterfaceId: call.Data.DebugInterfaceId,
			_currRequest:      domain.BaseRequest{},
			_currResponse:     domain.DebugResponse{},
		},
		ScenarioDebug: &ScenarioDebugSession{ // just put an empty
			ScopedVariables: map[uint][]domain.ExecVariable{},
			DatapoolCursor:  map[string]int{},
		},
	}
	session.ResetGojaVariables()
	session.ResetGojaLogs()

	InitGojaRuntimeWithSession(session, session.VuNo, session.TenantId)

	return
}

func NewScenarioExecSession(vuNo int, req *ScenarioExecObj, environmentId uint, wsMsg *websocket.Message) (session *ExecSession) {
	session = &ExecSession{
		Name:     req.Name,
		VuNo:     vuNo,
		ExecUuid: req.ExecUuid,

		EnvironmentId: environmentId,
		ProjectId:     req.RootProcessor.ProjectId,
		TenantId:      req.TenantId,

		ExecScene:     req.ExecScene,
		InterfaceStat: agentExecDomain.InterfaceStat{},

		ServerUrl:   req.ServerUrl,
		ServerToken: req.Token,

		GojaSetValueFunc: func(name string, value interface{}) {},

		ScenarioDebug: &ScenarioDebugSession{
			ScenarioId: req.ScenarioId,

			RootProcessor: req.RootProcessor,

			DatapoolCursor: map[string]int{},

			ScopedVariables: map[uint][]domain.ExecVariable{},
			ScopedCookies:   map[uint][]domain.ExecCookie{},
			ScopeHierarchy:  map[uint]*[]uint{},

			WsMsg: wsMsg,
		},
		InterfaceDebug: &InterfaceDebugSession{}, // just put an empty
	}
	session.ResetGojaVariables()
	session.ResetGojaLogs()

	InitGojaRuntimeWithSession(session, session.VuNo, session.TenantId)
	ComputerScopeHierarchy(req.RootProcessor, &session.ScenarioDebug.ScopeHierarchy)

	jar, _ := cookiejar.New(nil)
	session.ScenarioDebug.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar:     jar, // insert response cookies into request
		Timeout: 120 * time.Second,
	}

	session.ScenarioDebug.Http2Client = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 120 * time.Second,
	}

	return
}

func (s *ExecSession) Run() {
	s.ScenarioDebug.RootProcessor.Run(s)
}
