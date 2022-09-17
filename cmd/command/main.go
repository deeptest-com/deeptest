package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/internal/command/action"
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

	scenario string
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

	flagSet = flag.NewFlagSet("ztf", flag.ContinueOnError)

	flagSet.StringVar(&scenario, "s", "", "")
	flagSet.StringVar(&scenario, "scenario", "", "")

	flagSet.BoolVar(&_consts.Verbose, "verbose", false, "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "run", ".")
	}

	switch os.Args[1] {
	case "help", "-h", "-help", "--help":
		action.PrintUsage(language)

	default: // run
		if scenario == "" {
			action.PrintUsage(language)
			return
		} else {
			run(scenario)
		}
	}
}

func run(scenario string) {
	command, _ := NewCommand()
	command.Run(scenario)
}

func init() {
	cleanup()
	commandConfig.InitConfig()
	zapLog.Init("command")
}

func cleanup() {
	color.Unset()
}
