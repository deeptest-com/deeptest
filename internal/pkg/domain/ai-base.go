package domain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
)

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

type AiMetricsEntityBase struct {
	Name         string              `json:"name"`
	Output       string              `gorm:"type:longtext;" json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `gorm:"type:longtext" json:"resultMsg"`

	MetricsId         uint               `json:"conditionId"`
	MetricsEntityId   uint               `gorm:"-" json:"conditionEntityId"`   // refer to po id in domain object
	MetricsEntityType consts.MetricsType `gorm:"-" json:"conditionEntityType"` // for log only
	InvokeId          uint               `json:"invokeId"`                     // for log only

	Disabled bool `json:"disabled"`
}

func (to AiMetricsEntityBase) SetInfo(id, entityId uint, typ consts.MetricsType, disabled bool) {
	to.Output = ""
	to.MetricsId = id
	to.MetricsEntityId = entityId
	to.MetricsEntityType = typ
	to.Disabled = disabled
}

type EntityToInterface interface {
	SetInfo(id, entityId uint, typ consts.MetricsType, disabled bool)
}

type AiMetricsSummarizationBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsAnswerRelevancyBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsFaithfulnessBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsContextualPrecisionBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsContextualRecallBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsContextualRelevancyBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsHallucinationBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsBiasBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
}
type AiMetricsToxicityBase struct {
	TemplStatements string `json:"templStatements" gorm:"type:text"`
	TemplVerdicts   string `json:"templ_verdicts" gorm:"type:text"`
	TemplReason     string `json:"templReason" gorm:"type:text"`

	AiMetricsEntityBase
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
	Ordr    int         `json:"ordr"`

	EntityType consts.MetricsType `json:"entityType"`
	EntityId   uint               `json:"entityId"`

	DebugInterfaceId    uint `gorm:"default:0" json:"debugInterfaceId"`
	EndpointInterfaceId uint `gorm:"default:0" json:"endpointInterfaceId"`

	//EntityObj interface{} `json:"entityObj" gorm:"-"`
}
