package domain

type ExecScene struct {
	GlobalVars   []GlobalVar   `json:"globalVar"`
	GlobalParams []GlobalParam `json:"globalParam"`

	ShareVars []GlobalVar `json:"shareVars"`

	InterfaceToEnvMap InterfaceToEnvMap `json:"interfaceToEnvMap"`
	EnvToVariables    EnvToVariables    `json:"envToVariables"`
	Datapools         Datapools         `json:"datapool"`
}
