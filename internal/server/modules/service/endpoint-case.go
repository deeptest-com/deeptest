package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"log"
	"strconv"
)

type EndpointCaseService struct {
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	PreConditionRepo      *repo.PreConditionRepo      `inject:""`
	PostConditionRepo     *repo.PostConditionRepo     `inject:""`
	CategoryRepo          *repo.CategoryRepo          `inject:""`

	EndpointService       *EndpointService       `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`
}

func (s *EndpointCaseService) List(endpointId uint) (ret []model.EndpointCase, err error) {
	ret, err = s.EndpointCaseRepo.List(endpointId)

	return
}

func (s *EndpointCaseService) Get(id int) (ret model.EndpointCase, err error) {
	ret, err = s.EndpointCaseRepo.Get(uint(id))
	// its debug data will load in webpage

	return
}

func (s *EndpointCaseService) Create(req serverDomain.EndpointCaseSaveReq) (casePo model.EndpointCase, err error) {
	debugInterfaceId, endpointInterfaceId := s.EndpointInterfaceRepo.GetByMethod(req.EndpointId, consts.HttpMethod(req.Method))
	if debugInterfaceId > 0 {
		req.DebugData, err = s.DebugInterfaceService.GetDebugDataFromDebugInterface(debugInterfaceId)
	} else if endpointInterfaceId > 0 {
		req.DebugData, err = s.DebugInterfaceService.ConvertDebugDataFromEndpointInterface(endpointInterfaceId)
	}

	casePo, err = s.SaveFromDebugInterface(req)

	if casePo.DebugInterfaceId > 0 {
		values := map[string]interface{}{
			"case_interface_id": casePo.ID,
		}
		err = s.DebugInterfaceRepo.UpdateDebugInfo(casePo.DebugInterfaceId, values)
	}

	return
}

func (s *EndpointCaseService) Copy(id int, userId uint, userName string) (po model.EndpointCase, err error) {
	endpointCase, _ := s.EndpointCaseRepo.Get(uint(id))
	debugData, _ := s.DebugInterfaceService.GetDebugDataFromDebugInterface(endpointCase.DebugInterfaceId)

	req := serverDomain.EndpointCaseSaveReq{
		Name:       "copy-" + endpointCase.Name,
		EndpointId: endpointCase.EndpointId,
		ServeId:    endpointCase.ServeId,
		ProjectId:  endpointCase.ProjectId,

		CreateUserId:   userId,
		CreateUserName: userName,

		DebugData: debugData,
	}

	s.CopyValueFromRequest(&po, req)

	endpoint, err := s.EndpointRepo.Get(req.EndpointId)

	// create new DebugInterface
	url := req.DebugData.Url
	if url == "" {
		url = endpoint.Path
	}

	debugInterface := model.DebugInterface{}

	s.DebugInterfaceService.CopyValueFromRequest(&debugInterface, req.DebugData)
	debugInterface.Name = req.Name
	debugInterface.Url = url

	err = s.DebugInterfaceRepo.Save(&debugInterface)

	// clone conditions
	s.PreConditionRepo.CloneAll(req.DebugData.DebugInterfaceId, 0, debugInterface.ID)
	s.PostConditionRepo.CloneAll(req.DebugData.DebugInterfaceId, 0, debugInterface.ID)

	// save case
	po.ProjectId = endpoint.ProjectId
	po.ServeId = endpoint.ServeId
	po.DebugInterfaceId = debugInterface.ID
	err = s.EndpointCaseRepo.Save(&po)

	if po.DebugInterfaceId > 0 {
		values := map[string]interface{}{
			"case_interface_id": po.ID,
		}
		err = s.DebugInterfaceRepo.UpdateDebugInfo(po.DebugInterfaceId, values)
	}

	return
}

