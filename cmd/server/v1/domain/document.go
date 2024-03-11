package serverDomain

type DocumentReq struct {
	EndpointIds []uint `json:"endpointIds"`
	ServeIds    []uint `json:"serveIds"`
	ProjectId   uint   `json:"projectId"`
	DocumentId  uint   `json:"documentId"`
	NeedDetail  bool   `json:"needDetail"`
}

type DocumentRep struct {
	ProjectReq
	Environment  []EnvironmentReq       `json:"environment"`
	Serves       []DocumentServe        `json:"serves"`
	GlobalParams map[string]interface{} `json:"globalParams"`
	GlobalVars   []EnvironmentParam     `json:"globalVars"`
	Version      string                 `json:"version"`
	DocumentId   uint                   `json:"documentId"`
	Mock         []interface{}          `json:"mock"`
	Components   []ServeSchemaReq       `json:"components"`
}

type DocumentServe struct {
	ServeReq
	Servers    []ServeServer      `json:"servers"`
	Component  []ServeSchemaReq   `json:"component"`
	Endpoints  []EndpointReq      `json:"endpoints"`
	Securities []ServeSecurityReq `json:"securities"`
}

type Endpoints struct {
	endpoint EndpointReq
	children []*Endpoints
}

type DocumentVersionReq struct {
	EndpointIds []uint `json:"endpointIds"`
	Name        string `json:"name"`
	Version     string `json:"version" validate:"required"`
}

type UpdateDocumentVersionReq struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DocumentVersionListReq struct {
	NeedLatest bool `json:"needLatest"` //需要展示实时版本这条假数据
}

type DocumentShareReq struct {
	ProjectId  uint `json:"projectId"`
	DocumentId uint `json:"documentId"`
	EndpointId uint `json:"endpointId"`
	NeedDetail bool `json:"needDetail"`
}
