package service

import (
	"encoding/json"
	"errors"
	"fmt"
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	curlHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/gcurl"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	repo "github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"log"
	"strings"
)

type ScenarioNodeService struct {
	ScenarioNodeRepo         *repo.ScenarioNodeRepo      `inject:""`
	ScenarioProcessorRepo    *repo.ScenarioProcessorRepo `inject:""`
	ScenarioProcessorService *ScenarioProcessorService   `inject:""`
	ScenarioRepo             *repo.ScenarioRepo          `inject:""`
	DebugInterfaceRepo       *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	ExtractorRepo            *repo.ExtractorRepo         `inject:""`
	CheckpointRepo           *repo.CheckpointRepo        `inject:""`
	ServeServerRepo          *repo.ServeServerRepo       `inject:""`

	DebugInterfaceService    *DebugInterfaceService    `inject:""`
	DiagnoseInterfaceService *DiagnoseInterfaceService `inject:""`
}

func (s *ScenarioNodeService) GetTree(tenantId consts.TenantId, scenario model.Scenario, withDetail bool) (root *agentExec.Processor, err error) {
	pos, err := s.ScenarioNodeRepo.ListByScenario(tenantId, scenario.ID)
	if err != nil {
		return
	}

	tos := s.ToTos(tenantId, pos, withDetail)

	root = tos[0]
	root.Name = scenario.Name
	root.Slots = iris.Map{"icon": "icon"}

	s.ScenarioNodeRepo.MakeTree(tenantId, tos[1:], root)

	root.ScenarioId = scenario.ID
	root.ProjectId = scenario.ProjectId

	return
}

func (s *ScenarioNodeService) ToTos(tenantId consts.TenantId, pos []*model.Processor, withDetail bool) (tos []*agentExec.Processor) {
	for _, po := range pos {
		to := agentExec.Processor{
			ProcessorBase: agentExec.ProcessorBase{
				ScenarioId: po.ScenarioId,
			},
		}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})
		to.Disable = po.Disabled
		to.Comments = po.Comments

		if withDetail {
			entity, _ := s.ScenarioProcessorService.GetEntityTo(tenantId, &to)
			to.EntityRaw, _ = json.Marshal(entity)

			log.Println("")
		}

		// just to avoid json marshal error for IProcessorEntity
		to.Entity = agentExec.ProcessorGroup{}

		tos = append(tos, &to)
	}

	return
}

