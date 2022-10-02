package run

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

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

func (ts *TestScenario) Category() consts.ProcessorCategory {
	return consts.ProcessorScenario
}
func (ts *TestScenario) Type() consts.ProcessorType {
	return consts.ProcessorScenarioDefault
}

// TScenario represents testcase data structure.
// Each testcase includes one public config and several sequential teststages.
type TScenario struct {
	Config     *TConfig  `json:"config" yaml:"config"`
	TestStages []*TStage `json:"teststages" yaml:"teststages"`
}
