package main

import (
	"flag"
	agentServe "github.com/deeptest-com/deeptest/cmd/agent/serve"
	v1 "github.com/deeptest-com/deeptest/cmd/agent/v1"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/core/cron"
	websocketHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/websocket"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	_consts "github.com/deeptest-com/deeptest/pkg/consts"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	AppVersion string
	BuildTime  string
	GoVersion  string
	GitHash    string

	flagSet *flag.FlagSet
)

// @title 第三方API文档
// @version 1.0
// @contact.name API Support

func main() {
	flagSet = flag.NewFlagSet("deeptest", flag.ContinueOnError)
	flagSet.IntVar(&consts.Port, "p", 0, "")
	flagSet.BoolVar(&_consts.Verbose, "verbose", false, "")
	flagSet.Parse(os.Args[1:])

	websocketHelper.InitMq()

	agent := agentServe.Init()
	if agent == nil {
		return
	}

	injectModule(agent)

	agent.Start()

	_logUtils.Infof("start agent")
}

func injectModule(ws *agentServe.AgentServer) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	cron := cron.NewServerCron()
	cron.Init()
	indexModule := v1.NewIndexModule()

	// inject objects
	if err := g.Provide(
		&inject.Object{Value: dao.GetDB("")},
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
