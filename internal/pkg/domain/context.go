package domain

type GlobalEnvVar map[string]interface{}
type GlobalParamVar map[string]interface{}

type EnvToVariablesMap map[uint]map[string]VarKeyValuePair // envId -> varName -> varObj
type InterfaceToEnvMap map[uint]uint                       // interfaceId -> envId

type VarKeyValuePair map[string]interface{}

type Datapools map[string][]map[string]interface{} // datapoolName -> obj array
