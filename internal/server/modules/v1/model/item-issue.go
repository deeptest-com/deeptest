package model

import ()

type Issue struct {
	BaseModel
	BaseItem

	// build in
	Severity      string `json:"severity" gorm:"comment:'严重程度'"`
	Origin        string `json:"origin" gorm:"comment:'发现阶段,起源'"`
	Source        string `json:"source" gorm:"comment:'来源'"`
	Cause         string `json:"cause" gorm:"comment:'根源'"`
	DetectWay     string `json:"detectWay" gorm:"comment:'发现方式'"`
	DetectVersion uint   `json:"detectVersion" gorm:"comment:'发现版本'"`
	FixVersion    uint   `json:"fixVersion" gorm:"comment:'修复版本'"`
	Iteration     uint   `json:"iteration" gorm:"comment:'所属迭代'"`
	Module        uint   `json:"module" gorm:"comment:'所属模块'"`
}

func (Issue) TableName() string {
	return "item_issue"
}
