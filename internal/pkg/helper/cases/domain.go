package casesHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type AlternativeCase struct {
	Id uint `json:"id"`

	Title string `json:"title"`
	Desc  string `json:"desc"`
	IsDir bool   `json:"isDir"`

	FieldRequired bool         `json:"fieldRequired"`
	FieldType     OasFieldType `json:"fieldType"`

	Sample interface{} `json:"sample"`

	Category consts.AlternativeCaseCategory `json:"category"`
	Type     consts.AlternativeCaseType     `json:"type"`
	ParentId int                            `json:"parentId"`

	Ordr     int                `json:"ordr"`
	Children []*AlternativeCase `json:"children"`
	Slots    iris.Map           `json:"slots"`
}
