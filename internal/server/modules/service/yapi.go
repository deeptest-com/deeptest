package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	m "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/mitchellh/mapstructure"
)

type YapiService struct {
	ImportService *ImportService `inject:""`
}

func (s *YapiService) ImportYapiProject(req v1.InterfaceYapiReq) (err error) {
	//yapiHost := req.YapiHost
	//token := req.Token
	//projectId := req.ProjectId
	//target := req.Target
	//yapiCatMenu := getYapiCatMenu(yapiHost, token)
	//yapiCatMenuDatas := yapiCatMenu.Data
	//for i := 0; i < len(yapiCatMenuDatas); i++ {
	//	yapiCatMenuData := yapiCatMenuDatas[i]
	//	interf := m.Interface{}
	//	interf.ProjectId = uint(projectId)
	//	interf.Name = yapiCatMenuData.Name
	//	dropPos := serverConsts.Inner
	//	interf.ParentId, interf.Ordr = s.InterfaceRepo.UpdateOrder(dropPos, uint(target))
	//	err = s.InterfaceRepo.SaveDebugData(&interf)
	//	menuId := interf.ID
	//	catid := strconv.Itoa(yapiCatMenuData.ID)
	//
	//	yapiInterfaceList := s.GetYapiMenuInterfaceList(yapiHost, token, catid)
	//	catMenuInterfaces := yapiInterfaceList.Data.Index
	//	for j := 0; j < len(catMenuInterfaces); j++ {
	//		catMenuInterface := catMenuInterfaces[j]
	//		interfaceId := strconv.Itoa(catMenuInterface.ID)
	//		ret := s.GetYapiInterface(yapiHost, token, interfaceId)
	//		yapiInterf := s.YapiInterfaceInfoToInterf(ret)
	//		yapiInterf.ProjectId = uint(projectId)
	//		yapiInterf.ParentId, yapiInterf.Ordr = s.InterfaceRepo.UpdateOrder(dropPos, menuId)
	//		err = s.ImportService.CreateExpression(&yapiInterf)
	//		if err != nil {
	//			logUtils.Infof("update yapiInterf to db error, %s", err.Error())
	//			return
	//		}
	//	}
	//
	//}
	return

}

func getYapiCatMenu(yapiHost, token string) (yapiCatMenu YapiCatMenu) {
	var param []domain.Param
	tokenParam := domain.Param{
		Name:  "token",
		Value: token}
	param = append(param, tokenParam)

	req := domain.BaseRequest{
		Url:         yapiHost + "/api/interface/getCatMenu",
		QueryParams: &param,
	}
	resp, err := httpHelper.Get(req, nil)
	if err != nil {
		logUtils.Infof("from api get yapi catMenu error, %s", err.Error())
		return
	}
	content := resp.Content
	err1 := json.Unmarshal([]byte(content), &yapiCatMenu)
	if err1 != nil {
		logUtils.Infof("from api get yapi catMenu object error, %s", err1.Error())
		return
	}
	return
}

func (s *YapiService) GetYapiMenuInterfaceList(yapiHost, token, catid string) (yapiInterfaceList YapiInterfaceList) {
	var param []domain.Param
	tokenParam := domain.Param{
		Name:  "token",
		Value: token}
	param = append(param, tokenParam)
	catidParam := domain.Param{
		Name:  "catid",
		Value: catid}
	param = append(param, catidParam)
	limitParam := domain.Param{
		Name:  "limit",
		Value: "1000"}
	param = append(param, limitParam)

	req := domain.BaseRequest{
		Url:         yapiHost + "/api/interface/list_cat",
		QueryParams: &param,
	}
	resp, err := httpHelper.Get(req, nil)
	if err != nil {
		logUtils.Infof("from api get yapi catMenu error, %s", err.Error())
		return
	}
	content := resp.Content
	err1 := json.Unmarshal([]byte(content), &yapiInterfaceList)
	if err1 != nil {
		logUtils.Infof("from api get yapi interfaceList object error, %s", err1.Error())
		//fmt.Println(err1.Error())
		return
	}
	return
}

func (s *YapiService) GetYapiInterface(yapiHost, token, interfaceId string) (ret domain.DebugResponse) {
	var param []domain.Param
	tokenParam := domain.Param{
		Name:  "token",
		Value: token}
	param = append(param, tokenParam)
	interfaceIdParam := domain.Param{
		Name:  "id",
		Value: interfaceId}
	param = append(param, interfaceIdParam)
	req := domain.BaseRequest{
		Url:         yapiHost + "/api/interface/get",
		QueryParams: &param,
	}
	resp, err := httpHelper.Get(req, nil)
	if err != nil {
		logUtils.Infof("from api get yapi interface info error, %s", err.Error())
		return
	}
	return resp
}

