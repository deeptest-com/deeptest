package main

import (
	"fmt"
	commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	pth := ""
	var cmd *exec.Cmd
	if commonUtils.IsWin() {
		pth = "gui\\deeptest.exe"
		cmd = exec.Command("cmd", "/C", "start", pth)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start gui, path %s, err %s", pth, err.Error())
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
