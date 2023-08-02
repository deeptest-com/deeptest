package serverDomain

type BatchUpdateReq struct {
	FieldName   string      `json:"fieldName"`
	Value       interface{} `json:"value"`
	EndpointIds []uint      `json:"endpointIds"`
	Ids         []uint      `json:"ids"`
	Module      string      `json:"module" validate:"required"`
}