func (s *YapiService) YapiInterfaceInfoToInterf(ret domain.DebugResponse) (interf m.EndpointInterface) {
	////content := ret.Content
	//yapiRes := YapiRes{}
	//content := ret.Content
	//err := json.Unmarshal([]byte(content), &yapiRes)
	//if err != nil {
	//	logUtils.Infof("get yapi interface info map error, %s", err.Error())
	//	//fmt.Println(err.Error())
	//}
	//fmt.Println(yapiRes)
	//interf.Name = yapiRes.Data.Title
	//interf.Url = yapiRes.Data.Path
	//interf.Method = getMethod(yapiRes.Data.Method)
	//interf.BodyType = getReqBodyType(yapiRes.Data.ReqBodyType)
	//interf.Body = getReqBodyOther(yapiRes.Data.ReqBodyOther)
	//interf.QueryParams = getReqParams(yapiRes.Data.ReqQuery)
	//interf.Headers = getReqHeaders(yapiRes.Data.ReqHeaders)
	//interf.BodyFormData = getReqBodyForm(yapiRes.Data.ReqBodyForm)
	return
}

func getMethod(yapiMethod string) (method consts.HttpMethod) {
	if yapiMethod == "GET" {
		method = consts.GET
	} else if yapiMethod == "POST" {
		method = consts.POST
	} else if yapiMethod == "PUT" {
		method = consts.PUT
	} else if yapiMethod == "DELETE" {
		method = consts.DELETE
	} else if yapiMethod == "PATCH" {
		method = consts.PATCH
	} else if yapiMethod == "HEAD" {
		method = consts.HEAD
	} else if yapiMethod == "CONNECT" {
		method = consts.CONNECT
	} else if yapiMethod == "OPTIONS" {
		method = consts.OPTIONS
	} else if yapiMethod == "TRACE" {
		method = consts.TRACE
	}
	return
}

func getReqBodyType(yapiBodytype string) (httpContentType consts.HttpContentType) {
	if yapiBodytype == "json" {
		httpContentType = consts.ContentTypeJSON
	} else if yapiBodytype == "form" {
		httpContentType = consts.ContentTypeFormData
	} else if yapiBodytype == "raw" {
		httpContentType = consts.ContentTypeTEXT
	}
	return
}

func getReqBodyOther(yapiReqBodyOther string) (reqBodyOtherStr string) {
	if yapiReqBodyOther == "" {
		return
	}
	reqBodyOther := ReqBodyOther{}
	err := json.Unmarshal([]byte(yapiReqBodyOther), &reqBodyOther)
	if err != nil {
		logUtils.Infof("get yapi interface reqBodyOther error, %s", err.Error())
	}
	reqBodyMap := getJsonbody(reqBodyOther)
	b, err := json.Marshal(reqBodyMap)
	if err != nil {
		logUtils.Infof("get reqBodyMap error, %s", err.Error())
		return
	}
	reqBodyOtherStr = string(b)
	return
}

func getJsonbody(reqBodyOther ReqBodyOther) (reqBodyMap map[string]interface{}) {
	reqBodyMap = make(map[string]interface{})

	//requiredList, ok := required.([]interface{})
	bodyProperties := reqBodyOther.Properties
	bodyPropertiesObject := BodyProperties{}
	for key, value := range bodyProperties {
		if err := mapstructure.Decode(value, &bodyPropertiesObject); err != nil {
			logUtils.Infof("get bodyPropertiesObject error, %s", err.Error())
		}
		if bodyPropertiesObject.Type == "object" {
			propertiesObject := ReqBodyOther{}
			err := mapstructure.Decode(value, &propertiesObject)
			if err != nil {
				logUtils.Infof("get propertiesObject error, %s", err.Error())
			}
			reqBodyMap[key] = getJsonbody(propertiesObject)
		} else {
			reqBodyMap[key] = bodyPropertiesObject.Type
		}

	}
	return
}

func getReqParams(reqQuerys []ReqQuery) (param []m.DebugInterfaceParam) {
	for i := 0; i < len(reqQuerys); i++ {
		interfaceParam := m.DebugInterfaceParam{}
		interfaceParam.Name = reqQuerys[i].Name
		param = append(param, interfaceParam)
	}
	return
}

