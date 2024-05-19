package agentExec

import (
	"context"
	"crypto/tls"
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
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

	// exec context and status
	ExecContext context.Context
	ExecCancel  context.CancelFunc

	ExecScene domain.ExecScene
	IsRunning bool
	ForceStop bool

	InterfaceStat agentExecDomain.InterfaceStat

	// communication
	ServerUrl   string
	ServerToken string

	// goja engine
	GojaRuntime      *goja.Runtime
	GojaRequire      *require.RequireModule
	GojaVariables    *[]domain.ExecVariable
	GojaLogs         *[]string
	GojaSetValueFunc func(name string, value interface{}) // call this as call js _setData method

	InterfaceDebug *InterfaceDebugSession
	ScenarioDebug  *ScenarioDebugSession
}

type InterfaceDebugSession struct {
	DebugInterfaceId uint

	AllVariables map[uint][]domain.ExecVariable

	// used to exchange request and response data between goja and go
	CurrRequest  domain.BaseRequest
	CurrResponse domain.DebugResponse
}

type ScenarioDebugSession struct {
	ScenarioId uint

	ScopedVariables map[uint][]domain.ExecVariable
	ScopedCookies   map[uint][]domain.ExecCookie
	ScopeHierarchy  map[uint]*[]uint // processId -> ancestorProcessIds
	DatapoolCursor  map[string]int

	CurrProcessorId uint // for interface, pass an empty param to variable opt methods
	CurrProcessor   *Processor
	RootProcessor   *Processor
	Report          *Report

	WsMsg *websocket.Message

	HttpClient  *http.Client
	Http2Client *http.Client
	Failfast    bool
}

func NewInterfaceExecSession(call agentDomain.InterfaceCall) (session *ExecSession) {
	ctx, cancel := context.WithCancel(context.Background())

	session = &ExecSession{
		ExecContext: ctx,
		ExecCancel:  cancel,
		Name:        call.Data.Name,
		VuNo:        0,
		ExecUuid:    call.ExecUuid,

		EnvironmentId: call.Data.EnvironmentId,
		ProjectId:     call.Data.ProjectId,
		TenantId:      call.TenantId,

		ExecScene: call.ExecScene,

		GojaVariables: &[]domain.ExecVariable{},
		GojaLogs:      &[]string{},

		ServerUrl:   call.ServerUrl,
		ServerToken: call.Token,

		InterfaceDebug: &InterfaceDebugSession{
			AllVariables: map[uint][]domain.ExecVariable{},

			DebugInterfaceId: call.Data.DebugInterfaceId,
			CurrRequest:      domain.BaseRequest{},
			CurrResponse:     domain.DebugResponse{},
		},
	}

	InitGojaRuntimeWithSession(session, session.VuNo, session.TenantId)

	return
}

func NewScenarioExecSession(vuNo int, req *ScenarioExecObj, failfast bool, environmentId uint, wsMsg *websocket.Message) (session *ExecSession) {
	ctx, cancel := context.WithCancel(context.Background())

	session = &ExecSession{
		ExecContext: ctx,
		ExecCancel:  cancel,

		Name:     req.Name,
		VuNo:     vuNo,
		ExecUuid: req.ExecUuid,

		EnvironmentId: environmentId,
		ProjectId:     req.RootProcessor.ProjectId,
		TenantId:      req.TenantId,

		ServerUrl:   req.ServerUrl,
		ServerToken: req.Token,

		GojaSetValueFunc: func(name string, value interface{}) {},
		GojaVariables:    &[]domain.ExecVariable{},
		GojaLogs:         &[]string{},

		ScenarioDebug: &ScenarioDebugSession{
			ScenarioId: req.ScenarioId,

			RootProcessor: req.RootProcessor,

			Failfast:       failfast,
			DatapoolCursor: map[string]int{},

			ScopedVariables: map[uint][]domain.ExecVariable{},
			ScopedCookies:   map[uint][]domain.ExecCookie{},
			ScopeHierarchy:  map[uint]*[]uint{},

			WsMsg: wsMsg,
		},
	}

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
