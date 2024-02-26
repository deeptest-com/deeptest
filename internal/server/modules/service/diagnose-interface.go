package service

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	curlHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/gcurl"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"github.com/jinzhu/copier"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type DiagnoseInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	ExtractorRepo         *repo.ExtractorRepo         `inject:""`
	CheckpointRepo        *repo.CheckpointRepo        `inject:""`

	DebugInterfaceService *DebugInterfaceService `inject:""`
}

func (s *DiagnoseInterfaceService) Load(tenantId consts.TenantId, projectId int) (ret []*serverDomain.DiagnoseInterface, err error) {
	root, err := s.DiagnoseInterfaceRepo.GetTree(tenantId, uint(projectId))

	if root != nil {
		ret = root.Children
	}

	return
}

func (s *DiagnoseInterfaceService) Get(tenantId consts.TenantId, id int) (ret model.DiagnoseInterface, err error) {
	ret, err = s.DiagnoseInterfaceRepo.Get(tenantId, uint(id))
	// its debug data will load in webpage

	return
}

func (s *DiagnoseInterfaceService) Save(tenantId consts.TenantId, req serverDomain.DiagnoseInterfaceSaveReq) (diagnoseInterface model.DiagnoseInterface, err error) {
	s.CopyValueFromRequest(tenantId, &diagnoseInterface, req)

	//if diagnoseInterface.ID != 0 {
	//	oldDiagnoseInterface, err := s.DiagnoseInterfaceRepo.Get(diagnoseInterface.ID)
	//	if err != nil {
	//		return diagnoseInterface, err
	//	}
	//	diagnoseInterface.CreatedBy = oldDiagnoseInterface.CreatedBy
	//}

	if diagnoseInterface.Type == serverConsts.DiagnoseInterfaceTypeInterface {
		if req.ID == 0 {
			//server, _ := s.ServeServerRepo.GetDefaultByServe(diagnoseInterface.ServeId)
			debugInterface := model.DebugInterface{
				InterfaceBase: model.InterfaceBase{
					Name: req.Title,
					InterfaceConfigBase: model.InterfaceConfigBase{
						//			Url:    server.Url,
						Method: consts.GET,
					},
					ProjectId: req.ProjectId,
				},
				//	ServeId:  diagnoseInterface.ServeId,
				//	ServerId: server.ID,
				BaseUrl: "",
			}
			err = s.DebugInterfaceRepo.Save(tenantId, &debugInterface)
			diagnoseInterface.DebugInterfaceId = debugInterface.ID
			diagnoseInterface.Method = debugInterface.Method

			err = s.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)

			if diagnoseInterface.DebugInterfaceId > 0 {
				values := map[string]interface{}{
					"diagnose_interface_id": diagnoseInterface.ID,
				}
				err = s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, diagnoseInterface.DebugInterfaceId, values)
			}
		} else {
			diagnoseInterface, _ = s.DiagnoseInterfaceRepo.Get(tenantId, req.ID)
			diagnoseInterface.Title = req.Title

			err = s.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)
		}

		// create new DebugInterface
		//debugInterface := model.DebugInterface{
		//	InterfaceBase: model.InterfaceBase{
		//		Name: req.Title,
		//		InterfaceConfigBase: model.InterfaceConfigBase{
		//			//Url:    server.Url,
		//			Method: consts.GET,
		//		},
		//	},
		//	ServeId: diagnoseInterface.ServeId,
		//	//ServerId: server.ID,
		//	BaseUrl: "",
		//}

	} else {
		if req.ID != 0 {
			diagnoseInterface, _ = s.DiagnoseInterfaceRepo.Get(tenantId, req.ID)
			diagnoseInterface.Title = req.Title
		}
		err = s.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)
	}

	return
}

func (s *DiagnoseInterfaceService) Remove(tenantId consts.TenantId, id int, typ serverConsts.DiagnoseInterfaceType) (err error) {
	err = s.DiagnoseInterfaceRepo.Remove(tenantId, uint(id), typ)
	return
}

