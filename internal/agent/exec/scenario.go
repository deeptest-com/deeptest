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

	InterfaceToEnvMap domain.InterfaceToEnvMap `json:"interfaceToEnvMap"`
	EnvToVariables    domain.EnvToVariables    `json:"envToVariables"` // envId -> vars

	GlobalVars   []domain.GlobalVar   `json:"globalVars"`
	GlobalParams []domain.GlobalParam `json:"globalParams"`

	Datapools domain.Datapools

	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
