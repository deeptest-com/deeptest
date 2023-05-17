package model

type SummaryProjectUserRanking struct {
	BaseModel

	Sort          int64  `json:"sort"`
	ProjectId     int64  `json:"projectId"`
	UserId        int64  `json:"userId"`
	UserName      string `gorm:"type:varchar(90)" json:"userName"`
	ScenarioTotal int64  `json:"scenarioTotal"`
	TestCaseTotal int64  `json:"testCaseTotal"`
}

func (SummaryProjectUserRanking) TableName() string {
	return "biz_summary_project_user_ranking"
}
