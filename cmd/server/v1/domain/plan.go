package serverDomain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"time"
)

type PlanReqPaginate struct {
	_domain.PaginateReq

	ProjectId  uint   `json:"projectId"`
	CategoryId int64  `json:"categoryId"`
	Status     string `json:"status"`
	AdminId    string `json:"adminId"`
	Keywords   string `json:"keywords"`
	Enabled    string `json:"enabled"`
}

//type PlanAddScenariosReq struct {
//	SelectedNodes []ScenarioSimple `json:"selectedNodes"`
//
//	TargetId  uint `json:"targetId"`
//	ProjectId int  `json:"projectId"`
//}

type PlanAddScenariosReq struct {
	ScenarioIds []uint `json:"scenarioIds"`
}

type PlanAndReportDetail struct {
	Id             uint              `json:"id"`        //计划ID
	AdminName      string            `json:"adminName"` //负责人姓名
	CreatedAt      *time.Time        `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time        `json:"updatedAt,omitempty"`
	UpdateUserName string            `json:"updateUserName"` //最近更新人姓名
	CreateUserName string            `json:"createUserName"` //创建人姓名
	Status         consts.TestStatus `json:"status"`         //状态
	TestPassRate   string            `json:"testPassRate"`   //执行通过率
	ExecTimes      int64             `json:"execTimes"`      //执行次数
	ExecutorName   string            `json:"executorName"`   //执行人姓名
	ExecTime       *time.Time        `json:"execTime"`       //执行时间
	ExecEnv        string            `json:"execEnv"`        //执行环境
	CurrEnvId      uint              `json:"currEnvId"`
}

type PlanScenariosReqPaginate struct {
	_domain.PaginateReq

	CreateUserId uint   `json:"createUserId"`
	Priority     string `json:"priority"`
	Keywords     string `json:"keywords"`
	Enabled      string `json:"enabled"`
}

type NotRelationScenarioReqPaginate struct {
	_domain.PaginateReq

	PlanId       uint   `json:"planId"`
	Keywords     string `json:"keywords"`
	Enabled      string `json:"enabled"`
	Status       string `json:"status"`
	Priority     string `json:"priority"`
	Type         string `json:"type"`
	CreateUserId uint   `json:"createUserId"`
}

type MoveReq struct {
	PlanId        uint `json:"planId"`
	SourceId      uint `json:"sourceId"`
	DestinationId uint `json:"destinationId"`
}
