package source

import (
	"github.com/gookit/color"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"gorm.io/gorm"
)

type PermSource struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermSource() *PermSource {
	return &PermSource{}
}

func (s *PermSource) GetSources() []model.Permission {
	permRouteLen := len(serverConsts.PermRoutes)
	ch := make(chan model.Permission, permRouteLen)

	for _, permRoute := range serverConsts.PermRoutes {
		p := permRoute
		go func(permRoute map[string]string) {
			perm := model.Permission{BasePermission: model.BasePermission{
				Name:        permRoute["path"],
				DisplayName: permRoute["name"],
				Description: permRoute["name"],
				Act:         permRoute["act"],
			}}
			ch <- perm
		}(p)
	}
	perms := make([]model.Permission, permRouteLen)
	for i := 0; i < permRouteLen; i++ {
		perms[i] = <-ch
	}
	return perms
}

func (s *PermSource) Init() error {
	sources := s.GetSources()

	return s.PermRepo.DB.Transaction(func(tx *gorm.DB) error {
		//if tx.Model(&model.Permission{}).Where("id IN ?", []int{1}).Find(&[]model.Permission{}).RowsAffected == 1 {
		//	color.Danger.Printf("\n[Mysql] --> %s 表的初始数据已存在!\n", model.Permission{}.TableName())
		//	return nil
		//}
		//
		//if err := s.PermRepo.CreateInBatches(sources); err != nil { // 遇到错误时回滚事务
		//	return err
		//}

		count, err := s.PermRepo.CreateIfNotExist(sources)
		if err == nil {
			color.Info.Printf("\n[Mysql] --> %s 表成功初始化%d行数据!\n", model.Permission{}.TableName(), count)
		}

		return nil
	})
}
