package serverDomain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
)

type InterfaceReq struct {
	Mode      serverConsts.NodeCreateMode `json:"mode"`
	Type      serverConsts.NodeCreateType `json:"type"`
	Target    int                         `json:"target"`
	Name      string                      `json:"name"`
	Id        int                         `json:"id"`
	ProjectId int                         `json:"projectId"`
}

type InterfaceMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}

type InterfaceSimple struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ProjectId int    `json:"projectId"`
	IsDir     bool   `json:"isDir"`

	Ordr     int               `json:"ordr"`
	Children []InterfaceSimple `json:"children"`
}

/*type InterfaceResp struct {
	Url               string          `json:"url"`
	Method            string          `gorm:"default:GET" json:"method"`
	Params            []domain.Param  `gorm:"-" json:"params"`
	Headers           []domain.Header `gorm:"-" json:"headers"`
	Body              string          `gorm:"default:{}" json:"body"`
	BodyType          consts.HttpContentType          `gorm:"default:''" json:"bodyType"`
	AuthorizationType string          `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string          `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string          `gorm:"default:''" json:"validationScript"`

	BasicAuth   domain.BasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken domain.BearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     domain.OAuth20     `gorm:"-" json:"oAuth20"`
	ApiKey      domain.ApiKey      `gorm:"-" json:"apiKey"`
}*/

type Variable struct {
	id    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
