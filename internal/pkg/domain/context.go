package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type GlobalVar struct {
	Name        string `json:"name"`
	RightValue  string `json:"rightValue"`
	LocalValue  string `json:"localValue"`
	RemoteValue string `json:"remoteValue"`
}
type GlobalParam struct {
	Name         string           `json:"name"`
	Type         consts.ParamType `json:"type"`
	In           consts.ParamIn   `json:"in"`
	Required     bool             `json:"Required"`
	DefaultValue string           `json:"defaultValue"`
}

type EnvToVariablesMap map[uint]map[string]VarKeyValuePair // envId -> varName -> varObj
type InterfaceToEnvMap map[uint]uint                       // interfaceId -> envId

type VarKeyValuePair map[string]interface{}

type Datapools map[string][]map[string]interface{} // datapoolName -> obj array
