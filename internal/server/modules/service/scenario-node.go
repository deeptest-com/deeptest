package service

import (
	"encoding/json"
	"errors"
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	curlHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/gcurl"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
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

func (s *ScenarioNodeService) GetTree(scenario model.Scenario, withDetail bool) (root *agentExec.Processor, err error) {
	pos, err := s.ScenarioNodeRepo.ListByScenario(scenario.ID)
	if err != nil {
		return
	}

	tos := s.ToTos(pos, withDetail)

	root = tos[0]
	root.Name = scenario.Name
	root.Slots = iris.Map{"icon": "icon"}

	s.ScenarioNodeRepo.MakeTree(tos[1:], root)

	root.Session = agentExec.Session{}

	root.ScenarioId = scenario.ID
	root.ProjectId = scenario.ProjectId

	return
}

func (s *ScenarioNodeService) ToTos(pos []*model.Processor, withDetail bool) (tos []*agentExec.Processor) {
	for _, po := range pos {
		to := agentExec.Processor{
			ProcessorBase: agentExec.ProcessorBase{
				Session:    agentExec.Session{},
				ScenarioId: po.ScenarioId,
			},
		}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})
		to.Disable = po.Disabled
		to.Comments = po.Comments

		if withDetail {
			entity, _ := s.ScenarioProcessorService.GetEntityTo(&to)
			to.EntityRaw, _ = json.Marshal(entity)
		}

		// just to avoid json marshal error for IProcessorEntity
		to.Entity = agentExec.ProcessorGroup{}

		tos = append(tos, &to)
	}

	return
}

func (s *ScenarioNodeService) AddProcessor(req serverDomain.ScenarioAddScenarioReq, source string) (ret model.Processor, err *_domain.BizErr) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(uint(req.TargetProcessorId))
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
		s.ScenarioNodeRepo.MoveMaxOrder(ret.ParentId, uint(targetProcessor.Ordr), 1)
		ret.Ordr = targetProcessor.Ordr + 1
	} else {
		ret.Ordr = s.ScenarioNodeRepo.GetMaxOrder(ret.ParentId)
	}

	s.ScenarioNodeRepo.Save(&ret)

	if req.Mode == "parent" { // move interface to new folder
		targetProcessor.ParentId = ret.ID
		s.ScenarioNodeRepo.Save(&targetProcessor)
	}

	if source == "copy" {
		if ret.EntityType == consts.ProcessorInterfaceDefault {
			s.CopyInterfaceEntity(req.SrcProcessorId, ret.ID)
		} else {
			s.ScenarioProcessorRepo.CopyEntity(req.SrcProcessorId, ret.ID)
		}
	} else {
		if ret.EntityType == consts.ProcessorInterfaceDefault { // create debug interface
			debugInterfaceId, _ := s.DebugInterfaceService.CreateDefault(ret.ProcessorInterfaceSrc, req.ProjectId)
			s.ScenarioProcessorRepo.UpdateInterfaceId(ret.ID, debugInterfaceId)
		} else if ret.EntityType == consts.ProcessorLogicElse { // create default entity
			entity := model.ProcessorLogic{
				ProcessorEntityBase: agentExec.ProcessorEntityBase{
					ProcessorID:       ret.ID,
					ProcessorCategory: ret.EntityCategory,
					ProcessorType:     ret.EntityType,
				},
			}
			_ = s.ScenarioProcessorRepo.SaveLogic(&entity)

		}
	}

	return
}

func (s *ScenarioNodeService) CopyInterfaceEntity(srcProcessorId, distProcessorId uint) (err error) {
	srcProcessor, err := s.ScenarioNodeRepo.Get(srcProcessorId)
	if err != nil {
		return
	}

	debugData, err := s.DebugInterfaceService.GetDebugDataFromDebugInterface(srcProcessor.EntityId)
	if err != nil {
		return
	}

	debugData.ScenarioProcessorId = distProcessorId
	debugData.UsedBy = consts.ScenarioDebug
	debugInterface, err := s.DebugInterfaceService.SaveAs(debugData, debugData.DebugInterfaceId, debugData.UsedBy)
	if err != nil {
		return
	}

	err = s.ScenarioProcessorRepo.UpdateInterfaceId(distProcessorId, debugInterface.ID)
	if err != nil {
		return err
	}

	return
}
func (s *ScenarioNodeService) AddInterfacesFromDiagnose(req serverDomain.ScenarioAddInterfacesFromTreeReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(targetProcessor.ParentId)
	}

	for _, interfaceNode := range req.SelectedNodes {
		ret, _ = s.createDirOrInterfaceFromDiagnose(&interfaceNode, targetProcessor, 0)
	}

	return
}

