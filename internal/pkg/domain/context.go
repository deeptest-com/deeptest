package domain

type GlobalEnvVar map[string]interface{}
type GlobalParamVar map[string]interface{}

type EnvToVariablesMap map[uint]map[string]EnvVar // envId -> varName -> varObj
type InterfaceToEnvMap map[uint]uint              // interfaceId -> envId

type EnvVar map[string]interface{}
type ShareVars map[string]interface{}

type Datapools map[string][]map[string]interface{} // datapoolName -> obj array
