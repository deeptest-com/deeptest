package run

// ITestScenario represents interface for testcases,
// includes TestScenario and TestScenarioPath.
type ITestScenario interface {
	GetPath() string
	ToTestScenario() (*TestScenario, error)
}

// TestScenario is a container for one testcase, which is used for testcase runner.
// TestScenario implements ITestScenario interface.
type TestScenario struct {
	Id         uint
	Name       string
	TestStages []IStage
}

// TScenario represents testcase data structure.
// Each testcase includes one public config and several sequential teststages.
type TScenario struct {
	Config     *TConfig  `json:"config" yaml:"config"`
	TestStages []*TStage `json:"teststages" yaml:"teststages"`
}
