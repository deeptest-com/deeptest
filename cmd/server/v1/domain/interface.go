package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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
	ParentId  int    `json:"parentId"`
	ProjectId int    `json:"projectId"`
	IsLeaf    bool   `json:"isLeaf"`

	Ordr     int               `json:"ordr"`
	Children []InterfaceSimple `json:"children"`
}

type InterfaceImportReq struct {
	Src  string `json:"src"`
	Type string `json:"type"`
	File string `json:"file"`
}

type InterfaceResp struct {
	ID             int64                 `json:"id"`
	Url            string                `json:"url"`
	Name           string                `json:"name"`
	OperationId    string                `json:"operationId"`
	Description    string                `json:"description"`
	Security       string                `json:"security"`
	Method         string                `gorm:"default:GET" json:"method"`
	Params         []domain.Param        `gorm:"-" json:"params"`
	Headers        []domain.Header       `gorm:"-" json:"headers"`
	Cookies        []domain.Cookie       `gorm:"-" json:"cookies"`
	RequestBody    domain.RequestBody    `gorm:"default:{}" json:"requestBody"`
	ResponseBodies []domain.ResponseBody `gorm:"default:{}" json:"ResponseBodies"`
	//Body              string                 `gorm:"default:{}" json:"body"`
	//BodyType          consts.HttpContentType `gorm:"default:''" json:"bodyType"`
	//AuthorizationType string                 `gorm:"default:''" json:"authorizationType"`
	//PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	//ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	//BasicAuth   domain.BasicAuth   `gorm:"-" json:"basicAuth"`
	//BearerToken domain.BearerToken `gorm:"-" json:"bearerToken"`
	//OAuth20     domain.OAuth20     `gorm:"-" json:"oAuth20"`
	//ApiKey      domain.ApiKey      `gorm:"-" json:"apiKey"`
}

type Variable struct {
	Id                    uint   `json:"id"`
	Name                  string `json:"name"`
	Value                 string `json:"value"`
	AvailableForCurrScope bool   `json:"availableForCurrScope"`
}

type InterfaceYapiReq struct {
	Target    int    `json:"target"`
	YapiHost  string `json:"yapiHost"`
	Token     string `json:"token"`
	ProjectId int    `json:"projectId"`
}
