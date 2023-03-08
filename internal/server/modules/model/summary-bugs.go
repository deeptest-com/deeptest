package model

type SummaryBugs struct {
	BaseModel
	ProjectId     int64  `json:"project_id"`
	BugId         string `gorm:"type:text" json:"bug_id"`
	Source        string `gorm:"type:text" json:"source"`
	BugSeverity   string `gorm:"type:text" json:"bug_severity"`
	BugCreateDate string `gorm:"type:text" json:"bug_create_date"`
	BugClassify   string `gorm:"type:text" json:"bug_classify"`
	ProductId     string `gorm:"type:text" json:"product_id"`
	ProductName   string `gorm:"type:text" json:"product_name"`
	BugState      string `gorm:"type:text" json:"bug_state"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (SummaryBugs) TableName() string {
	return "biz_summary_bugs"
}
