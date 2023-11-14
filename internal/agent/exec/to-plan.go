package agentExec

type PlanExecReq struct {
	ExecUuid      string `json:"execUuid"`
	ServerUrl     string `json:"serverUrl"`
	Token         string `json:"token"`
	PlanId        int    `json:"planId"`
	EnvironmentId int    `json:"environmentId"`
}

type PlanExecObj struct {
	Name      string            `json:"name"`
	Scenarios []ScenarioExecObj `json:"scenarios"`

	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
