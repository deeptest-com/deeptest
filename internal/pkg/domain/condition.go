package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ExtractorBase struct {
	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `json:"expression"`
	Prop       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string                `json:"variable"`
	Scope    consts.ExtractorScope `json:"scope" gorm:"default:public"`

	Result       string              `json:"result"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg"`

	ConditionId       uint `json:"conditionId"`
	ConditionEntityId uint `json:"conditionEntityId"`
	InvokeId          uint `json:"invokeId"`

	Disabled bool `json:"disabled"`
}

func (condition ExtractorBase) GetType() consts.ConditionType {
	return consts.ConditionTypeExtractor
}

type CheckpointBase struct {
	Type consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator     consts.ComparisonOperator `json:"operator"`
	Value        string                    `json:"value"`
	ActualResult string                    `json:"actualResult"`

	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg"`

	ConditionId       uint `json:"conditionId"`
	ConditionEntityId uint `json:"conditionEntityId"`
	InvokeId          uint `json:"invokeId"`

	Disabled bool `json:"disabled"`
}

func (condition CheckpointBase) GetType() consts.ConditionType {
	return consts.ConditionTypeCheckpoint
}

type ScriptBase struct {
	ConditionSrc consts.ConditionSrc `json:"conditionType"`

	Content string `json:"content"`

	Output       string              `json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	ResultMsg    string              `json:"resultMsg"`

	ConditionId       uint `json:"conditionId"`
	ConditionEntityId uint `json:"conditionEntityId"`
	InvokeId          uint `json:"invokeId"`

	Disabled bool `json:"disabled"`
}

func (condition ScriptBase) GetType() consts.ConditionType {
	return consts.ConditionTypeScript
}