func (s *DiagnoseInterfaceService) Move(tenantId consts.TenantId, srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcScenarioNode model.DiagnoseInterface, err error) {
	srcScenarioNode, err = s.DiagnoseInterfaceRepo.Get(tenantId, srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.DiagnoseInterfaceRepo.UpdateOrder(tenantId, pos, targetId, projectId)
	err = s.DiagnoseInterfaceRepo.UpdateOrdAndParent(tenantId, srcScenarioNode)

	return
}

func (s *DiagnoseInterfaceService) CopyValueFromRequest(tenantId consts.TenantId, interf *model.DiagnoseInterface, req serverDomain.DiagnoseInterfaceSaveReq) {
	copier.CopyWithOption(interf, req, copier.Option{
		DeepCopy: true,
	})
}

func (s *DiagnoseInterfaceService) CopyDebugDataValueFromRequest(tenantId consts.TenantId, interf *model.DiagnoseInterface, req domain.DebugData) (err error) {
	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}

func (s *DiagnoseInterfaceService) ImportInterfaces(tenantId consts.TenantId, req serverDomain.DiagnoseInterfaceImportReq) (ret model.DiagnoseInterface, err error) {
	parent, _ := s.DiagnoseInterfaceRepo.Get(tenantId, req.TargetId)

	if parent.Type != serverConsts.DiagnoseInterfaceTypeDir {
		parent, _ = s.DiagnoseInterfaceRepo.Get(tenantId, parent.ParentId)
	}

	for _, interfaceId := range req.InterfaceIds {
		ret, err = s.createInterfaceFromDefine(tenantId, interfaceId, req.CreateBy, parent)
	}

	return
}

func (s *DiagnoseInterfaceService) createInterfaceFromDefine(tenantId consts.TenantId, endpointInterfaceId int, createBy uint, parent model.DiagnoseInterface) (
	ret model.DiagnoseInterface, err error) {

	endpointInterface, err := s.EndpointInterfaceRepo.Get(tenantId, uint(endpointInterfaceId))
	if err != nil {
		return
	}

	// convert or clone a debug interface obj
	debugData, err := s.DebugInterfaceService.GetDebugDataFromEndpointInterface(tenantId, uint(endpointInterfaceId))
	debugData.UsedBy = "" // mark src usedBy for pre/post-condition loading, empty for no pre/post conditions

	debugData.EndpointInterfaceId = uint(endpointInterfaceId)
	//debugData.ServeId = parent.ServeId

	if debugData.ServerId == 0 {
		server, _ := s.ServeServerRepo.GetDefaultByServe(tenantId, debugData.ServeId)
		debugData.ServerId = server.ID
	}

	server, _ := s.ServeServerRepo.Get(tenantId, debugData.ServerId)
	debugData.BaseUrl = "" // no need to bind to env in debug page
	debugData.Url = _httpUtils.CombineUrls(server.Url, debugData.Url)

	debugData.UsedBy = consts.DiagnoseDebug
	srcDebugInterfaceId := debugData.DebugInterfaceId
	debugInterface, err := s.DebugInterfaceService.SaveAs(tenantId, debugData, srcDebugInterfaceId, "")

	// save test interface
	diagnoseInterface := model.DiagnoseInterface{
		Title:            endpointInterface.Name,
		Type:             serverConsts.DiagnoseInterfaceTypeInterface,
		Ordr:             s.DiagnoseInterfaceRepo.GetMaxOrder(tenantId, parent.ID),
		Method:           debugData.Method,
		DebugInterfaceId: debugInterface.ID,
		ParentId:         parent.ID,
		ServeId:          parent.ServeId,
		ProjectId:        parent.ProjectId,
		CreatedBy:        createBy,
	}
	s.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)

	// update diagnose_interface_id
	values := map[string]interface{}{
		"diagnose_interface_id": diagnoseInterface.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, debugInterface.ID, values)

	ret = diagnoseInterface

	return
}

