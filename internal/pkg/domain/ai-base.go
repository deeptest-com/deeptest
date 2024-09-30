package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type AiMeasurementBase struct {
	Name             string `json:"name"`
	Desc             string `json:"desc"`
	Input            string `json:"input"`
	ActualOutput     string `json:"actualOutput"`
	ExpectedOutput   string `json:"expectedOutput"`
	Context          string `json:"context"`
	RetrievalContext string `json:"retrievalContext"`
	ToolsCalled      string `json:"toolsCalled"`
	ExpectedTools    string `json:"expectedTools"`
	Reasoning        string `json:"reasoning"`

	MetricsIds string `json:"metricsIds"`
}

type AiTemplateBase struct {
	Name string `json:"name"`
}

type AiModelBase struct {
	Name string `json:"name"`

	ModelType string `json:"modelType"`
	ModelName string `json:"modelName"`

	ApiUrl string `json:"apiUrl"`
	ApiKey string `json:"apiKey"`
}

type AiMetricsAnswerRelevancyBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`
}

type AiMetricsBase struct {
	Name string `json:"name"`

	Threshold     float32 `json:"threshold" gorm:"default:0.5"`
	IncludeReason bool    `json:"includeReason" gorm:"default:true"`
	AsyncMode     bool    `json:"asyncMode" gorm:"default:false"`
	StrictMode    bool    `json:"strictMode" gorm:"default:false"`
	VerboseMode   bool    `json:"verboseMode" gorm:"default:false"`

	Score       float32 `json:"score" gorm:"default:0"`
	Reason      string  `json:"reason"`
	Success     bool    `json:"success" gorm:"default:false"`
	VerboseLogs string  `json:"verboseLogs"`

	ModelId uint        `json:"modelId"`
	Model   AiModelBase `gorm:"-" json:"model"`

	EntityType consts.MetricsType `json:"entity_type"`
	EntityId   uint               `json:"entity_id"`

	//EntityObj interface{} `json:"entityObj" gorm:"-"`
}
