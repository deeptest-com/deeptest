package model

type SummaryProjectUserRanking struct {
	BaseModel

	Sort           int64  `json:"sort"`
	ProjectId      int64  `gorm:"type:text" json:"project_id"`
	UserId         int64  `json:"user_id"`
	UserName       string `gorm:"type:text" json:"user_name"`
	ScenarioTotal  int64  `gorm:"type:text" json:"scenario_total"`
	TestcasesTotal int64  `gorm:"type:text" json:"testcases_total"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (SummaryProjectUserRanking) TableName() string {
	return "biz_summary_project_user_ranking"
}
