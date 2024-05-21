package grpcDomain

type Desc struct {
	Schema            string `json:"schema"`
	Template          string `json:"template"`
	IsClientStreaming bool   `json:"isClientStreaming"`
	IsServerStreaming bool   `json:"isServerStreaming"`
}

type InvRes struct {
	Time   string `json:"timer"`
	Result string `json:"result"`
}
