package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type ProjectMockSetting struct {
	BaseModel

	Priority  consts.MockPriority `json:"priority"`
	ProjectId uint                `json:"projectId"`
}

func (ProjectMockSetting) TableName() string {
	return "biz_project_mock_settings"
}
