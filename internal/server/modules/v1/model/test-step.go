package model

import "gorm.io/datatypes"

type TestStep struct {
	BaseModel

	Version int    `json:"version"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Desc    string `json:"desc"`

	Selector string `json:"selector"`
	Value    string `json:"value"`
	TagName  string `json:"tagName"`
	Action   string `json:"action"`
	KeyCode  string `json:"keyCode"`
	Href     string `json:"href"`

	Uuid  string `json:"uuid" gorm:"-"`
	Image string `json:"image"`

	CoordinatesJson datatypes.JSON `json:"-"`
	DimensionsJson  datatypes.JSON `json:"-"`

	Coordinates Coordinates `json:"coordinates" gorm:"-"`
	Dimensions  Dimensions  `json:"dimensions" gorm:"-"`

	ScriptId  uint `json:"scriptId" gorm:"column:script_id"`
	ProjectId uint `json:"projectId"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Left   int `json:"left"`
	Top    int `json:"top"`
}

func (TestStep) TableName() string {
	return "biz_test_step"
}
