package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"sync"
)

var (
	breakMap sync.Map
)

type ScenarioExecReq struct {
	ServerUrl  string `json:"serverUrl"`
	Token      string `json:"token"`
	ScenarioId int    `json:"scenarioId"`
}

type ScenarioExecObj struct {
	RootProcessor *Processor       `json:"rootProcessor"`
	Variables     domain.Variables `json:"variables"`
	Datapools     domain.Datapools
	ServerUrl     string `json:"serverUrl"`
	Token         string `json:"token"`
}
