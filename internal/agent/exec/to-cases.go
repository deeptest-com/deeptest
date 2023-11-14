package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	casesHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
)

type CasesExecReq struct {
	ExecUuid  string `json:"execUuid"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	ProjectId     uint           `json:"projectId"`
	BaseCaseId    uint           `json:"baseCaseId"`
	UsedBy        consts.UsedBy  `json:"usedBy"`
	ExecUUid      string         `json:"execUuid"`
	Cases         []CasesExecObj `json:"cases"`
	EnvironmentId int            `json:"environmentId"`
}

type CasesExecObj struct {
	ExecUuid   uint          `json:"execUuid"`
	ProjectId  uint          `json:"projectId"`
	BaseCaseId uint          `json:"baseCaseId"`
	UsedBy     consts.UsedBy `json:"usedBy"`

	Key       string                   `json:"key"`
	Path      string                   `json:"path"`
	Sample    interface{}              `json:"sample"`
	FieldType casesHelper.OasFieldType `json:"fieldType"`
	Category  interface{}              `json:"Category"`
	Type      interface{}              `json:"Type"`
	Rule      interface{}              `json:"Rule"`
}

type CaseInterfaceExecObj struct {
	DebugData domain.DebugData `json:"debugData"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`

	ExecScene domain.ExecScene `json:"execScene"`
}
