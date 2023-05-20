package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/cmd/command/action"
	commandConfig "github.com/aaronchen2k/deeptest/internal/command/config"
	zapLog "github.com/aaronchen2k/deeptest/internal/pkg/log"
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
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

	scenarioId int
	planId     int
	envId      int
	server     string
	token      string

	language string

	flagSet *flag.FlagSet
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("deeptest", flag.ContinueOnError)

	flagSet.IntVar(&scenarioId, "s", 0, "")
	flagSet.IntVar(&scenarioId, "scenario", 0, "")

	flagSet.IntVar(&planId, "p", 0, "")
	flagSet.IntVar(&planId, "plan", 0, "")

	flagSet.IntVar(&envId, "e", 0, "")
	flagSet.IntVar(&envId, "env", 0, "")

	flagSet.StringVar(&server, "S", "", "")
	flagSet.StringVar(&server, "server", "", "")

	flagSet.StringVar(&token, "t", "", "")
	flagSet.StringVar(&token, "token", "", "")

	flagSet.BoolVar(&_consts.Verbose, "verbose", false, "")

	flagSet.Parse(os.Args[1:])

	switch os.Args[1] {
	case "help", "-h", "-help", "--help":
		action.PrintUsage(language)

	default:
		if (scenarioId == 0 && planId == 0) || envId == 0 || server == "" || token == "" {
			action.PrintUsage(language)
			return
		} else {
			run(scenarioId, planId, envId, server, token)
		}
	}
}

func run(scenarioId, planId, envId int, server, token string) {
	action.Run(scenarioId, planId, envId, server, token)
}

func init() {
	cleanup()
	commandConfig.InitConfig()
	zapLog.Init("command")
}

func cleanup() {
	color.Unset()
}
