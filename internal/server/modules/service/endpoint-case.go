package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"strconv"
)

type EndpointCaseService struct {
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ConditionRepo         *repo.ConditionRepo         `inject:""`
	CategoryRepo          *repo.CategoryRepo          `inject:""`

	EndpointService       *EndpointService       `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`
}

func (s *EndpointCaseService) Paginate(tenantId consts.TenantId, req serverDomain.EndpointCaseReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointCaseRepo.Paginate(tenantId, req)

	return
}

func (s *EndpointCaseService) List(tenantId consts.TenantId, endpointId uint) (ret []model.EndpointCase, err error) {
	ret, err = s.EndpointCaseRepo.List(tenantId, endpointId)

	return
}

func (s *EndpointCaseService) Get(tenantId consts.TenantId, id uint) (ret model.EndpointCase, err error) {
	ret, err = s.EndpointCaseRepo.Get(tenantId, id)
	// its debug data will load in webpage

	return
}

func (s *EndpointCaseService) Create(tenantId consts.TenantId, req serverDomain.EndpointCaseSaveReq) (casePo model.EndpointCase, err error) {
	_, endpointInterfaceId := s.EndpointInterfaceRepo.GetByMethod(tenantId, req.EndpointId, req.Method)
	req.DebugData, err = s.DebugInterfaceService.ConvertDebugDataFromEndpointInterface(tenantId, endpointInterfaceId)

	casePo, err = s.SaveFromDebugInterface(tenantId, req)

	if casePo.DebugInterfaceId > 0 {
		values := map[string]interface{}{
			"case_interface_id": casePo.ID,
		}
		err = s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, casePo.DebugInterfaceId, values)
	}

	return
}

func (s *EndpointCaseService) Copy(tenantId consts.TenantId, id int, newNamePrefix string, newEndpointId, baseCaseId, userId uint, userName,
	forAlternativeCase string) (po model.EndpointCase, err error) {

	endpointCase, _ := s.EndpointCaseRepo.Get(tenantId, uint(id))
	debugData, _ := s.DebugInterfaceService.GetDebugDataFromDebugInterface(tenantId, endpointCase.DebugInterfaceId)
	debugData.UsedBy = consts.CaseDebug

	if newEndpointId != 0 {
		endpointCase.EndpointId = newEndpointId
	}

	if baseCaseId != 0 {
		endpointCase.BaseCase = baseCaseId
	}

	if newNamePrefix == "" {
		newNamePrefix = "copy-" + endpointCase.Name
	}

	caseType := endpointCase.CaseType
	if caseType == consts.CaseBenchmark && newEndpointId == 0 { //复制接口定义导致的复制用例不改变，用例类型
		caseType = consts.CaseDefault
	}
	req := serverDomain.EndpointCaseSaveReq{
		Name:       newNamePrefix,
		EndpointId: endpointCase.EndpointId,
		ServeId:    endpointCase.ServeId,
		ProjectId:  endpointCase.ProjectId,

		CreateUserId:   userId,
		CreateUserName: userName,

		Method:    endpointCase.Method,
		DebugData: debugData,
		CaseType:  caseType,
		BaseCase:  endpointCase.BaseCase,
	}

	s.CopyValueFromRequest(tenantId, &po, req)

	endpoint, err := s.EndpointRepo.Get(tenantId, req.EndpointId)

	// create new DebugInterface
	url := req.DebugData.Url
	if url == "" {
		url = endpoint.Path
	}

	debugInterface := model.DebugInterface{}

	s.DebugInterfaceService.CopyValueFromRequest(tenantId, &debugInterface, req.DebugData)
	debugInterface.Name = req.Name
	debugInterface.Url = url

	err = s.DebugInterfaceRepo.Save(tenantId, &debugInterface)

	// clone conditions
	s.ConditionRepo.CloneAll(tenantId, req.DebugData.DebugInterfaceId, 0, debugInterface.ID, debugData.UsedBy, debugData.UsedBy, forAlternativeCase)

	// save case
	po.ProjectId = endpoint.ProjectId
	po.ServeId = endpoint.ServeId
	po.DebugInterfaceId = debugInterface.ID
	err = s.EndpointCaseRepo.Save(tenantId, &po)

	if po.DebugInterfaceId > 0 {
		values := map[string]interface{}{
			"case_interface_id": po.ID,
		}
		err = s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, po.DebugInterfaceId, values)
	}

	return
}