func (s *ScenarioNodeService) AddProcessor(tenantId consts.TenantId, req serverDomain.ScenarioAddScenarioReq, source string) (ret model.Processor, err *_domain.BizErr) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, uint(req.TargetProcessorId))
	if targetProcessor.ID == 0 {
		return
	}

	ret = model.Processor{
		Name:                  strings.TrimSpace(req.Name),
		EntityCategory:        req.ProcessorCategory,
		EntityType:            req.ProcessorType,
		ProcessorInterfaceSrc: req.ProcessorInterfaceSrc,
		EndpointInterfaceId:   targetProcessor.EndpointInterfaceId,
		ScenarioId:            targetProcessor.ScenarioId,
		ProjectId:             req.ProjectId,
		CreatedBy:             req.CreateBy,
		BaseModel:             model.BaseModel{Disabled: targetProcessor.Disabled},
		Comments:              req.Comments,
		Method:                req.Method,
	}

	if req.Mode == "child" {
		ret.ParentId = targetProcessor.ID
	} else if req.Mode == "brother" {
		ret.ParentId = targetProcessor.ParentId
	} else if req.Mode == "parent" && req.TargetProcessorCategory == consts.ProcessorInterface {
		ret.ParentId = targetProcessor.ParentId
	} else if req.Mode == "siblings" {
		ret.ParentId = targetProcessor.ParentId
	}

	//相邻节点需要插空，其他节点后移
	if req.Mode == "siblings" {
		s.ScenarioNodeRepo.MoveMaxOrder(tenantId, ret.ParentId, uint(targetProcessor.Ordr), 1)
		ret.Ordr = targetProcessor.Ordr + 1
	} else {
		ret.Ordr = s.ScenarioNodeRepo.GetMaxOrder(tenantId, ret.ParentId)
	}

	s.ScenarioNodeRepo.Save(tenantId, &ret)

	if req.Mode == "parent" { // move interface to new folder
		targetProcessor.ParentId = ret.ID
		s.ScenarioNodeRepo.Save(tenantId, &targetProcessor)
	}

	if source == "copy" {
		if ret.EntityType == consts.ProcessorInterfaceDefault {
			s.CopyInterfaceEntity(tenantId, req.SrcProcessorId, ret.ID)
		} else {
			s.ScenarioProcessorRepo.CopyEntity(tenantId, req.SrcProcessorId, ret.ID)
		}
	} else {
		if ret.EntityType == consts.ProcessorInterfaceDefault { // create debug interface
			debugInterfaceId, _ := s.DebugInterfaceService.CreateDefault(tenantId, ret.ProcessorInterfaceSrc, req.ProjectId)
			s.ScenarioProcessorRepo.UpdateInterfaceId(tenantId, ret.ID, debugInterfaceId)
		} else if ret.EntityType == consts.ProcessorLogicElse { // create default entity
			entity := model.ProcessorLogic{
				ProcessorEntityBase: agentExec.ProcessorEntityBase{
					ProcessorID:       ret.ID,
					ProcessorCategory: ret.EntityCategory,
					ProcessorType:     ret.EntityType,
				},
			}
			_ = s.ScenarioProcessorRepo.SaveLogic(tenantId, &entity)

		}
	}

	return
}

func (s *ScenarioNodeService) CopyInterfaceEntity(tenantId consts.TenantId, srcProcessorId, distProcessorId uint) (err error) {
	srcProcessor, err := s.ScenarioNodeRepo.Get(tenantId, srcProcessorId)
	if err != nil {
		return
	}

	debugData, err := s.DebugInterfaceService.GetDebugDataFromDebugInterface(tenantId, srcProcessor.EntityId)
	if err != nil {
		return
	}

	debugData.ScenarioProcessorId = distProcessorId
	debugData.UsedBy = consts.ScenarioDebug
	debugInterface, err := s.DebugInterfaceService.SaveAs(tenantId, debugData, debugData.DebugInterfaceId, debugData.UsedBy)
	if err != nil {
		return
	}

	err = s.ScenarioProcessorRepo.UpdateInterfaceId(tenantId, distProcessorId, debugInterface.ID)
	if err != nil {
		return err
	}

	return
}
func (s *ScenarioNodeService) AddInterfacesFromDiagnose(tenantId consts.TenantId, req serverDomain.ScenarioAddInterfacesFromTreeReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(tenantId, targetProcessor.ParentId)
	}

	for _, interfaceNode := range req.SelectedNodes {
		ret, _ = s.createDirOrInterfaceFromDiagnose(tenantId, &interfaceNode, targetProcessor, 0)
	}

	return
}

func (s *ScenarioNodeService) AddInterfacesFromDefine(tenantId consts.TenantId, req serverDomain.ScenarioAddInterfacesReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(tenantId, targetProcessor.ParentId)
	}

	serveId := uint(0)
	for _, interfaceId := range req.InterfaceIds {
		ret, err = s.createInterfaceFromDefine(tenantId, uint(interfaceId), &serveId, req.CreateBy, targetProcessor, "", 0)
	}

	return
}

func (s *ScenarioNodeService) AddInterfacesFromCase(tenantId consts.TenantId, req serverDomain.ScenarioAddCasesFromTreeReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(tenantId, targetProcessor.ParentId)
	}

	for _, interfaceNode := range req.SelectedNodes {
		ret, _ = s.createDirOrInterfaceFromCase(tenantId, &interfaceNode, targetProcessor, 0)
	}

	return
}

