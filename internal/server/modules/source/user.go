package source

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type UserSource struct {
	UserRepo    *repo2.UserRepo    `inject:""`
	RoleRepo    *repo2.RoleRepo    `inject:""`
	ProjectRepo *repo2.ProjectRepo `inject:""`
}

func (s *UserSource) GetSources(tenantId consts.TenantId) ([]v1.UserReq, v1.ProjectReq, error) {
	roleIds, err := s.RoleRepo.GetRoleIds(tenantId)
	if err != nil {
		return []v1.UserReq{}, v1.ProjectReq{}, err
	}
	users := []v1.UserReq{
		{
			UserBase: v1.UserBase{
				Username: serverConsts.AdminUserName,
				Name:     "管理员",
				Intro:    "超级管理员",
				Avatar:   "upload/images/avatar-m.svg",
				Password: serverConsts.AdminUserPassword,
				RoleIds:  roleIds,
			},
		},
	}

	project := v1.ProjectReq{ProjectBase: v1.ProjectBase{Name: "默认项目", AdminId: 1, ShortName: "T"}}

	return users, project, nil
}

func (s *UserSource) Init(tenantId consts.TenantId) error {
	if s.UserRepo.DB.Model(&model.SysUser{}).Where("id IN ?", []int{1}).Find(&[]model.SysUser{}).RowsAffected == 1 {
		color.Danger.Printf("\n[Mysql] --> %s 表的初始数据已存在!", model.SysUser{}.TableName())
		return nil
	}
	sources, project, err := s.GetSources(tenantId)
	if err != nil {
		return err
	}

	//创建项目
	s.ProjectRepo.Create(tenantId, project, 1)

	for _, source := range sources {
		if _, err := s.UserRepo.Create(tenantId, source); err != nil { // 遇到错误时回滚事务
			return err
		}
	}
	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.SysUser{}.TableName())
	return nil
}
