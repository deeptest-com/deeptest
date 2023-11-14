package agentExec

type MessageExecReq struct {
	ExecUuid  string `json:"execUuid"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}

type MessageExecObj struct {
	Count uint `json:"count"`
}
