package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type CasesExecReq struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	BaseCaseId    uint           `json:"baseCaseId"`
	ExecUUid      string         `json:"execUuid"`
	Cases         []CasesExecObj `json:"cases"`
	EnvironmentId int            `json:"environmentId"`
}

type CasesExecObj struct {
	BaseCaseId uint `json:"baseCaseId"`

	Path      interface{} `json:"path"`
	Sample    interface{} `json:"sample"`
	FieldType interface{} `json:"fieldType"`
	Category  interface{} `json:"Category"`
	Type      interface{} `json:"Type"`
	Rule      interface{} `json:"Rule"`
}

type CaseInterfaceExecObj struct {
	DebugData domain.DebugData `json:"debugData"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`

	ExecScene domain.ExecScene `json:"execScene"`
}
