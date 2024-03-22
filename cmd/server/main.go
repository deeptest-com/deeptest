package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/cmd/server/serve"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
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

	flagSet = flag.NewFlagSet("deeptest", flag.ContinueOnError)
	flagSet.IntVar(&consts.WebPort, "p", 0, "")
	flagSet.BoolVar(&_consts.Verbose, "verbose", false, "")
	flagSet.Parse(os.Args[1:])

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
