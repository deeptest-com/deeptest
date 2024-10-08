package agentServe

import (
	stdContext "context"
	"fmt"
	"github.com/deeptest-com/deeptest/cmd/agent/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/core/middleware"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	zapLog "github.com/deeptest-com/deeptest/internal/pkg/log"
	commService "github.com/deeptest-com/deeptest/internal/pkg/service"
	fileUtils "github.com/deeptest-com/deeptest/pkg/lib/file"
	_i118Utils "github.com/deeptest-com/deeptest/pkg/lib/i118"
	"github.com/facebookgo/inject"
	"github.com/go-playground/validator/v10"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/tests"

	"github.com/kataras/iris/v12/mvc"
)

var client *tests.Client

// Init 初始化web服务
func Init() *AgentServer {
	consts.RunFrom = consts.FromAgent
	//consts.WorkDir = commUtils.GetWorkDir()

	fileUtils.RmDir(consts.TmpDirRelativeAgent)
	config.Init()
	zapLog.Init()
	_i118Utils.Init(consts.Language, "")

	app := iris.New()
	app.Validator = validator.New() //参数验证
	app.Logger().SetLevel(config.CONFIG.System.Level)
	idleConnClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() { //优雅退出
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx) // close all hosts
		close(idleConnClosed)
	})

	// init grpc
	mvc.New(app)

	// init websocket
	websocketCtrl := handler.NewWebsocketCtrl()
	injectWebsocketModule(websocketCtrl)

	websocketAPI := app.Party(consts.WsPath)
	m := mvc.New(websocketAPI)
	m.Register(
		&commService.PrefixedLogger{Prefix: ""},
	)
	m.HandleWebsocket(websocketCtrl)

	websocketServer := websocket.New(gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), m)
	websocketAPI.Get("/", websocket.Handler(websocketServer))

	return &AgentServer{
		app:               app,
		addr:              config.CONFIG.System.AgentAddress,
		timeFormat:        config.CONFIG.System.TimeFormat,
		idleConnClosed:    idleConnClosed,
		globalMiddlewares: []context.Handler{middleware.Error()},
	}
}

// AddModule 添加模块
func (s *AgentServer) AddModule(module ...module.WebModule) {
	s.modules = append(s.modules, module...)
}

// GetModules 获取模块
func (s *AgentServer) GetModules() []module.WebModule {
	return s.modules
}

// GetTestAuth 获取测试验证客户端
func (s *AgentServer) GetTestAuth(t *testing.T) *tests.Client {
	var once sync.Once
	once.Do(
		func() {
			client = tests.New(str.Join("http://", s.addr), t, s.app)
			if client == nil {
				t.Fatalf("client is nil")
			}
		},
	)

	return client
}

// Init 启动web服务
func (s *AgentServer) Start() {
	s.app.UseGlobal(s.globalMiddlewares...)
	err := s.InitRouter()
	if err != nil {
		fmt.Printf("初始化路由错误： %v\n", err)
		panic(err)
	}
	// 添加上传文件路径
	s.app.Listen(
		s.addr,
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(s.timeFormat),
	)
	<-s.idleConnClosed
}

func injectWebsocketModule(websocketCtrl *handler.ExecByWebSocketCtrl) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: websocketCtrl},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
