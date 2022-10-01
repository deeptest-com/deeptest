package main

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/command/action"
	commandConsts "github.com/aaronchen2k/deeptest/internal/command/consts"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
)

type Command struct {
	RunAction *action.RunAction `inject:""`
}

func NewCommand() (command *Command, err error) {
	var g inject.Graph

	command = &Command{}

	// inject objects
	if err := g.Provide(
		&inject.Object{Value: commandConsts.DB},
		&inject.Object{Value: command},
	); err != nil {
		_logUtils.Errorf("provide usecase objects to the Graph: %v", err)
	}
	err = g.Populate()
	if err != nil {
		_logUtils.Errorf(fmt.Sprintf("populate the incomplete Objects: %v", err))
	}

	return
}

func (c *Command) Run(scenarioIdOrName string) {
	c.RunAction.Run(scenarioIdOrName)
}
