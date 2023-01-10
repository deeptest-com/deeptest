package agentExec

import (
	"crypto/tls"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12/websocket"
	"golang.org/x/net/http2"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Session struct {
	ScenarioId uint
	Name       string

	HttpClient  *http.Client
	Http2Client *http.Client
	Failfast    bool

	RootProcessor *Processor
	Report        *Report

	WsMsg *websocket.Message
}

func NewSession(req *ProcessorExecObj, failfast bool, wsMsg *websocket.Message) (ret *Session) {
	root := req.RootProcessor
	variables := req.Variables

	ImportVariables(root.ID, variables, consts.Global)

	session := Session{
		ScenarioId:    root.ScenarioId,
		Name:          root.Name,
		RootProcessor: root,
		Failfast:      failfast,
		WsMsg:         wsMsg,
	}

	jar, _ := cookiejar.New(nil)
	session.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar:     jar, // insert response cookies into request
		Timeout: 120 * time.Second,
	}

	session.Http2Client = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 120 * time.Second,
	}

	ret = &session

	return
}

func (s *Session) Run() {
	s.RootProcessor.Run(s)
}
