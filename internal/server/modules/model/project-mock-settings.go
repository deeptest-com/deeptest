package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	mockData "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
)

type ProjectMockSetting struct {
	BaseModel

	Priority    consts.MockPriority      `json:"priority"`
	UseExamples mockData.UseExamplesEnum `gorm:"-" json:"useExamplesEnum"`

	ProjectId uint `json:"projectId"`
}

func (ProjectMockSetting) TableName() string {
	return "biz_project_mock_settings"
}
