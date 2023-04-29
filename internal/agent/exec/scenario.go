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
	Name          string     `json:"name"`
	RootProcessor *Processor `json:"rootProcessor"`

	EnvToVariablesMap map[uint]map[string]domain.EnvVar `json:"envVariables"` // envId -> varId -> varObj
	InterfaceToEnvMap map[uint]uint                     `json:"interfaceToEnvMap"`

	GlobalEnvVars   []domain.GlobalEnvVars   `json:"globalEnvVars"`
	GlobalParamVars []domain.GlobalParamVars `json:"globalParamVars"`

	Datapools domain.Datapools

	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
