package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	casesHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
)

type CasesExecReq struct {
	ExecUuid  string `json:"execUuid"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	UserId     uint          `json:"projectId"`
	ProjectId  uint          `json:"projectId"`
	BaseCaseId uint          `json:"baseCaseId"`
	UsedBy     consts.UsedBy `json:"usedBy"`

	ExecObj  casesHelper.AlternativeCase `json:"cases"`
	ExecType string                      `json:"type"`

	EnvironmentId uint `json:"environmentId"`
}

//type CaseInterfaceExecObj struct {
//	DebugData domain.DebugData `json:"debugData"`
//
//	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
//	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`
//
//	ExecScene domain.ExecScene `json:"execScene"`
//}

type CaseExecProcessor struct {
	Title    string `json:"tile"`
	Key      string `json:"key"`
	ExecUUid string `json:"execUUid"`
	Category string `json:"type"`

	Children []*CaseExecProcessor `json:"children"`

	Data *InterfaceExecObj `json:"data"`
}
