package request

import (
	"fmt"
)

// ITestCase represents interface for testcases,
// includes TestCase and TestCasePath.
type ITestCase interface {
	GetPath() string
	ToTestCase() (*TestCase, error)
}

// TestCase is a container for one testcase, which is used for testcase runner.
// TestCase implements ITestCase interface.
type TestCase struct {
	TestSteps []IStep
}

func (tc *TestCase) ToTestCase() (*TestCase, error) {
	return tc, nil
}

// TestCasePath implements ITestCase interface.
type TestCasePath string

func (path *TestCasePath) GetPath() string {
	return fmt.Sprintf("%v", *path)
}

// TCase represents testcase data structure.
// Each testcase includes one public config and several sequential teststeps.
type TCase struct {
	Config    *TConfig  `json:"config" yaml:"config"`
	TestSteps []*TStage `json:"teststeps" yaml:"teststeps"`
}
