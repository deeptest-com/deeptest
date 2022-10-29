package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/cmd/server/server"
	"github.com/aaronchen2k/deeptest/cmd/server/v1"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	AppVersion string
	BuildTime  string
	GoVersion  string
	GitHash    string

	flagSet *flag.FlagSet
)

// @title DeepTest服务端API文档
// @version 1.0
// @contact.name API Support
// @contact.url https://github.com/aaronchen2k/deeptest/issues
// @contact.email 462626@qq.com
func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	websocketHelper.InitMq()

	server := server.Init()
	if server == nil {
		return
	}

	injectModule(server)
	server.Start()
}

func injectModule(ws *server.WebServer) {
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

	_logUtils.Infof("start server")
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
