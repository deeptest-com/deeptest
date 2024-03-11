package agentExec

import ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"

type ProcessorPerformanceGoal struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Type ptconsts.GoalType `json:"type,omitempty"`

	Duration int `json:"duration,omitempty"`
	Loop     int `json:"loop,omitempty"`

	ResponseTime float32 `json:"responseTime,omitempty"`
	Qps          float32 `json:"qps,omitempty"`
	FailRate     float32 `json:"failRate,omitempty"`
}

func (entity ProcessorPerformanceGoal) Run(processor *Processor, session *Session) (err error) {
	return
}
