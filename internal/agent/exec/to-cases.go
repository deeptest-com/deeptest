package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	casesHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
)

type CasesExecReq struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	UserId     uint          `json:"projectId"`
	ProjectId  uint          `json:"projectId"`
	BaseCaseId uint          `json:"baseCaseId"`
	UsedBy     consts.UsedBy `json:"usedBy"`

	ExecUUid string        `json:"execUuid"`
	ExecObj  *CasesExecObj `json:"cases"`
	ExecType string        `json:"type"`

	EnvironmentId int `json:"environmentId"`
}

type CasesExecObj struct {
	ProjectId  uint          `json:"projectId"`
	BaseCaseId uint          `json:"baseCaseId"`
	UsedBy     consts.UsedBy `json:"usedBy"`
	UserId     uint          `json:"userId"`

	Key       string                   `json:"key"`
	Path      string                   `json:"path"`
	Sample    interface{}              `json:"sample"`
	FieldType casesHelper.OasFieldType `json:"fieldType"`
	Category  interface{}              `json:"Category"`
	Type      interface{}              `json:"Type"`
	Rule      interface{}              `json:"Rule"`
	Title     string                   `json:"title"`
	NeedExec  bool                     `json:"needExec"`
	Children  []*CasesExecObj          `json:"children"`
}

type CaseInterfaceExecObj struct {
	DebugData domain.DebugData `json:"debugData"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`

	ExecScene domain.ExecScene `json:"execScene"`
}

type CaseExecProcessor struct {
	Title    string `json:"tile"`
	Key      string `json:"key"`
	ExecUUid string `json:"execUUid"`
	Category string `json:"type"`

	Children []*CaseExecProcessor `json:"children"`

	Data *InterfaceExecObj `json:"data"`
}
