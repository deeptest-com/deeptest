package runnerExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/runner/metrics"
)

type VuGenerator interface {
	Run(execCtx context.Context, sender metrics.MessageSender) error
}

type VuGeneratorParam struct {
	Type   ptconsts.GenerateType `json:"type,omitempty"`
	Target int                   `json:"target,omitempty"`
	Stages []VuGeneratorStage    `json:"stages,omitempty"`
}
type VuGeneratorStage struct {
	Duration int `json:"duration"`
	Target   int `json:"target"`
}