func (s *ScenarioNodeService) createInterfaceFromDefine(tenantId consts.TenantId, endpointInterfaceId uint, serveId *uint,
	createBy uint, parentProcessor model.Processor, name string, order int) (
	ret model.Processor, err error) {

	endpointInterface, err := s.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if err == gorm.ErrRecordNotFound {
		err = errors.New("interface is deleted")
		return
	}

	// get serve id once
	if *serveId == 0 {
		endpoint, _ := s.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)
		*serveId = endpoint.ServeId
	}

	// convert or clone a debug interface obj
	debugData, err := s.DebugInterfaceService.GetDebugDataFromEndpointInterface(tenantId, endpointInterfaceId)
	debugData.UsedBy = consts.ScenarioDebug

	debugData.EndpointInterfaceId = endpointInterfaceId

	debugData.ScenarioProcessorId = 0 // will be update after ScenarioProcessor saved
	debugData.ProcessorInterfaceSrc = consts.InterfaceSrcDefine

	debugData.ServeId = *serveId

	if debugData.ServerId == 0 {
		server, _ := s.ServeServerRepo.GetDefaultByServe(tenantId, debugData.ServeId)
		debugData.ServerId = server.ID
	}

	server, _ := s.ServeServerRepo.Get(tenantId, debugData.ServerId)
	debugData.ServerId = server.ID
	debugData.BaseUrl = server.Url

	srcDebugInterfaceId := debugData.DebugInterfaceId
	debugInterface, err := s.DebugInterfaceService.SaveAs(tenantId, debugData, srcDebugInterfaceId, "")

	if order == 0 {
		order = s.ScenarioNodeRepo.GetMaxOrder(tenantId, parentProcessor.ID)
	}
	processor := model.Processor{
		Name:                endpointInterface.Name,
		Method:              endpointInterface.Method,
		EntityCategory:      consts.ProcessorInterface,
		EntityType:          consts.ProcessorInterfaceDefault,
		EntityId:            debugInterface.ID, // as debugInterfaceId
		EndpointInterfaceId: debugInterface.EndpointInterfaceId,

		//Ordr: s.ScenarioNodeRepo.GetMaxOrder(parentProcessor.ID),
		Ordr:                  order,
		ParentId:              parentProcessor.ID,
		ScenarioId:            parentProcessor.ScenarioId,
		ProjectId:             parentProcessor.ProjectId,
		CreatedBy:             createBy,
		ProcessorInterfaceSrc: consts.InterfaceSrcDefine,
		BaseModel:             model.BaseModel{Disabled: parentProcessor.Disabled},
	}

	s.ScenarioNodeRepo.Save(tenantId, &processor)

	// update to new ScenarioProcessorId
	values := map[string]interface{}{
		"scenario_processor_id": processor.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, debugInterface.ID, values)

	ret = processor

	return
}

