package agentDomain

type PerformanceRunnerSelectionReq struct {
	Weight uint  `json:"weight"`
	Ids    []int `json:"ids"`

	ProjectId  int `json:"projectId"`
	ScenarioId int `json:"scenarioId"`
}
