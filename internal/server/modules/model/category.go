package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

type Category struct {
	BaseModel

	Name  string `json:"name"`
	Desc  string `json:"desc"`
	IsDir bool   `json:"isDir"`

	ParentId  int  `json:"parentId"`
	ProjectId uint `json:"projectId"`
	ServeId   uint `json:"serveId"`
	UseID     uint `json:"useId"`

	Ordr     int          `json:"ordr"`
	Children []*Processor `gorm:"-" json:"children"`
	Slots    iris.Map     `gorm:"-" json:"slots"`

	Type serverConsts.CategoryDiscriminator `gorm:"index:idx_entity_id_type,priority:2" json:"type"`

	SourceType consts.SourceType `json:"sourceType" gorm:"default:0"`
	EntityId   uint              `gorm:"default:0;index:idx_entity_id_type,priority:1" json:"entityId"`
}

func (Category) TableName() string {
	return "biz_category"
}
