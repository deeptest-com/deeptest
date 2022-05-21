package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type TestConfig struct {
	NumberOfThreads int `json:"numberOfThreads,omitempty" yaml:"numberOfThreads,omitempty"`
	RampUpPeriod    int `json:"rampUpPeriod,omitempty" yaml:"rampUpPeriod,omitempty"` // sec

	Duration  int  `json:"duration,omitempty" yaml:"duration,omitempty"` // sec, if set, loopCount will be ignore
	LoopCount int  `json:"loopCount,omitempty" yaml:"loopCount,omitempty"`
	Forever   bool `json:"forever,omitempty" yaml:"forever,omitempty"`

	ErrorAction consts.ErrorAction `json:"errorAction,omitempty" yaml:"errorAction,omitempty" gorm:"-"`
}

func (TestConfig) TableName() string {
	return "test_config"
}
