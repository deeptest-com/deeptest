package main

import (
	"flag"
	"github.com/deeptest-com/deeptest/cmd/server/serve"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/websocket"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	AppVersion string
	BuildTime  string
	GoVersion  string
	GitHash    string
	flagSet    *flag.FlagSet
)

// @title DeepTest服务端API文档
// @version 3.0
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

	serve.Start()

	_logUtils.Infof("start server")
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
