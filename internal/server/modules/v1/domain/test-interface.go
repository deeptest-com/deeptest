package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
)

type TestInterfaceReq struct {
	Mode      serverConsts.NodeCreateMode `json:"mode"`
	Type      serverConsts.NodeCreateType `json:"type"`
	Target    int                         `json:"target"`
	Name      string                      `json:"name"`
	Id        int                         `json:"id"`
	ProjectId int                         `json:"projectId"`
}

type TestInterfaceMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}

type TestInterfaceResp struct {
	Url               string          `json:"url"`
	Method            string          `gorm:"default:GET" json:"method"`
	Params            []domain.Param  `gorm:"-" json:"params"`
	Headers           []domain.Header `gorm:"-" json:"headers"`
	Body              string          `gorm:"default:{}" json:"body"`
	BodyType          string          `gorm:"default:''" json:"bodyType"`
	AuthorizationType string          `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string          `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string          `gorm:"default:''" json:"validationScript"`

	BasicAuth   domain.BasicAuth   `gorm:" -" json:"basicAuth"`
	BearerToken domain.BearerToken `gorm:" -" json:"bearerToken"`
	OAuth20     domain.OAuth20     `gorm:" -" json:"oAuth20"`
	ApiKey      domain.ApiKey      `gorm:" -" json:"apiKey"`
}
