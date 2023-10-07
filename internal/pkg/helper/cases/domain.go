package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type AlternativeCase struct {
	Key string `json:"key"`

	Title string `json:"title"`
	Desc  string `json:"desc"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`

	FieldRequired bool         `json:"fieldRequired"`
	FieldType     OasFieldType `json:"fieldType"`

	Sample interface{} `json:"sample"`

	Category consts.AlternativeCaseCategories `json:"category"`
	Type     consts.AlternativeCaseTypes      `json:"type"`
	Rule     consts.AlternativeCaseRules      `json:"rule"`
	ParentId int                              `json:"parentId"`

	Ordr     int                `json:"ordr"`
	Children []*AlternativeCase `json:"children"`
	Slots    iris.Map           `json:"slots"`
}
