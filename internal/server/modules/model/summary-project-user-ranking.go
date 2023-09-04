package model

import "time"

type SummaryProjectUserRanking struct {
	BaseModel

	Sort          int64 `json:"sort"`
	ProjectId     int64 `json:"projectId"`
	UserId        int64 `json:"userId"`
	ScenarioTotal int64 `json:"scenarioTotal"`
	TestCaseTotal int64 `json:"testCaseTotal"`
}

type UserTotal struct {
	Count        int64 `json:"count"`
	CreateUserId int64 `json:"createUserId"`
}

type ProjectUserTotal struct {
	UserTotal
	ProjectId int64 `json:"projectId"`
}

type RankingUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ProjectRanking struct {
	Interfaces []UserTotal `json:"-"`
	Scenarios  []UserTotal `json:"-"`
}

type UserUpdateTime struct {
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedBy int64      `json:"createdBy"`
}

func (SummaryProjectUserRanking) TableName() string {
	return "biz_summary_project_user_ranking"
}
