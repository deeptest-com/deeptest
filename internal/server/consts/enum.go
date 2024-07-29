package serverConsts

type NodeCreateMode string

const (
	Brother NodeCreateMode = "brother"
	Child   NodeCreateMode = "child"
)

func (e NodeCreateMode) String() string {
	return string(e)
}

type NodeCreateType string

const (
	Dir  NodeCreateType = "dir"
	Node NodeCreateType = "node"
)

func (e NodeCreateType) String() string {
	return string(e)
}

type DropPos int

const (
	Before DropPos = -1
	Inner  DropPos = 0
	After  DropPos = 1
)

func (e DropPos) Int() int {
	return int(e)
}

type CategoryDiscriminator string

const (
	EndpointCategory CategoryDiscriminator = "endpoint"
	ScenarioCategory CategoryDiscriminator = "scenario"
	PlanCategory     CategoryDiscriminator = "plan"
	SchemaCategory   CategoryDiscriminator = "schema"
)

func (e CategoryDiscriminator) String() string {
	return string(e)
}

type DiagnoseInterfaceType string

const (
	DiagnoseInterfaceTypeDir       DiagnoseInterfaceType = "dir"
	DiagnoseInterfaceTypeInterface DiagnoseInterfaceType = "interface"

	DiagnoseInterfaceTypeWebsocketInterface DiagnoseInterfaceType = "websocket_interface"
	DiagnoseInterfaceTypeGrpcInterface      DiagnoseInterfaceType = "grpc_interface"
)

func (e DiagnoseInterfaceType) String() string {
	return string(e)
}

type AuthType string

const (
	ApiKey      AuthType = "apiKey"
	BearerToken AuthType = "bearerToken"
	BasicAuth   AuthType = "basicAuth"
)

func (e AuthType) String() string {
	return string(e)
}

type ProjectType string

const (
	Full  ProjectType = "full"
	Debug ProjectType = "debug"
)

func (e ProjectType) String() string {
	return string(e)
}

type EndpointCaseTreeType string

const (
	EndpointCaseTreeTypeDir      EndpointCaseTreeType = "dir"
	EndpointCaseTreeTypeEndpoint EndpointCaseTreeType = "endpoint"
	EndpointCaseTreeTypeCase     EndpointCaseTreeType = "case"
)

func (e EndpointCaseTreeType) String() string {
	return string(e)
}

type ProjectSource string

const (
	ProjectSourceLY ProjectSource = "ly"
)

func (e ProjectSource) String() string {
	return string(e)
}
