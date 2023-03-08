package model

type SummaryProjectUserRanking struct {
	BaseModel

	Sort           string `json:"sort"`
	Hb             int    `json:"hb"`
	ProjectId      string `gorm:"type:text" json:"project_id"`
	User           string `gorm:"type:text" json:"user"`
	ScenarioTotal  string `gorm:"type:text" json:"scenario_total"`
	TestcasesTotal string `gorm:"type:text" json:"testcases_total"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (SummaryProjectUserRanking) TableName() string {
	return "biz_summary_project_user_ranking"
}
