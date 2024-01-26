package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type ExtractorBase struct {
	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `gorm:"default:''" json:"expression"`
	Prop       string `json:"prop"`

	BoundaryStart    string `gorm:"default:''" json:"boundaryStart"`
	BoundaryEnd      string `gorm:"default:''" json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string                `gorm:"default:''" json:"variable"`
	Scope    consts.ExtractorScope `json:"scope" gorm:"default:public"`

	Default string `gorm:"default:''" json:"default"` // for cookie

	Result     string                     `json:"result" gorm:"type:text"`
	ResultType consts.ExtractorResultType `json:"resultType"`

	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg" gorm:"type:text"`

	ConditionId         uint                 `json:"conditionId"`
	ConditionEntityId   uint                 `gorm:"-" json:"conditionEntityId"`   // refer to po id in domain object
	ConditionEntityType consts.ConditionType `gorm:"-" json:"conditionEntityType"` // for log only
	InvokeId            uint                 `json:"invokeId"`                     // for log only

	Disabled bool `json:"disabled"`
}

func (condition ExtractorBase) GetType() consts.ConditionType {
	return consts.ConditionTypeExtractor
}

type CheckpointBase struct {
	Type consts.CheckpointType `json:"type"`

	Expression          string               `json:"expression"`
	ExtractorVariable   string               `json:"extractorVariable"`
	ExtractorType       consts.ExtractorType `json:"extractorType"`
	ExtractorExpression string               `json:"extractorExpression"`

	Operator     consts.ComparisonOperator `json:"operator"`
	Value        string                    `json:"value"`
	ActualResult string                    `json:"actualResult" gorm:"type:text"`

	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg" gorm:"type:text"`
	Variables    string              `json:"variables" gorm:"type:text"` // for checkpoint log only

	ConditionId         uint                 `json:"conditionId"`
	ConditionEntityId   uint                 `gorm:"-" json:"conditionEntityId"`   // refer to entity po id in domain object
	ConditionEntityType consts.ConditionType `gorm:"-" json:"conditionEntityType"` // for log only
	InvokeId            uint                 `json:"invokeId"`                     // for log only

	Disabled bool `json:"disabled"`
}

type DatabaseConnBase struct {
	Name string              `json:"name"`
	Type consts.DatabaseType `json:"type"`

	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DatabaseOptBase struct {
	DbConnId uint `json:"dbConnId"`
	DatabaseConnBase
	DatabaseConnIsDisabled bool `json:"databaseConnIsDisabled" gorm:"-"`

	ConditionSrc consts.ConditionSrc `json:"conditionSrc"`

	Sql        string                     `json:"sql"`
	Variable   string                     `json:"variable"`
	Scope      consts.ExtractorScope      `json:"scope" gorm:"default:public"`
	JsonPath   string                     `json:"jsonPath"`
	Result     string                     `json:"result" gorm:"type:text"`
	ResultType consts.ExtractorResultType `json:"resultType"`

	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg" gorm:"type:text"`

	ConditionId         uint                 `json:"conditionId"`
	ConditionEntityId   uint                 `gorm:"-" json:"conditionEntityId"`   // refer to entity po id in domain object
	ConditionEntityType consts.ConditionType `gorm:"-" json:"conditionEntityType"` // for log only
	InvokeId            uint                 `json:"invokeId"`                     // for log only

	Disabled bool `json:"disabled"`
}

func (condition CheckpointBase) GetType() consts.ConditionType {
	return consts.ConditionTypeCheckpoint
}

type ScriptBase struct {
	ConditionSrc consts.ConditionSrc `json:"conditionSrc"`

	Content string `gorm:"type:longtext;" json:"content"`

	Output       string              `gorm:"type:longtext;" json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `gorm:"type:longtext" json:"resultMsg"`

	ConditionId         uint                 `json:"conditionId"`
	ConditionEntityId   uint                 `gorm:"-" json:"conditionEntityId"`   // refer to po id in domain object
	ConditionEntityType consts.ConditionType `gorm:"-" json:"conditionEntityType"` // for log only
	InvokeId            uint                 `json:"invokeId"`                     // for log only

	Disabled bool `json:"disabled"`

	VariableSettings []ExecVariable `gorm:"-" json:"variableSettings"`
}

func (condition ScriptBase) GetType() consts.ConditionType {
	return consts.ConditionTypeScript
}

type ResponseDefineBase struct {
	ResponseCode string   `json:"responseCode"`
	Schema       string   `gorm:"-" json:"schema"`
	Codes        []string `gorm:"-" json:"codes"`
	Code         string   `json:"code"`

	Output       string              `gorm:"type:longtext;" json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg"`

	ConditionId         uint                 `json:"conditionId"`
	ConditionEntityId   uint                 `gorm:"-" json:"conditionEntityId"`   // refer to po id in domain object
	ConditionEntityType consts.ConditionType `gorm:"-" json:"conditionEntityType"` // for log only
	InvokeId            uint                 `json:"invokeId"`                     // for log only
	MediaType           string               `json:"mediaType"`
	Disabled            bool                 `json:"disabled"`
	Component           string               `gorm:"-" json:"component"`
}

func (condition ResponseDefineBase) GetType() consts.ConditionType {
	return consts.ConditionTypeResponseDefine
}
