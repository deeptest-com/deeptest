package main

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/cron"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/core/web"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1"
	websocketHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/websocket"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
)

// @title DeepTest服务端API文档
// @version 1.0
// @contact.name API Support
// @contact.url https://github.com/aaronchen2k/deeptest/issues
// @contact.email 462626@qq.com
func main() {
	cron.NewServerCron().Init()
	websocketHelper.InitMq()

	webServer := web.Init()
	if webServer == nil {
		return
	}
	injectModule(webServer)
	webServer.Run()
}

func injectModule(ws *web.WebServer) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	indexModule := v1.NewIndexModule()

	// inject objects
	if err := g.Provide(
		&inject.Object{Value: dao.GetDB()},
		&inject.Object{Value: indexModule},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	ws.AddModule(indexModule.Party())

	_logUtils.Infof("start server")
}
