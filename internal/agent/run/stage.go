package run

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

// IStage represents interface for all types for teststages, includes:
// StepRequest, StepRequestWithOptionalArgs, StepRequestValidation, StepRequestExtraction,
// StepTestCaseWithOptionalArgs,
// StepTransaction, StepRendezvous, StepWebSocket.
type IStage interface {
	Name() string
	Struct() *TStage
	Run(*SessionRunner) (*StageResult, error)

	Category() consts.ProcessorCategory
}

// TStage represents teststage data structure.
// Each stage maybe three different types: make one request or reference another api/testcase.
type TStage struct {
	Id        uint        `json:"id" yaml:"id"`
	Name      string      `json:"name" yaml:"name"` // required
	Processor interface{} `json:"processor,omitempty" yaml:"processor,omitempty"`

	Variables     map[string]interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`
	SetupHooks    []string               `json:"setupHooks,omitempty" yaml:"setupHooks,omitempty"`
	TeardownHooks []string               `json:"teardownHooks,omitempty" yaml:"teardownHooks,omitempty"`
	Extract       map[string]string      `json:"extract,omitempty" yaml:"extract,omitempty"`
	Validators    []interface{}          `json:"validate,omitempty" yaml:"validate,omitempty"`
	Export        []string               `json:"export,omitempty" yaml:"export,omitempty"`

	Children []IStage `json:"children,omitempty" yaml:"children,omitempty"`
}

type StageResult struct {
	Name          string                   `json:"name" yaml:"name"`                                 // stage name
	StageCategory consts.ProcessorCategory `json:"stageCategory" yaml:"stageCategory"`               // stage Category
	StageType     consts.ProcessorType     `json:"stageType" yaml:"stageType"`                       // stage type
	Success       bool                     `json:"success" yaml:"success"`                           // stage execution result
	Elapsed       int64                    `json:"elapsedMs" yaml:"elapsedMs"`                       // stage execution time in millisecond(ms)
	HttpStat      map[string]int64         `json:"httpstat,omitempty" yaml:"httpstat,omitempty"`     // httpstat in millisecond(ms)
	Data          interface{}              `json:"data,omitempty" yaml:"data,omitempty"`             // session data or slice of stage data
	ContentSize   int64                    `json:"contentSize" yaml:"contentSize"`                   // response body length
	ExportVars    map[string]interface{}   `json:"exportVars,omitempty" yaml:"exportVars,omitempty"` // extract variables
	Attachment    string                   `json:"attachment,omitempty" yaml:"attachment,omitempty"` // stage error information
}
