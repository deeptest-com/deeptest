package serverDomain

type DocumentReq struct {
	EndpointIds []uint `json:"endpointIds"`
	ServeIds    []uint `json:"serveIds"`
	ProjectId   uint   `json:"projectId"`
}

type DocumentRep struct {
	ProjectReq
	Environment  []EnvironmentReq       `json:"environment"`
	Serves       []DocumentServe        `json:"serves"`
	GlobalParams map[string]interface{} `json:"globalParams"`
	GlobalVars   []EnvironmentParam     `json:"globalVars"`
}

type DocumentServe struct {
	ServeReq
	Servers    []ServeServer      `json:"servers"`
	Component  []ServeSchemaReq   `json:"component"`
	Endpoints  []EndpointReq      `json:"endpoints"`
	Securities []ServeSecurityReq `json:"securities"`
}