func (s *EndpointCaseService) SaveFromDebugInterface(tenantId consts.TenantId, req serverDomain.EndpointCaseSaveReq) (po model.EndpointCase, err error) {
	if req.Method != req.DebugData.Method {
		debugInterfaceId, endpointInterfaceId := s.EndpointInterfaceRepo.GetByMethod(tenantId, req.EndpointId, req.Method)
		info := domain.DebugInfo{DebugInterfaceId: debugInterfaceId, EndpointInterfaceId: endpointInterfaceId}

		req.DebugData, err = s.DebugInterfaceService.Load(tenantId, info)
	}

	// save debug data
	srcDebugUsedBy := req.DebugData.UsedBy
	req.DebugData.UsedBy = consts.CaseDebug
	srcDebugInterfaceId := req.DebugData.DebugInterfaceId
	debugInterface, err := s.DebugInterfaceService.SaveAs(tenantId, req.DebugData, srcDebugInterfaceId, srcDebugUsedBy)

	// save case
	s.CopyValueFromRequest(tenantId, &po, req)

	if po.EndpointId == 0 {
		po.EndpointId = req.EndpointId
	}
	endpoint, err := s.EndpointRepo.Get(tenantId, po.EndpointId)
	po.ProjectId = endpoint.ProjectId
	po.ServeId = endpoint.ServeId

	po.Method = debugInterface.Method
	po.DebugInterfaceId = debugInterface.ID
	po.ID = 0
	err = s.EndpointCaseRepo.Save(tenantId, &po)

	if po.DebugInterfaceId > 0 {
		values := map[string]interface{}{
			"case_interface_id": po.ID,
		}
		err = s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, po.DebugInterfaceId, values)
	}

	if err != nil {
		return
	}

	return
}

func (s *EndpointCaseService) UpdateName(tenantId consts.TenantId, req serverDomain.EndpointCaseSaveReq) (err error) {
	err = s.EndpointCaseRepo.UpdateName(tenantId, req)

	return
}

func (s *EndpointCaseService) Remove(tenantId consts.TenantId, id uint) (err error) {
	err = s.EndpointCaseRepo.Remove(tenantId, id)
	return
}

func (s *EndpointCaseService) CopyValueFromRequest(tenantId consts.TenantId, po *model.EndpointCase, req serverDomain.EndpointCaseSaveReq) {
	copier.CopyWithOption(po, req, copier.Option{
		DeepCopy: true,
	})
}

func (s *EndpointCaseService) EndpointCaseToTo(tenantId consts.TenantId, po *serverDomain.InterfaceCase) (to *serverDomain.EndpointCaseTree) {
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

func (s *EndpointCaseService) LoadTree(tenantId consts.TenantId, projectId uint, serveIds consts.Integers) (ret []*serverDomain.EndpointCaseTree, err error) {
	list, err := s.EndpointCaseRepo.GetCategoryEndpointCase(tenantId, projectId, serveIds)
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

	categories, err := s.CategoryRepo.ListByProject(tenantId, serverConsts.EndpointCategory, projectId)
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
		Slots:     iris.Map{"icon": "icon"},
	}

	ret = s.MapToTree(tenantId, entityMap, "category_"+strconv.FormatInt(int64(categories[0].ID), 10))
	return s.GetNodeCaseNum(tenantId, ret), nil

}

func (s *EndpointCaseService) MapToTree(tenantId consts.TenantId, mapData map[string]*serverDomain.EndpointCaseTree, parentId string) (res []*serverDomain.EndpointCaseTree) {
	for k, v := range mapData {
		if v.ParentId == parentId {
			v.Children = s.MapToTree(tenantId, mapData, k)
			res = append(res, v)
		}
	}
	return
}

func (s *EndpointCaseService) GetNodeCaseNum(tenantId consts.TenantId, req []*serverDomain.EndpointCaseTree) (ret []*serverDomain.EndpointCaseTree) {

	root := &serverDomain.EndpointCaseTree{Children: req}
	s.GetNodeCaseNumNew(tenantId, root)
	return root.Children

}

func (s *EndpointCaseService) GetNodeCaseNumNew(tenantId consts.TenantId, node *serverDomain.EndpointCaseTree) (num int64) {
	if node.Type == serverConsts.EndpointCaseTreeTypeCase {
		return 1
	}

	var children []*serverDomain.EndpointCaseTree
	for _, child := range node.Children {
		n := s.GetNodeCaseNumNew(tenantId, child)
		if n != 0 {
			children = append(children, child)
		}
		node.Count += n
	}
	node.Children = children

	return node.Count

}

func (s *EndpointCaseService) ListByCaseType(tenantId consts.TenantId, endpointId uint, caseType consts.CaseType) (ret []model.EndpointCase, err error) {
	ret, err = s.EndpointCaseRepo.ListByCaseType(tenantId, endpointId, []consts.CaseType{caseType})

	return
}

func (s *EndpointCaseService) CopyChildrenCases(tenantId consts.TenantId, caseId, newCaseId, endpointId, userId uint, username string) (err error) {
	childrenCases, err := s.EndpointCaseRepo.ListByCaseTypeAndBaseCase(tenantId, consts.CaseAlternative, caseId)
	if err != nil {
		return err
	}

	for _, item := range childrenCases {

		_, err = s.Copy(tenantId, int(item.ID), item.Name, endpointId, newCaseId, userId, username, "false")
		if err != nil {
			return err
		}
	}

	return
}
