package main

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/agent"
	v1 "github.com/aaronchen2k/deeptest/cmd/agent/v1"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

// @title DeepTest代理API文档
// @version 1.0
// @contact.name API Support
// @contact.url https://github.com/aaronchen2k/deeptest/issues
// @contact.email 462626@qq.com
func main() {
	websocketHelper.InitMq()

	server := agent.Init()
	if server == nil {
		return
	}

	injectModule(server)
	server.Start()
}

func injectModule(ws *agent.WebServer) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	cron := cron.NewServerCron()
	cron.Init()
	indexModule := v1.NewIndexModule()

	// inject objects
	if err := g.Provide(
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

	_logUtils.Infof("start agent")
}
