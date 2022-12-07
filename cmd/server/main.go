package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/cmd/server/serverServe"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
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

	serverServe.Start()
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
