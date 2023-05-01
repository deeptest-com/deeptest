package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type InterfaceExecObj struct {
	DebugData domain.DebugData `json:"debugData"`

	InterfaceToEnvMap domain.InterfaceToEnvMap `json:"interfaceToEnvMap"`
	EnvToVariables    domain.EnvToVariables    `json:"envToVariables"` // envId -> vars

	GlobalVars   []domain.GlobalVar   `json:"globalVars"`
	GlobalParams []domain.GlobalParam `json:"globalParams"`

	Datapools domain.Datapools
}
