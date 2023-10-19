package agentExec

type CasesExecReq struct {
	ExecUUid string `json:"execUuid"`

	Cases         []CasesExecObj `json:"cases"`
	EnvironmentId int            `json:"environmentId"`
}

type CasesExecObj struct {
	Path      interface{} `json:"path"`
	Sample    interface{} `json:"sample"`
	FieldType interface{} `json:"fieldType"`
	Category  interface{} `json:"Category"`
	Type      interface{} `json:"Type"`
	Rule      interface{} `json:"Rule"`
}
