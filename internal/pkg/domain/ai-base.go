package domain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
)

type AiMeasurementBase struct {
	Input            string `json:"input"`
	ActualOutput     string `json:"actualOutput"`
	RetrievalContext string `json:"retrievalContext"`
}

type AiTemplateBase struct {
	Name string `json:"name"`
}

type AiMetricsEntityBase struct {
	Threshold     float32 `json:"threshold" gorm:"default:0.5"`
	IncludeReason bool    `json:"includeReason" gorm:"default:true"`
	AsyncMode     bool    `json:"asyncMode" gorm:"default:false"`
	StrictMode    bool    `json:"strictMode" gorm:"default:false"`
	VerboseMode   bool    `json:"verboseMode" gorm:"default:false"`

	Score       float32 `json:"score" gorm:"default:0"`
	Reason      string  `json:"reason"`
	Success     bool    `json:"success" gorm:"default:false"`
	VerboseLogs string  `json:"verboseLogs"`

	Output       string              `gorm:"type:longtext;" json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `gorm:"type:longtext" json:"resultMsg"`

	MetricsId   uint               `gorm:"-" json:"metricsId"`
	MetricsType consts.MetricsType `gorm:"-" json:"metricsType"` // for log only
	InvokeId    uint               `gorm:"-" json:"invokeId"`    // for log only

	Disabled bool `json:"disabled"`
}

type AiMetricsBase struct {
	Name string `json:"name"`

	ModelId uint `json:"modelId"`
	Ordr    int  `json:"ordr"`

	EntityType consts.MetricsType `json:"entityType"`

	DebugInterfaceId    uint `gorm:"default:0" json:"debugInterfaceId"`
	EndpointInterfaceId uint `gorm:"default:0" json:"endpointInterfaceId"`

	//EntityObj interface{} `json:"entityObj" gorm:"-"`
}

type ToolModelBase struct {
	ModelProvider consts.LlmType `json:"modelProvider"`
	Name          string         `json:"name"`
	Desc          string         `gorm:"type:text" json:"desc"`

	ApiBase string `json:"apiBase"`
	ApiKey  string `json:"apiKey"`
	Model   string `json:"model"`
	Version string `json:"version"`

	IsDefault bool `json:"isDefault"`
	ProjectId uint `json:"projectId"`
}