func getReqHeaders(reqHeaders []ReqHeaders) (header []m.DebugInterfaceHeader) {
	for i := 0; i < len(reqHeaders); i++ {
		interfaceHeader := m.DebugInterfaceHeader{}
		interfaceHeader.Name = reqHeaders[i].Name
		interfaceHeader.Value = reqHeaders[i].Value
		header = append(header, interfaceHeader)
	}
	return
}

func getReqBodyForm(reqBodyForm []ReqBodyForm) (bodyFormData []m.DebugInterfaceBodyFormDataItem) {
	for i := 0; i < len(reqBodyForm); i++ {
		bodyForm := m.DebugInterfaceBodyFormDataItem{}
		bodyForm.Name = reqBodyForm[i].Name
		if reqBodyForm[i].Type == "text" {
			bodyForm.Type = consts.FormDataTypeText
		} else if reqBodyForm[i].Type == "file" {
			bodyForm.Type = consts.FormDataTypeFile
		}
		bodyFormData = append(bodyFormData, bodyForm)
	}
	return
}

type ReqBodyOther struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Title      string                 `json:"title"`
	Ref        string                 `json:"$$ref"`
	Required   []string               `json:"required"`
}

type BodyProperties struct {
	Type        string        `json:"type"`
	Description string        `json:"description"`
	Format      string        `json:"format"`
	Enum        []interface{} `json:"enum"`
	EnumDesc    string        `json:"enumDesc"`
}

type YapiRes struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    Data   `json:"data"`
}
type QueryPath struct {
	Path   string        `json:"path"`
	Params []interface{} `json:"params"`
}
type ReqBodyForm struct {
	Required string `json:"required"`
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Type     string `json:"type"`
}
type ReqHeaders struct {
	Required string `json:"required"`
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}
type ReqQuery struct {
	Required string `json:"required"`
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
}
type Data struct {
	QueryPath           QueryPath     `json:"query_path"`
	EditUID             int           `json:"edit_uid"`
	Status              string        `json:"status"`
	Type                string        `json:"type"`
	ReqBodyIsJSONSchema bool          `json:"req_body_is_json_schema"`
	ResBodyIsJSONSchema bool          `json:"res_body_is_json_schema"`
	APIOpened           bool          `json:"api_opened"`
	Index               int           `json:"index"`
	Tag                 []string      `json:"tag"`
	ID                  int           `json:"_id"`
	Method              string        `json:"method"`
	Title               string        `json:"title"`
	Path                string        `json:"path"`
	ReqParams           []interface{} `json:"req_params"`
	ReqBodyForm         []ReqBodyForm `json:"req_body_form"`
	ReqHeaders          []ReqHeaders  `json:"req_headers"`
	ReqQuery            []ReqQuery    `json:"req_query"`
	ReqBodyType         string        `json:"req_body_type"`
	ResBodyType         string        `json:"res_body_type"`
	ResBody             string        `json:"res_body"`
	ProjectID           int           `json:"project_id"`
	Catid               int           `json:"catid"`
	UID                 int           `json:"uid"`
	AddTime             int           `json:"add_time"`
	UpTime              int           `json:"up_time"`
	V                   int           `json:"__v"`
	Desc                string        `json:"desc"`
	Markdown            string        `json:"markdown"`
	ReqBodyOther        string        `json:"req_body_other"`
	Username            string        `json:"username"`
}

type YapiInterfaceList struct {
	Errcode int                   `json:"errcode"`
	Errmsg  string                `json:"errmsg"`
	Data    YapiInterfaceListData `json:"data"`
}
type YapiInterfaceListDataList struct {
	EditUID   int      `json:"edit_uid"`
	Status    string   `json:"status"`
	APIOpened bool     `json:"api_opened"`
	Tag       []string `json:"tag"`
	ID        int      `json:"_id"`
	Method    string   `json:"method"`
	Title     string   `json:"title"`
	Path      string   `json:"path"`
	ProjectID int      `json:"project_id"`
	Catid     int      `json:"catid"`
	UID       int      `json:"uid"`
	AddTime   int      `json:"add_time"`
}
type YapiInterfaceListData struct {
	Count int                         `json:"count"`
	Total int                         `json:"total"`
	List  []YapiInterfaceListDataList `json:"list"`
}

type YapiCatMenu struct {
	Errcode int               `json:"errcode"`
	Errmsg  string            `json:"errmsg"`
	Data    []YapiCatMenuData `json:"data"`
}
type YapiCatMenuData struct {
	Index     int    `json:"index"`
	ID        int    `json:"_id"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	Desc      string `json:"desc"`
	UID       int    `json:"uid"`
	AddTime   int    `json:"add_time"`
	UpTime    int    `json:"up_time"`
	V         int    `json:"__v"`
}
