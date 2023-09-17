package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	casesHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
	"github.com/kataras/iris/v12"
)

type AlternativeCase struct {
	Id uint `json:"id"`

	Title string `json:"title"`
	Desc  string `json:"desc"`
	IsDir bool   `json:"isDir"`

	FieldRequired bool                  `json:"fieldRequired"`
	FieldType     casesHelper.FieldType `json:"fieldType"`

	Category consts.AlternativeCaseCategory `json:"category"`
	Type     consts.AlternativeCaseType     `json:"type"`
	ParentId int                            `json:"parentId"`

	Ordr     int                `json:"ordr"`
	Children []*AlternativeCase `gorm:"-" json:"children"`
	Slots    iris.Map           `gorm:"-" json:"slots"`
}