func (s *ScenarioNodeService) createDirOrInterfaceFromDiagnose(tenantId consts.TenantId, diagnoseInterfaceNode *serverDomain.DiagnoseInterface, parentProcessor model.Processor, order int) (
	ret model.Processor, err error) {

	debugData, _ := s.DebugInterfaceService.GetDebugDataFromDebugInterface(tenantId, diagnoseInterfaceNode.DebugInterfaceId)
	debugData.UsedBy = consts.ScenarioDebug

	if diagnoseInterfaceNode.IsDir && len(diagnoseInterfaceNode.Children) > 0 { // dir
		for _, child := range diagnoseInterfaceNode.Children {
			ret, _ = s.createDirOrInterfaceFromDiagnose(tenantId, child, parentProcessor, 0)
		}

	} else if !diagnoseInterfaceNode.IsDir { // interface
		if order == 0 {
			order = s.ScenarioNodeRepo.GetMaxOrder(tenantId, parentProcessor.ID)
		}
		processor := model.Processor{
			Name:                diagnoseInterfaceNode.Title,
			Method:              debugData.Method,
			EntityCategory:      consts.ProcessorInterface,
			EntityType:          consts.ProcessorInterfaceDefault,
			EntityId:            0, // as debugInterfaceId
			EndpointInterfaceId: debugData.EndpointInterfaceId,

			//Ordr: s.ScenarioNodeRepo.GetMaxOrder(parentProcessor.ID),
			Ordr: order,

			ParentId:              parentProcessor.ID,
			ScenarioId:            parentProcessor.ScenarioId,
			ProjectId:             parentProcessor.ProjectId,
			CreatedBy:             parentProcessor.CreatedBy,
			ProcessorInterfaceSrc: consts.InterfaceSrcDiagnose,
			BaseModel:             model.BaseModel{Disabled: parentProcessor.Disabled},
		}

		//processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
		s.ScenarioNodeRepo.Save(tenantId, &processor)

		// convert or clone a debug interface obj
		debugData.ScenarioProcessorId = processor.ID
		debugData.ProcessorInterfaceSrc = consts.InterfaceSrcDiagnose

		debugData.ServeId = diagnoseInterfaceNode.ServeId

		debugInterfaceOfDiagnoseInterfaceNode, _ := s.DebugInterfaceRepo.Get(tenantId, diagnoseInterfaceNode.DebugInterfaceId)
		debugData.ServerId = debugInterfaceOfDiagnoseInterfaceNode.ServerId

		debugData.BaseUrl = "" // no need to bind to env in debug page
		debugData.Url = debugInterfaceOfDiagnoseInterfaceNode.Url

		srcDebugInterfaceId := debugData.DebugInterfaceId
		debugInterface, _ := s.DebugInterfaceService.SaveAs(tenantId, debugData, srcDebugInterfaceId, consts.DiagnoseDebug)

		processor.EntityId = debugInterface.ID
		s.ScenarioProcessorRepo.UpdateEntityId(tenantId, processor.ID, processor.EntityId)

		ret = processor
	}

	return
}

func (s *ScenarioNodeService) createDirOrInterfaceFromCase(tenantId consts.TenantId, caseNode *serverDomain.EndpointCaseTree, parentProcessor model.Processor, order int) (
	processor model.Processor, err error) {
	if caseNode.IsDir && len(caseNode.Children) > 0 { // dir
		/*processor = model.Processor{
			Name:           caseNode.Name,
			ScenarioId:     parentProcessor.ScenarioId,
			EntityCategory: consts.ProcessorGroup,
			EntityType:     consts.ProcessorGroupDefault,
			ParentId:       parentProcessor.ID,
			ProjectId:      parentProcessor.ProjectId,
		}
		processor.Ordr = s.ScenarioNodeRepo.GetMaxOrder(processor.ParentId)
		s.ScenarioNodeRepo.CreateExpression(&processor)*/

		for _, child := range caseNode.Children {
			processor, _ = s.createDirOrInterfaceFromCase(tenantId, child, parentProcessor, 0)
		}
	} else if !caseNode.IsDir { // interface
		debugData, _ := s.DebugInterfaceService.GetDebugDataFromDebugInterface(tenantId, caseNode.DebugInterfaceId)
		debugData.UsedBy = consts.ScenarioDebug

		if order == 0 {
			order = s.ScenarioNodeRepo.GetMaxOrder(tenantId, parentProcessor.ID)
		}
		processor = model.Processor{
			Name:                  caseNode.Name,
			Method:                debugData.Method,
			EntityCategory:        consts.ProcessorInterface,
			EntityType:            consts.ProcessorInterfaceDefault,
			EntityId:              0, // as debugInterfaceId
			EndpointInterfaceId:   debugData.EndpointInterfaceId,
			Ordr:                  order,
			ParentId:              parentProcessor.ID,
			ScenarioId:            parentProcessor.ScenarioId,
			ProjectId:             parentProcessor.ProjectId,
			CreatedBy:             parentProcessor.CreatedBy,
			ProcessorInterfaceSrc: consts.InterfaceSrcCase,
			BaseModel:             model.BaseModel{Disabled: parentProcessor.Disabled},
		}

		s.ScenarioNodeRepo.Save(tenantId, &processor)

		debugData.ScenarioProcessorId = processor.ID
		debugData.ProcessorInterfaceSrc = consts.InterfaceSrcCase

		debugInterfaceOfCaseNode, _ := s.DebugInterfaceRepo.GetByCaseInterfaceId(tenantId, uint(caseNode.Key))
		debugData.Body = debugInterfaceOfCaseNode.Body
		debugData.ServerId = debugInterfaceOfCaseNode.ServerId

		debugData.BaseUrl = "" // no need to bind to env in debug page
		debugData.Url = debugInterfaceOfCaseNode.Url

		srcDebugInterfaceId := debugData.DebugInterfaceId
		debugInterface, _ := s.DebugInterfaceService.SaveAs(tenantId, debugData, srcDebugInterfaceId, consts.CaseDebug)

		processor.EntityId = debugInterface.ID
		s.ScenarioProcessorRepo.UpdateEntityId(tenantId, processor.ID, processor.EntityId)

	}

	return
}

