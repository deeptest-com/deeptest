package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/cmd/agent/serve"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/performance"
	ptwebsocket "github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	commService "github.com/aaronchen2k/deeptest/internal/pkg/service"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/fatih/color"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	AppVersion string
	BuildTime  string
	GoVersion  string
	GitHash    string

	flagSet *flag.FlagSet
)

// @title 乐研API文档
// @version 1.0
// @contact.name API Support

func main() {
	flagSet = flag.NewFlagSet("deeptest", flag.ContinueOnError)
	flagSet.IntVar(&consts.Port, "p", 0, "")
	flagSet.BoolVar(&_consts.Verbose, "verbose", false, "")
	flagSet.Parse(os.Args[1:])

	// grpc service
	go performance.StartGrpcServe()

	// init queues
	websocketHelper.InitMq()
	ptwebsocket.InitTestMq()
	ptwebsocket.InitLogMq()

	agent := agentServe.Init()
	if agent == nil {
		return
	}

	injectWebSocketModule((*agent).GetApp())

	injectWebModule(agent)

	agent.Start()

	_logUtils.Infof("start agent")
}

func injectWebSocketModule(app *iris.Application) {
	// init websocket
	websocketCtrl := handler.NewWebsocketCtrl()
	websocketTestCtrl := handler.NewPerformanceTestWebSocketCtrl()
	websocketLogCtrl := handler.NewPerformanceLogWebSocketCtrl()

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	err := g.Provide(
		&inject.Object{Value: websocketCtrl},
		&inject.Object{Value: websocketTestCtrl},
		&inject.Object{Value: websocketLogCtrl},
	)
	if err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err = g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	websocketAPI := app.Party(consts.WsPath)
	mvcApp := mvc.New(websocketAPI)
	mvcApp.Register(&commService.PrefixedLogger{Prefix: ""})
	mvcApp.HandleWebsocket(websocketCtrl)
	mvcApp.HandleWebsocket(websocketTestCtrl)
	mvcApp.HandleWebsocket(websocketLogCtrl)

	websocketServer := websocket.New(gorilla.Upgrader(
		gorillaWs.Upgrader{
			CheckOrigin: func(*http.Request) bool { return true },
		}), mvcApp)

	websocketAPI.Get("/", websocket.Handler(websocketServer))
}

func injectWebModule(ws *agentServe.AgentServer) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	cron := cron.NewServerCron()
	cron.Init()
	indexModule := v1.NewIndexModule()

	// inject objects
	if err := g.Provide(
		&inject.Object{Value: dao.GetDB()},
		&inject.Object{Value: cron},
		&inject.Object{Value: indexModule},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	ws.AddModule(indexModule.Party())
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
