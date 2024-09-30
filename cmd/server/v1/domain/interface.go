package serverDomain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
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
	IsDir     bool   `json:"isDir"`

	Ordr     int               `json:"ordr"`
	Children []InterfaceSimple `json:"children"`
}

type InterfaceImportReq struct {
	Src  string `json:"src"`
	Type string `json:"type"`
	File string `json:"file"`
}

type OpenApiHeader OpenApiParam
type OpenApiCookie OpenApiParam

type InterfaceResp struct {
	ID             int64                 `json:"id"`
	Url            string                `json:"url"`
	Name           string                `json:"name"`
	OperationId    string                `json:"operationId"`
	Description    string                `json:"description"`
	Security       string                `json:"security"`
	Method         string                `gorm:"default:GET" json:"method"`
	Params         []OpenApiParam        `gorm:"-" json:"params"`
	Headers        []OpenApiHeader       `gorm:"-" json:"headers"`
	Cookies        []OpenApiCookie       `gorm:"-" json:"cookies"`
	RequestBody    domain.RequestBody    `gorm:"default:{}" json:"requestBody"`
	ResponseBodies []domain.ResponseBody `gorm:"default:{}" json:"responseBodies"`
	Body           string                `gorm:"default:{}" json:"body"`
	BodyType       string                `gorm:"default:''" json:"bodyType"`
	ResponseCodes  string                `json:"responseCodes"`
	Mock           []interface{}         `gorm:"-" json:"mock"`
	//AuthorizationType string                 `gorm:"default:''" json:"authorizationType"`
	//PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	//ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	//BasicAuth   domain.BasicAuth   `gorm:"-" json:"basicAuth"`
	//BearerToken domain.BearerToken `gorm:"-" json:"bearerToken"`
	//OAuth20     domain.OAuth20     `gorm:"-" json:"oAuth20"`
	//ApiKey      domain.ApiKey      `gorm:"-" json:"apiKey"`
	ProjectId    uint                 `json:"projectId"`
	GlobalParams []domain.GlobalParam `json:"globalParams"`
}

type InterfaceYapiReq struct {
	Target    int    `json:"target"`
	YapiHost  string `json:"yapiHost"`
	Token     string `json:"token"`
	ProjectId int    `json:"projectId"`
}