func (s *EndpointCaseService) SaveFromDebugInterface(req serverDomain.EndpointCaseSaveReq) (po model.EndpointCase, err error) {
	if req.Method != req.DebugData.Method {
		debugInterfaceId, endpointInterfaceId := s.EndpointInterfaceRepo.GetByMethod(req.EndpointId, req.Method)
		info := domain.DebugInfo{DebugInterfaceId: debugInterfaceId, EndpointInterfaceId: endpointInterfaceId}

		req.DebugData, err = s.DebugInterfaceService.Load(info)
	}

	// save debug data
	req.DebugData.UsedBy = consts.CaseDebug
	srcDebugInterfaceId := req.DebugData.DebugInterfaceId
	debugInterface, err := s.DebugInterfaceService.SaveAs(req.DebugData, srcDebugInterfaceId)

	// save case
	s.CopyValueFromRequest(&po, req)

	if po.EndpointId == 0 {
		po.EndpointId = req.EndpointId
	}
	endpoint, err := s.EndpointRepo.Get(po.EndpointId)
	po.ProjectId = endpoint.ProjectId
	po.ServeId = endpoint.ServeId

	po.DebugInterfaceId = debugInterface.ID
	po.ID = 0
	err = s.EndpointCaseRepo.Save(&po)

	if po.DebugInterfaceId > 0 {
		values := map[string]interface{}{
			"case_interface_id": po.ID,
		}
		err = s.DebugInterfaceRepo.UpdateDebugInfo(po.DebugInterfaceId, values)
	}

	if err != nil {
		return
	}

	return
}

