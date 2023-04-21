package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ServeServerRepo struct {
	*BaseRepo       `inject:""`
	CategoryRepo    *CategoryRepo    `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
}

func NewServerRepo() *ServeServerRepo {
	return &ServeServerRepo{}
}

func (r *ServeServerRepo) Get(id uint) (res model.ServeServer, err error) {
	err = r.DB.Where("NOT deleted").First(&res, id).Error
	return
}