func (s *ScenarioNodeService) UpdateName(tenantId consts.TenantId, req serverDomain.ScenarioNodeReq) (err error) {
	err = s.ScenarioNodeRepo.UpdateName(tenantId, req.Id, req.Name)
	return
}

func (s *ScenarioNodeService) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = s.deleteScenarioNodeAndChildren(tenantId, id)

	return
}

func (s *ScenarioNodeService) DisableOrNot(tenantId consts.TenantId, id uint) (err error) {
	err = s.disableScenarioNodeAndChildren(tenantId, id)

	return
}

func (s *ScenarioNodeService) Move(tenantId consts.TenantId, srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcScenarioNode model.Processor, err error) {
	srcScenarioNode, err = s.ScenarioNodeRepo.Get(tenantId, srcId)

	//获取下一个节点，如果是else准备移动
	srcScenarioNextNode, nextNodeErr := s.ScenarioNodeRepo.GetNextNode(tenantId, srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr, srcScenarioNode.Disabled = s.ScenarioNodeRepo.UpdateOrder(tenantId, pos, targetId)

	err = s.ScenarioNodeRepo.UpdateOrdAndParent(tenantId, srcScenarioNode)

	//判断节点是else一起移动
	if nextNodeErr == nil && srcScenarioNextNode.EntityType == consts.ProcessorLogicElse {
		s.Move(tenantId, srcScenarioNextNode.ID, srcScenarioNode.ID, serverConsts.After, srcScenarioNode.ProjectId)
	}

	return
}

func (s *ScenarioNodeService) deleteScenarioNodeAndChildren(tenantId consts.TenantId, nodeId uint) (err error) {
	srcScenarioNextNode, nextNodeErr := s.ScenarioNodeRepo.GetNextNode(tenantId, nodeId)
	err = s.ScenarioNodeRepo.DeleteWithChildren(tenantId, nodeId)
	if nextNodeErr == nil && srcScenarioNextNode.EntityType == consts.ProcessorLogicElse {
		err = s.ScenarioNodeRepo.DeleteWithChildren(tenantId, srcScenarioNextNode.ID)
	}

	return
}
func (s *ScenarioNodeService) disableScenarioNodeAndChildren(tenantId consts.TenantId, nodeId uint) (err error) {
	node, err := s.ScenarioNodeRepo.Get(tenantId, nodeId)
	srcScenarioNextNode, nextNodeErr := s.ScenarioNodeRepo.GetNextNode(tenantId, nodeId)

	err = s.ScenarioNodeRepo.DisableWithChildren(tenantId, nodeId)

	//如果要禁用if，则else也要禁用
	if !node.Disabled && nextNodeErr == nil && !srcScenarioNextNode.Disabled && srcScenarioNextNode.EntityType == consts.ProcessorLogicElse {
		err = s.ScenarioNodeRepo.DisableWithChildren(tenantId, srcScenarioNextNode.ID)
	}

	//如果启动else，则if则被启用
	if node.Disabled && node.EntityType == consts.ProcessorLogicElse {
		srcScenarioPreNode, preNodeErr := s.ScenarioNodeRepo.GetPreNode(tenantId, nodeId)
		if preNodeErr == nil && srcScenarioPreNode.Disabled {
			s.disableScenarioNodeAndChildren(tenantId, srcScenarioPreNode.ID)
		}
	}

	return
}

func (s *ScenarioNodeService) ListToByScenario(tenantId consts.TenantId, id uint) (ret []*agentExec.Processor, err error) {
	pos, _ := s.ScenarioNodeRepo.ListByScenario(tenantId, id)

	for _, po := range pos {
		to := agentExec.Processor{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		ret = append(ret, &to)
	}

	return
}

func (s *ScenarioNodeService) ImportCurl(tenantId consts.TenantId, req serverDomain.ScenarioCurlImportReq) (ret model.Processor, err error) {
	//req.Content = "curl 'https://thirdpartyapi.nancalcloud.com/api/v1/endpoint/detail?id=1783&ts=1691139567363&currProjectId=2&lang=zh-CN' \\\n  -H 'Accept: application/json, text/plain, */*' \\\n  -H 'Accept-Language: zh-CN,zh;q=0.9' \\\n  -H 'Authorization: Bearer WldJME16Z3lZelV6T0RSbFlXWXdaV0psWWpKbU56TTROR0U1TjJJd01UVS5OV1l3TkRBMU16WmlZekk1WldGa05UTmtOelF4WldRNVlUTXlaamczTURr' \\\n  -H 'Connection: keep-alive' \\\n  -H 'Cookie: td_cookie=3753276482; HWWAFSESID=8f1d573ac34f71a4c5; HWWAFSESTIME=1690431343235' \\\n  -H 'Referer: https://thirdpartyapi.nancalcloud.com/' \\\n  -H 'Sec-Fetch-Dest: empty' \\\n  -H 'Sec-Fetch-Mode: cors' \\\n  -H 'Sec-Fetch-Site: same-origin' \\\n  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36' \\\n  -H 'sec-ch-ua: \"Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"115\", \"Chromium\";v=\"115\"' \\\n  -H 'sec-ch-ua-mobile: ?0' \\\n  -H 'sec-ch-ua-platform: \"Windows\"' \\\n  --compressed"
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(tenantId, targetProcessor.ParentId)
	}

	curlObj := curlHelper.Parse(req.Content)
	wf := curlObj.CreateTemporary(curlObj.CreateSession())

	url := fmt.Sprintf("%s://%s%s", curlObj.ParsedURL.Scheme, curlObj.ParsedURL.Host, curlObj.ParsedURL.Path)
	//title := fmt.Sprintf("%s %s", url, curlObj.Method)
	queryParams := s.DiagnoseInterfaceService.getQueryParams(curlObj.ParsedURL.Query())
	headers := s.DiagnoseInterfaceService.getHeaders(wf.Header)
	cookies := s.DiagnoseInterfaceService.getCookies(wf.Cookies)

	bodyType := ""
	contentType := strings.Split(curlObj.ContentType, ";")
	if len(contentType) > 1 {
		bodyType = contentType[0]
	}

	debugData := domain.DebugData{
		Name:    url,
		BaseUrl: "",
		BaseRequest: domain.BaseRequest{
			Method:      s.DiagnoseInterfaceService.getMethod(bodyType, curlObj.Method),
			QueryParams: &queryParams,
			Headers:     &headers,
			Cookies:     &cookies,
			Body:        wf.Body.String(),
			BodyType:    consts.HttpContentType(bodyType),
			Url:         url,
		},
		//ServeId:   parent.ServeId,
		//ServerId:  server.ID,
		ProjectId:             targetProcessor.ProjectId,
		ProcessorInterfaceSrc: consts.InterfaceSrcCurl,

		UsedBy: consts.ScenarioDebug,
	}

	debugInterface, err := s.DebugInterfaceService.SaveAs(tenantId, debugData, 0, "")

	processor := model.Processor{
		Name:                url,
		Method:              debugData.Method,
		EntityCategory:      consts.ProcessorInterface,
		EntityType:          consts.ProcessorInterfaceDefault,
		EntityId:            debugInterface.ID, // as debugInterfaceId
		EndpointInterfaceId: debugInterface.EndpointInterfaceId,

		Ordr: s.ScenarioNodeRepo.GetMaxOrder(tenantId, targetProcessor.ID),

		ParentId:              targetProcessor.ID,
		ScenarioId:            targetProcessor.ScenarioId,
		ProjectId:             targetProcessor.ProjectId,
		CreatedBy:             req.CreateBy,
		ProcessorInterfaceSrc: consts.InterfaceSrcCurl,
		BaseModel:             model.BaseModel{Disabled: targetProcessor.Disabled},
	}

	err = s.ScenarioNodeRepo.Save(tenantId, &processor)
	if err != nil {
		return
	}

	// update to new ScenarioProcessorId
	values := map[string]interface{}{
		"scenario_processor_id": processor.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(tenantId, debugInterface.ID, values)

	ret = processor

	return

}

func (s *ScenarioNodeService) CopyProcessor(tenantId consts.TenantId, req *agentExec.Processor, CreateBy uint, mod string, rootId *uint) (err *_domain.BizErr) {
	currentNodeReq := s.toProcessorReq(req, CreateBy, mod)

	currentProcessor, err := s.AddProcessor(tenantId, currentNodeReq, "copy")
	if err != nil {
		return err
	}

	if *rootId == 0 && mod == "siblings" {
		*rootId = currentProcessor.ID
	}

	for _, child := range req.Children {
		child.ParentId = currentProcessor.ID
		if err = s.CopyProcessor(tenantId, child, CreateBy, "child", rootId); err != nil {
			return err
		}
	}

	return
}

func (s *ScenarioNodeService) toProcessorReq(req *agentExec.Processor, createBy uint, mod string) (ret serverDomain.ScenarioAddScenarioReq) {
	if mod == "siblings" {
		ret.TargetProcessorId = int(req.ID)
	} else if mod == "child" {
		ret.TargetProcessorId = int(req.ParentId)
	}
	ret.Name = req.Name
	ret.ProcessorCategory = req.EntityCategory
	ret.ProcessorType = req.EntityType
	ret.ProcessorInterfaceSrc = req.ProcessorInterfaceSrc
	ret.ProjectId = req.ProjectId
	ret.CreateBy = createBy
	ret.Mode = mod
	ret.Comments = req.Comments
	ret.Method = req.Method
	ret.SrcProcessorId = req.ID

	return
}

func (s *ScenarioNodeService) GetNodeTree(tenantId consts.TenantId, scenarioId uint, node model.Processor) (root *agentExec.Processor, err error) {
	pos, err := s.ScenarioNodeRepo.ListByScenario(tenantId, scenarioId)
	if err != nil {
		return
	}

	tos := s.ToTos(tenantId, pos, false)

	root = &agentExec.Processor{}
	copier.CopyWithOption(root, &node, copier.Option{DeepCopy: true})

	root.Slots = iris.Map{"icon": "icon"}

	s.ScenarioNodeRepo.MakeTree(tenantId, tos[1:], root)

	return
}