func (s *DiagnoseInterfaceService) ImportCurl(tenantId consts.TenantId, req serverDomain.DiagnoseCurlImportReq) (ret model.DiagnoseInterface, err error) {
	parent, _ := s.DiagnoseInterfaceRepo.Get(tenantId, req.TargetId)
	if parent.Type != serverConsts.DiagnoseInterfaceTypeDir {
		parent, _ = s.DiagnoseInterfaceRepo.Get(tenantId, parent.ParentId)
	}

	curlObj := curlHelper.Parse(req.Content)
	wf := curlObj.CreateTemporary(curlObj.CreateSession())

	url := fmt.Sprintf("%s://%s%s", curlObj.ParsedURL.Scheme,
		curlObj.ParsedURL.Host, curlObj.ParsedURL.Path)
	//title := fmt.Sprintf("%s %s", url, curlObj.Method)
	queryParams := s.getQueryParams(curlObj.ParsedURL.Query())
	headers := s.getHeaders(wf.Header)
	cookies := s.getCookies(wf.Cookies)

	server, _ := s.ServeServerRepo.GetDefaultByServe(tenantId, parent.ServeId)
	bodyType := ""
	contentType := strings.Split(curlObj.ContentType, ";")
	if len(contentType) > 1 {
		bodyType = contentType[0]
	}

	debugData := domain.DebugData{
		Name:    url,
		BaseUrl: "",
		BaseRequest: domain.BaseRequest{
			Method:      s.getMethod(bodyType, curlObj.Method),
			QueryParams: &queryParams,
			Headers:     &headers,
			Cookies:     &cookies,
			Body:        wf.Body.String(),
			BodyType:    consts.HttpContentType(bodyType),
			Url:         url,
		},
		ServeId:   parent.ServeId,
		ServerId:  server.ID,
		ProjectId: parent.ProjectId,

		UsedBy: consts.DiagnoseDebug,
	}

	debugInterface, err := s.DebugInterfaceService.SaveAs(tenantId, debugData, 0, "")

	// save test interface
	diagnoseInterface := model.DiagnoseInterface{
		Title:            url,
		Type:             serverConsts.DiagnoseInterfaceTypeInterface,
		Ordr:             s.DiagnoseInterfaceRepo.GetMaxOrder(tenantId, parent.ID),
		Method:           debugData.Method,
		DebugInterfaceId: debugInterface.ID,
		ParentId:         parent.ID,
		ServeId:          parent.ServeId,
		ProjectId:        parent.ProjectId,
		CreatedBy:        req.CreateBy,
	}
	s.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)

	// update diagnose_interface_id
	values := map[string]interface{}{
		"diagnose_interface_id": diagnoseInterface.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, debugInterface.ID, values)

	ret = diagnoseInterface

	return
}

func (s *DiagnoseInterfaceService) getQueryParams(params url.Values) (ret []domain.Param) {
	for key, arr := range params {
		for _, item := range arr {
			ret = append(ret, domain.Param{
				Name:    key,
				Value:   item,
				ParamIn: consts.ParamInQuery,
			})
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Name < ret[j].Name
	})
	return
}

func (s *DiagnoseInterfaceService) getHeaders(header http.Header) (ret []domain.Header) {
	for key, arr := range header {
		for _, item := range arr {
			ret = append(ret, domain.Header{
				Name:  key,
				Value: item,
			})
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Name < ret[j].Name
	})
	return
}

func (s *DiagnoseInterfaceService) getCookies(cookies map[string]*http.Cookie) (ret []domain.ExecCookie) {
	for _, item := range cookies {
		ret = append(ret, domain.ExecCookie{
			Name:       item.Name,
			Value:      item.Value,
			ExpireTime: &item.Expires,
			Domain:     item.Domain,
		})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Name < ret[j].Name
	})
	return
}

func (s *DiagnoseInterfaceService) getMethod(contentType, method string) (ret consts.HttpMethod) {

	if method == "" && contentType == "application/json" {
		method = "POST"
	}

	return consts.HttpMethod(method)
}
