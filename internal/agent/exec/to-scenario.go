package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type ScenarioExecReq struct {
	ExecUuid  string `json:"execUuid"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	ScenarioId int `json:"scenarioId"`

	EnvironmentId int `json:"environmentId"`
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
	ScenarioId uint   `json:"scenarioId"`
	Name       string `json:"name"`

	ExecScene domain.ExecScene `json:"execScene"`

	ExecUuid  string `json:"execUuid"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
