package agentExecDomain

type InterfaceStat struct {
	InterfaceDurationTotal   int64 `json:"interfaceDurationTotal"` // milliseconds
	InterfaceDurationAverage int64 `json:"interfaceDurationAverage"`
	InterfaceCount           int64 `json:"interfaceCount"`

	InterfacePass int `json:"interfacePass"`
	InterfaceFail int `json:"interfaceFail"`
	InterfaceSkip int `json:"interfaceSkip"`

	CheckpointPass int `json:"checkpointPass"`
	CheckpointFail int `json:"checkpointFail"`
}
