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
	ScenarioExecObjBase
	RootProcessor *Processor `json:"rootProcessor"`
}

type ScenarioExecObjMsg struct {
	ScenarioExecObjBase
	RootProcessor *ProcessorMsg `json:"rootProcessor"`
}

type ScenarioExecObjBase struct {
	Name string `json:"name"`

	BaseUrl string `json:"baseUrl"`

	EnvToVariablesMap domain.EnvToVariablesMap `json:"envVariables"` // envId -> varId -> varObj
	InterfaceToEnvMap domain.InterfaceToEnvMap `json:"interfaceToEnvMap"`

	GlobalEnvVars   []domain.GlobalEnvVar   `json:"globalEnvVars"`
	GlobalParamVars []domain.GlobalParamVar `json:"globalParamVars"`

	Datapools domain.Datapools

	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
