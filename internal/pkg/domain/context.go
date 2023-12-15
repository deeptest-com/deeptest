package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type GlobalVar struct {
	VarId       uint                       `gorm:"-" json:"varId"`
	Name        string                     `json:"name"`
	RightValue  string                     `gorm:"type:text" json:"rightValue"`
	LocalValue  string                     `gorm:"type:text" json:"localValue"`
	RemoteValue string                     `gorm:"type:text" json:"remoteValue"`
	ValueType   consts.ExtractorResultType `json:"valueType"`
}
type GlobalParam struct {
	Name         string           `json:"name"`
	Type         consts.ParamType `json:"type"`
	In           consts.ParamIn   `json:"in"`
	Disabled     bool             `json:"disabled"`
	Required     bool             `json:"required"`
	DefaultValue string           `gorm:"type:text" json:"defaultValue"`
}

type InterfaceToEnvMap map[uint]uint        // interfaceId -> envId
type EnvToVariables map[uint][]GlobalVar    // envId -> vars
type Datapools map[string][]VarKeyValuePair // datapoolName -> array of map<colName, colValue>

type VarKeyValuePair map[string]interface{}
