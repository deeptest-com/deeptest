package source

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/gookit/color"
)

type RoleSource struct {
	RoleRepo *repo2.RoleRepo `inject:""`
	PermRepo *repo2.PermRepo `inject:""`
}

func NewRoleSource() *RoleSource {
	return &RoleSource{}
}

func (s *RoleSource) GetSources() ([]v1.RoleReq, error) {
	perms, err := s.PermRepo.GetPermsForRoles()
	if err != nil {
		return []v1.RoleReq{}, err
	}

	sources := []v1.RoleReq{
		{
			RoleBase: v1.RoleBase{
				Name:        "admin",
				DisplayName: "管理员",
				Perms:       perms[consts.Admin],
			},
		},
		{
			RoleBase: v1.RoleBase{
				Name:        "user",
				DisplayName: "用户",
				Perms:       perms[consts.User],
			},
		},
	}
	return sources, err
}

func (s *RoleSource) Init() error {
	if s.RoleRepo.DB.Model(&model.SysRole{}).
		Where("id IN ?", []int{1}).
		Find(&[]model.SysRole{}).RowsAffected == 2 {
		color.Danger.Printf("\n[Mysql] --> %s 表的初始数据已存在!", model.SysRole{}.TableName())
		return nil
	}

	sources, err := s.GetSources()
	if err != nil {
		return err
	}

	for _, source := range sources {
		if _, err := s.RoleRepo.Create(source); err != nil { // 遇到错误时回滚事务
			return err
		}
	}

	color.Info.Printf("\n[Mysql] --> %s 表初始数据成功!", model.SysRole{}.TableName())
	return nil
}
