package serverDomain

type DocumentReq struct {
	EndpointIds []uint `json:"endpointIds"`
	ServeId     []uint `json:"serveId"`
	ProjectId   uint   `json:"projectId"`
}

type DocumentRep struct {
	ProjectReq
	Environment  []EnvironmentReq     `json:"environment"`
	Serves       []DocumentServe      `json:"serves"`
	GlobalParams EnvironmentParamsReq `json:"globalParams"`
	GlobalVars   []EnvironmentParam   `json:"globalVars"`
}

type DocumentServe struct {
	ServeReq
	Servers   []ServeServer    `json:"servers"`
	Component []ServeSchemaReq `json:"component"`
	Endpoints []EndpointReq    `json:"endpoints"`
}