func (s *ScenarioNodeService) AddInterfacesFromDefine(req serverDomain.ScenarioAddInterfacesReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(targetProcessor.ParentId)
	}

	serveId := uint(0)
	for _, interfaceId := range req.InterfaceIds {
		ret, err = s.createInterfaceFromDefine(uint(interfaceId), &serveId, req.CreateBy, targetProcessor, "", 0)
	}

	return
}

func (s *ScenarioNodeService) AddInterfacesFromCase(req serverDomain.ScenarioAddCasesFromTreeReq) (ret model.Processor, err error) {
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(targetProcessor.ParentId)
	}

	for _, interfaceNode := range req.SelectedNodes {
		ret, _ = s.createDirOrInterfaceFromCase(&interfaceNode, targetProcessor, 0)
	}

	return
}

func (s *ScenarioNodeService) createInterfaceFromDefine(endpointInterfaceId uint, serveId *uint,
	createBy uint, parentProcessor model.Processor, name string, order int) (
	ret model.Processor, err error) {

	endpointInterface, err := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if err == gorm.ErrRecordNotFound {
		err = errors.New("interface is deleted")
		return
	}

	// get serve id once
	if *serveId == 0 {
		endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)
		*serveId = endpoint.ServeId
	}

	// convert or clone a debug interface obj
	debugData, err := s.DebugInterfaceService.GetDebugDataFromEndpointInterface(endpointInterfaceId)
	debugData.UsedBy = consts.ScenarioDebug

	debugData.EndpointInterfaceId = endpointInterfaceId

	debugData.ScenarioProcessorId = 0 // will be update after ScenarioProcessor saved
	debugData.ProcessorInterfaceSrc = consts.InterfaceSrcDefine

	debugData.ServeId = *serveId

	if debugData.ServerId == 0 {
		server, _ := s.ServeServerRepo.GetDefaultByServe(debugData.ServeId)
		debugData.ServerId = server.ID
	}

	server, _ := s.ServeServerRepo.Get(debugData.ServerId)
	debugData.ServerId = server.ID
	debugData.BaseUrl = server.Url

	srcDebugInterfaceId := debugData.DebugInterfaceId
	debugInterface, err := s.DebugInterfaceService.SaveAs(debugData, srcDebugInterfaceId, "")

	if order == 0 {
		order = s.ScenarioNodeRepo.GetMaxOrder(parentProcessor.ID)
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

	s.ScenarioNodeRepo.Save(&processor)

	// update to new ScenarioProcessorId
	values := map[string]interface{}{
		"scenario_processor_id": processor.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(debugInterface.ID, values)

	ret = processor

	return
}

func (s *ScenarioNodeService) createDirOrInterfaceFromDiagnose(diagnoseInterfaceNode *serverDomain.DiagnoseInterface, parentProcessor model.Processor, order int) (
	ret model.Processor, err error) {

	debugData, _ := s.DebugInterfaceService.GetDebugDataFromDebugInterface(diagnoseInterfaceNode.DebugInterfaceId)
	debugData.UsedBy = consts.ScenarioDebug

	if diagnoseInterfaceNode.IsDir && len(diagnoseInterfaceNode.Children) > 0 { // dir
		for _, child := range diagnoseInterfaceNode.Children {
			ret, _ = s.createDirOrInterfaceFromDiagnose(child, parentProcessor, 0)
		}

	} else if !diagnoseInterfaceNode.IsDir { // interface
		if order == 0 {
			order = s.ScenarioNodeRepo.GetMaxOrder(parentProcessor.ID)
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
		s.ScenarioNodeRepo.Save(&processor)

		// convert or clone a debug interface obj
		debugData.ScenarioProcessorId = processor.ID
		debugData.ProcessorInterfaceSrc = consts.InterfaceSrcDiagnose

		debugData.ServeId = diagnoseInterfaceNode.ServeId

		debugInterfaceOfDiagnoseInterfaceNode, _ := s.DebugInterfaceRepo.Get(diagnoseInterfaceNode.DebugInterfaceId)
		debugData.ServerId = debugInterfaceOfDiagnoseInterfaceNode.ServerId

		debugData.BaseUrl = "" // no need to bind to env in debug page
		debugData.Url = debugInterfaceOfDiagnoseInterfaceNode.Url

		srcDebugInterfaceId := debugData.DebugInterfaceId
		debugInterface, _ := s.DebugInterfaceService.SaveAs(debugData, srcDebugInterfaceId, consts.DiagnoseDebug)

		processor.EntityId = debugInterface.ID
		s.ScenarioProcessorRepo.UpdateEntityId(processor.ID, processor.EntityId)

		ret = processor
	}

	return
}

func (s *ScenarioNodeService) createDirOrInterfaceFromCase(caseNode *serverDomain.EndpointCaseTree, parentProcessor model.Processor, order int) (
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
			processor, _ = s.createDirOrInterfaceFromCase(child, parentProcessor, 0)
		}
	} else if !caseNode.IsDir { // interface
		debugData, _ := s.DebugInterfaceService.GetDebugDataFromDebugInterface(caseNode.DebugInterfaceId)
		debugData.UsedBy = consts.ScenarioDebug

		if order == 0 {
			order = s.ScenarioNodeRepo.GetMaxOrder(parentProcessor.ID)
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

		s.ScenarioNodeRepo.Save(&processor)

		debugData.ScenarioProcessorId = processor.ID
		debugData.ProcessorInterfaceSrc = consts.InterfaceSrcCase

		debugInterfaceOfCaseNode, _ := s.DebugInterfaceRepo.GetByCaseInterfaceId(uint(caseNode.Key))
		debugData.Body = debugInterfaceOfCaseNode.Body
		debugData.ServerId = debugInterfaceOfCaseNode.ServerId

		debugData.BaseUrl = "" // no need to bind to env in debug page
		debugData.Url = debugInterfaceOfCaseNode.Url

		srcDebugInterfaceId := debugData.DebugInterfaceId
		debugInterface, _ := s.DebugInterfaceService.SaveAs(debugData, srcDebugInterfaceId, consts.CaseDebug)

		processor.EntityId = debugInterface.ID
		s.ScenarioProcessorRepo.UpdateEntityId(processor.ID, processor.EntityId)

	}

	return
}

func (s *ScenarioNodeService) UpdateName(req serverDomain.ScenarioNodeReq) (err error) {
	err = s.ScenarioNodeRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *ScenarioNodeService) Delete(id uint) (err error) {
	err = s.deleteScenarioNodeAndChildren(id)

	return
}

func (s *ScenarioNodeService) DisableOrNot(id uint) (err error) {
	err = s.disableScenarioNodeAndChildren(id)

	return
}

func (s *ScenarioNodeService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcScenarioNode model.Processor, err error) {
	srcScenarioNode, err = s.ScenarioNodeRepo.Get(srcId)

	//获取下一个节点，如果是else准备移动
	srcScenarioNextNode, nextNodeErr := s.ScenarioNodeRepo.GetNextNode(srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr, srcScenarioNode.Disabled = s.ScenarioNodeRepo.UpdateOrder(pos, targetId)

	err = s.ScenarioNodeRepo.UpdateOrdAndParent(srcScenarioNode)

	//判断节点是else一起移动
	if nextNodeErr == nil && srcScenarioNextNode.EntityType == consts.ProcessorLogicElse {
		s.Move(srcScenarioNextNode.ID, srcScenarioNode.ID, serverConsts.After, srcScenarioNode.ProjectId)
	}

	return
}

func (s *ScenarioNodeService) deleteScenarioNodeAndChildren(nodeId uint) (err error) {
	srcScenarioNextNode, nextNodeErr := s.ScenarioNodeRepo.GetNextNode(nodeId)
	err = s.ScenarioNodeRepo.DeleteWithChildren(nodeId)
	if nextNodeErr == nil && srcScenarioNextNode.EntityType == consts.ProcessorLogicElse {
		err = s.ScenarioNodeRepo.DeleteWithChildren(srcScenarioNextNode.ID)
	}

	return
}
func (s *ScenarioNodeService) disableScenarioNodeAndChildren(nodeId uint) (err error) {
	node, err := s.ScenarioNodeRepo.Get(nodeId)
	srcScenarioNextNode, nextNodeErr := s.ScenarioNodeRepo.GetNextNode(nodeId)

	err = s.ScenarioNodeRepo.DisableWithChildren(nodeId)

	//如果要禁用if，则else也要禁用
	if !node.Disabled && nextNodeErr == nil && !srcScenarioNextNode.Disabled && srcScenarioNextNode.EntityType == consts.ProcessorLogicElse {
		err = s.ScenarioNodeRepo.DisableWithChildren(srcScenarioNextNode.ID)
	}

	//如果启动else，则if则被启用
	if node.Disabled && node.EntityType == consts.ProcessorLogicElse {
		srcScenarioPreNode, preNodeErr := s.ScenarioNodeRepo.GetPreNode(nodeId)
		if preNodeErr == nil && srcScenarioPreNode.Disabled {
			s.disableScenarioNodeAndChildren(srcScenarioPreNode.ID)
		}
	}

	return
}

func (s *ScenarioNodeService) ListToByScenario(id uint) (ret []*agentExec.Processor, err error) {
	pos, _ := s.ScenarioNodeRepo.ListByScenario(id)

	for _, po := range pos {
		to := agentExec.Processor{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		ret = append(ret, &to)
	}

	return
}

func (s *ScenarioNodeService) ImportCurl(req serverDomain.ScenarioCurlImportReq) (ret model.Processor, err error) {
	//req.Content = "curl 'https://leyanapi.nancalcloud.com/api/v1/endpoint/detail?id=1783&ts=1691139567363&currProjectId=2&lang=zh-CN' \\\n  -H 'Accept: application/json, text/plain, */*' \\\n  -H 'Accept-Language: zh-CN,zh;q=0.9' \\\n  -H 'Authorization: Bearer WldJME16Z3lZelV6T0RSbFlXWXdaV0psWWpKbU56TTROR0U1TjJJd01UVS5OV1l3TkRBMU16WmlZekk1WldGa05UTmtOelF4WldRNVlUTXlaamczTURr' \\\n  -H 'Connection: keep-alive' \\\n  -H 'Cookie: td_cookie=3753276482; HWWAFSESID=8f1d573ac34f71a4c5; HWWAFSESTIME=1690431343235' \\\n  -H 'Referer: https://leyanapi.nancalcloud.com/' \\\n  -H 'Sec-Fetch-Dest: empty' \\\n  -H 'Sec-Fetch-Mode: cors' \\\n  -H 'Sec-Fetch-Site: same-origin' \\\n  -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36' \\\n  -H 'sec-ch-ua: \"Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"115\", \"Chromium\";v=\"115\"' \\\n  -H 'sec-ch-ua-mobile: ?0' \\\n  -H 'sec-ch-ua-platform: \"Windows\"' \\\n  --compressed"
	targetProcessor, _ := s.ScenarioProcessorRepo.Get(req.TargetId)

	if !s.ScenarioNodeRepo.IsDir(targetProcessor) {
		targetProcessor, _ = s.ScenarioProcessorRepo.Get(targetProcessor.ParentId)
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

	debugInterface, err := s.DebugInterfaceService.SaveAs(debugData, 0, "")

	processor := model.Processor{
		Name:                url,
		Method:              debugData.Method,
		EntityCategory:      consts.ProcessorInterface,
		EntityType:          consts.ProcessorInterfaceDefault,
		EntityId:            debugInterface.ID, // as debugInterfaceId
		EndpointInterfaceId: debugInterface.EndpointInterfaceId,

		Ordr: s.ScenarioNodeRepo.GetMaxOrder(targetProcessor.ID),

		ParentId:              targetProcessor.ID,
		ScenarioId:            targetProcessor.ScenarioId,
		ProjectId:             targetProcessor.ProjectId,
		CreatedBy:             req.CreateBy,
		ProcessorInterfaceSrc: consts.InterfaceSrcCurl,
		BaseModel:             model.BaseModel{Disabled: targetProcessor.Disabled},
	}

	err = s.ScenarioNodeRepo.Save(&processor)
	if err != nil {
		return
	}

	// update to new ScenarioProcessorId
	values := map[string]interface{}{
		"scenario_processor_id": processor.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(debugInterface.ID, values)

	ret = processor

	return

}

func (s *ScenarioNodeService) CopyProcessor(req *agentExec.Processor, CreateBy uint, mod string, rootId *uint) (err *_domain.BizErr) {
	currentNodeReq := s.toProcessorReq(req, CreateBy, mod)

	currentProcessor, err := s.AddProcessor(currentNodeReq, "copy")
	if err != nil {
		return err
	}

	if *rootId == 0 && mod == "siblings" {
		*rootId = currentProcessor.ID
	}

	for _, child := range req.Children {
		child.ParentId = currentProcessor.ID
		if err = s.CopyProcessor(child, CreateBy, "child", rootId); err != nil {
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

func (s *ScenarioNodeService) GetNodeTree(scenarioId uint, node model.Processor) (root *agentExec.Processor, err error) {
	pos, err := s.ScenarioNodeRepo.ListByScenario(scenarioId)
	if err != nil {
		return
	}

	tos := s.ToTos(pos, false)

	root = &agentExec.Processor{}
	copier.CopyWithOption(root, &node, copier.Option{DeepCopy: true})

	root.Slots = iris.Map{"icon": "icon"}

	s.ScenarioNodeRepo.MakeTree(tos[1:], root)

	return
}
