package agentExec

type MessageExecReq struct {
	UserId    uint   `json:"userId"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}

type MessageExecObj struct {
	Count uint `json:"count"`
}