func (s *EndpointCaseService) GenerateFromSpec(req serverDomain.EndpointCaseGenerateReq) (err error) {
	endpointInterfaceId := req.EndpointInterfaceId
	if endpointInterfaceId == 0 {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(uint(endpointInterfaceId))
	endpoint, err := s.EndpointRepo.GetWithInterface(endpointInterface.EndpointId, "v0.1.0")

	// get spec
	spec := s.EndpointService.Yaml(endpoint)
	doc3 := spec.(*openapi3.T)
	apiPathItem, _ := cases.GetApiPathItem(doc3)

	for _, interf := range endpoint.Interfaces {
		basicDebugData, err1 := s.getBaseRequest(interf)
		if err1 != nil {
			continue
		}
		log.Println(basicDebugData)

		apiOperation, err1 := cases.GetApiOperation(interf.Method, apiPathItem)
		if err1 != nil {
			continue
		}
		log.Println(apiOperation)

		alternativeCases, err1 := cases.GenerateAlternativeCase(basicDebugData, apiOperation)
		if err1 != nil {
			continue
		}
		log.Println(alternativeCases)
	}

	return
}

func (s *EndpointCaseService) getBaseRequest(endpointInterface model.EndpointInterface) (debugData domain.DebugData, err error) {
	info := domain.DebugInfo{
		DebugInterfaceId:    endpointInterface.DebugInterfaceId,
		EndpointInterfaceId: endpointInterface.ID,
	}
	debugData, err = s.DebugInterfaceService.Load(info)

	return
}

func (s *EndpointCaseService) UpdateName(req serverDomain.EndpointCaseSaveReq) (err error) {
	err = s.EndpointCaseRepo.UpdateName(req)

	return
}

func (s *EndpointCaseService) Remove(id uint) (err error) {
	err = s.EndpointCaseRepo.Remove(id)
	return
}

func (s *EndpointCaseService) CopyValueFromRequest(po *model.EndpointCase, req serverDomain.EndpointCaseSaveReq) {
	copier.CopyWithOption(po, req, copier.Option{
		DeepCopy: true,
	})
}

func (s *EndpointCaseService) EndpointCaseToTo(po *serverDomain.InterfaceCase) (to *serverDomain.EndpointCaseTree) {
	to = &serverDomain.EndpointCaseTree{
		//Id:               uuid.NewV4(),
		Id:               "case_" + strconv.FormatInt(int64(po.ID), 10),
		Key:              int64(po.ID),
		Name:             po.Name,
		Method:           po.Method,
		Desc:             po.Desc,
		Type:             serverConsts.EndpointCaseTreeTypeCase,
		IsDir:            false,
		EndpointId:       po.EndpointId,
		DebugInterfaceId: po.DebugInterfaceId,
		CaseInterfaceId:  po.ID,
		ProjectId:        po.ProjectId,
		ServeId:          po.ServeId,
	}

	return
}

func (s *EndpointCaseService) LoadTree(projectId, serveId uint) (ret []*serverDomain.EndpointCaseTree, err error) {
	list, err := s.EndpointCaseRepo.GetCategoryEndpointCase(projectId, serveId)
	if err != nil {
		return
	}

	entityMap := make(map[string]*serverDomain.EndpointCaseTree)
	for _, v := range list {
		entityMap[v.CaseUniqueId] = &serverDomain.EndpointCaseTree{
			Id:               v.CaseUniqueId,
			Key:              int64(v.CaseId),
			Name:             v.CaseName,
			Method:           v.Method,
			Desc:             v.CaseDesc,
			Type:             serverConsts.EndpointCaseTreeTypeCase,
			IsDir:            false,
			EndpointId:       v.EndpointId,
			DebugInterfaceId: v.CaseDebugInterfaceId,
			CaseInterfaceId:  v.CaseId,
			ParentId:         v.EndpointUniqueId,
			ProjectId:        v.ProjectId,
			ServeId:          v.ServeId,
			Slots:            iris.Map{"icon": "icon"},
		}

		entityMap[v.EndpointUniqueId] = &serverDomain.EndpointCaseTree{
			Id:         v.EndpointUniqueId,
			Key:        int64(v.EndpointId),
			Name:       v.EndpointTitle,
			Desc:       v.EndpointDescription,
			Type:       serverConsts.EndpointCaseTreeTypeEndpoint,
			IsDir:      true,
			CategoryId: v.CategoryId,
			ParentId:   "category_" + strconv.FormatInt(v.CategoryId, 10),
			ProjectId:  v.ProjectId,
			ServeId:    v.ServeId,
			Slots:      iris.Map{"icon": "icon"},
		}
	}

	categories, err := s.CategoryRepo.ListByProject(serverConsts.EndpointCategory, projectId, 0)
	for _, v := range categories {
		uniqueId := "category_" + strconv.FormatInt(int64(v.ID), 10)
		entityMap[uniqueId] = &serverDomain.EndpointCaseTree{
			Id:        uniqueId,
			Key:       int64(v.ID),
			Name:      v.Name,
			Desc:      v.Desc,
			Type:      serverConsts.EndpointCaseTreeTypeDir,
			IsDir:     true,
			ParentId:  "category_" + strconv.FormatInt(int64(v.ParentId), 10),
			ProjectId: v.ProjectId,
			ServeId:   v.ServeId,
			Slots:     iris.Map{"icon": "icon"},
		}
	}
	//未分类
	entityMap["category_-1"] = &serverDomain.EndpointCaseTree{
		Id:        "category_-1",
		Key:       -1,
		Name:      "未分类",
		Type:      serverConsts.EndpointCaseTreeTypeDir,
		IsDir:     true,
		ParentId:  "category_" + strconv.FormatInt(int64(categories[0].ID), 10),
		ProjectId: projectId,
		ServeId:   serveId,
		Slots:     iris.Map{"icon": "icon"},
	}

	ret = s.MapToTree(entityMap, "category_"+strconv.FormatInt(int64(categories[0].ID), 10))
	s.GetNodeCaseNum(ret)
	return
}

func (s *EndpointCaseService) MapToTree(mapData map[string]*serverDomain.EndpointCaseTree, parentId string) (res []*serverDomain.EndpointCaseTree) {
	for k, v := range mapData {
		if v.ParentId == parentId {
			v.Children = s.MapToTree(mapData, k)
			res = append(res, v)
		}
	}
	return
}

func (s *EndpointCaseService) GetNodeCaseNum(res []*serverDomain.EndpointCaseTree) (num int64) {
	for _, v := range res {
		if v.Type == serverConsts.EndpointCaseTreeTypeCase {
			num = 1
			v.Count = 0
		} else if v.Type == serverConsts.EndpointCaseTreeTypeEndpoint {
			num += int64(len(v.Children))
			v.Count = int64(len(v.Children))
		} else {
			num = s.GetNodeCaseNum(v.Children)
			v.Count += num
		}
	}
	return
}
