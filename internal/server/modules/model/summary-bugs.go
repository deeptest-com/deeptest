package model

type SummaryBugs struct {
	BaseModel
	ProjectId    int64  `json:"projectId"`
	BugId        int64  `gorm:"type:varchar(128)" json:"bugId"`
	Source       string `gorm:"type:varchar(128)" json:"source"`
	BugSeverity  string `gorm:"type:varchar(50)" json:"bugSeverity"`
	BugCreatedAt string `gorm:"type:varchar(90)" json:"bugCreatedAt"`
	BugClassify  string `gorm:"type:varchar(90)" json:"bugClassify"`
	ProductId    int64  `json:"productId"`
	ProductName  string `gorm:"type:varchar(128)" json:"productName"`
	BugState     string `gorm:"type:varchar(50)" json:"bugState"`
}

type SummaryBugsSeverity struct {
	Count       int64  `gorm:"column:count" json:"count"`
	BugSeverity string `gorm:"column:severity" json:"severity"`
}

func (SummaryBugs) TableName() string {
	return "biz_summary_bugs"
}
