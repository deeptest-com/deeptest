package ptconsts

type TestRole string

const (
	Conductor TestRole = "conductor"
	Runner    TestRole = "runner"
)

func (e TestRole) String() string {
	return string(e)
}

type ResultStatus string

const (
	Pass    ResultStatus = "pass"
	Fail    ResultStatus = "fail"
	Error   ResultStatus = "error"
	Block   ResultStatus = "block"
	Unknown ResultStatus = "unknown"
)

func (e ResultStatus) String() string {
	return string(e)
}

type MsgCategory string

const (
	MsgCategoryInstruction MsgCategory = "instruction"
	MsgCategoryResult      MsgCategory = "result"
	MsgCategoryLog         MsgCategory = "log"
)

func (e MsgCategory) String() string {
	return string(e)
}

type MsgInstructionServerToRunner string

const (
	MsgInstructionJoinExist      MsgInstructionServerToRunner = "joinExist"
	MsgInstructionStart          MsgInstructionServerToRunner = "start"
	MsgInstructionEnd            MsgInstructionServerToRunner = "end"
	MsgInstructionTerminal       MsgInstructionServerToRunner = "terminal"
	MsgInstructionAlreadyRunning MsgInstructionServerToRunner = "alreadyRunning"
	MsgInstructionException      MsgInstructionServerToRunner = "exception"
)

func (e MsgInstructionServerToRunner) String() string {
	return string(e)
}

type MsgInstructionRunnerToServer string

const (
	MsgInstructionRunnerFinish MsgInstructionRunnerToServer = "runnerFinish"
)

func (e MsgInstructionRunnerToServer) String() string {
	return string(e)
}

type MsgResultTypeToWsClient string

const (
	MsgResultRecord MsgResultTypeToWsClient = "record"
)

func (e MsgResultTypeToWsClient) String() string {
	return string(e)
}

type GenerateType string

const (
	GeneratorConstant GenerateType = "constant"
	GeneratorRamp     GenerateType = "ramp"
)

func (e GenerateType) String() string {
	return string(e)
}

type GoalType string

const (
	GoalTypeDuration     GoalType = "duration"
	GoalTypeLoop         GoalType = "loop"
	GoalTypeResponseTime GoalType = "responseTime"
	GoalTypeQps          GoalType = "qps"
	GoalTypeFailRate     GoalType = "failRate"
)

func (e GoalType) String() string {
	return string(e)
}

type ExecType string

const (
	ExecStop ExecType = "stop"

	Init     ExecType = "init"
	ExecPlan ExecType = "execPlan"
)

func (e ExecType) String() string {
	return string(e)
}

type ExecMode string

const (
	Parallel ExecMode = "parallel"
	Serial   ExecMode = "serial"
)

func (e ExecMode) String() string {
	return string(e)
}

type ProcessorType string

const (
	Interface  ResultStatus = "interface"
	Rendezvous ResultStatus = "rendezvous"
)

func (e ProcessorType) String() string {
	return string(e)
}

type ChartType string

const (
	ChartSummaryVuCount ChartType = "vuCount"

	ChartSummaryStatusCount  ChartType = "statusCount"
	ChartSummaryResponseTime ChartType = "responseTime"
	ChartSummaryQps          ChartType = "qps"

	ChartRespTime ChartType = "respTime"

	ChartQps ChartType = "qps"

	ChartCpuUsage    ChartType = "CpuUsage"
	ChartMemoryUsage ChartType = "MemoryUsage"

	ChartDiskUsages    ChartType = "diskUsages"
	ChartNetworkUsages ChartType = "networkUsages"
)

func (e ChartType) String() string {
	return string(e)
}
