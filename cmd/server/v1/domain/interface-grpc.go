package serverDomain

type GrpcReq struct {
	Address string `json:"address"`
	Service string `json:"service"`
	Method  string `json:"method"`

	UseTls      bool           `json:"useTls"`
	RestartConn bool           `json:"restartConn"`
	MetaData    []MetaDataItem `json:"metaData"`

	IsClientStreaming bool `json:"isClientStreaming"`
	IsServerStreaming bool `json:"isServerStreaming"`

	Message string `json:"message"` // for invoke

	ProtoSrc  string `json:"protoSrc"`
	ProtoName string `json:"protoName"`
	ProtoPath string `json:"protoPath"`

	ExecUuid string `json:"execUuid"`
}
type MetaDataItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GrpcParseResp struct {
	Services []GrpcService `json:"services"`
	Methods  []GrpcMethod  `json:"methods"`
}

type GrpcService struct {
	Name string `json:"name"`
}

type GrpcMethod struct {
	Name string `json:"name"`
}
