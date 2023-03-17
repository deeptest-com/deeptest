package agentExec

type PlanExecReq struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
	PlanId    uint   `json:"planId"`
}

type PlanExecObj struct {
	Scenarios []ProcessorExecObj `json:"scenarios"`
	ServerUrl string             `json:"serverUrl"`
	Token     string             `json:"token"`
}
