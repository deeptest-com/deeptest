package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ScenarioNodeService struct {
	ScenarioNodeRepo      *repo.ScenarioNodeRepo      `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
}

func (s *ScenarioNodeService) GetTree(scenarioId int) (root *model.TestProcessor, err error) {
	root, err = s.ScenarioProcessorRepo.GetTree(scenarioId)

	return
}

//func (s *ScenarioNodeService) Get(interfId int) (interf model.ScenarioNode, err error) {
//	if interfId > 0 {
//		interf, err = s.ScenarioNodeRepo.Get(uint(interfId))
//
//		interf.Params, _ = s.ScenarioNodeRepo.ListParams(uint(interfId))
//		interf.Headers, _ = s.ScenarioNodeRepo.ListHeaders(uint(interfId))
//
//		interf.BasicAuth, _ = s.ScenarioNodeRepo.GetBasicAuth(uint(interfId))
//		interf.BearerToken, _ = s.ScenarioNodeRepo.GetBearerToken(uint(interfId))
//		interf.OAuth20, _ = s.ScenarioNodeRepo.GetOAuth20(uint(interfId))
//		interf.ApiKey, _ = s.ScenarioNodeRepo.GetApiKey(uint(interfId))
//	}
//
//	interf.Params = append(interf.Params, model.ScenarioNodeParam{Name: "", Value: ""})
//	interf.Headers = append(interf.Headers, model.ScenarioNodeHeader{Name: "", Value: ""})
//
//	return
//}
//
//func (s *ScenarioNodeService) Save(interf *model.ScenarioNode) (err error) {
//	err = s.ScenarioNodeRepo.Save(interf)
//
//	return
//}
//func (s *ScenarioNodeService) Create(req serverDomain.ScenarioNodeReq) (interf *model.ScenarioNode, err error) {
//	interf = &model.ScenarioNode{Name: req.Name, ProjectId: uint(req.ProjectId),
//		IsDir: req.Type == serverConsts.Dir}
//
//	var dropPos serverConsts.DropPos
//	if req.Mode == serverConsts.Child {
//		dropPos = serverConsts.Inner
//	} else {
//		dropPos = serverConsts.After
//	}
//
//	interf.ParentId, interf.Ordr = s.ScenarioNodeRepo.UpdateOrder(dropPos, uint(req.Target))
//	err = s.ScenarioNodeRepo.Save(interf)
//
//	return
//}
//func (s *ScenarioNodeService) Update(id int, req serverDomain.ScenarioNodeReq) (err error) {
//
//	return
//}
func (s *ScenarioNodeService) UpdateName(req serverDomain.ScenarioNodeReq) (err error) {
	err = s.ScenarioNodeRepo.UpdateName(req.Id, req.Name)
	return
}

//func (s *ScenarioNodeService) Delete(projectId, id uint) (err error) {
//	err = s.deleteScenarioNodeAndChildren(projectId, id)
//
//	return
//}
//
//func (s *ScenarioNodeService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
//	srcScenarioNode model.ScenarioNode, err error) {
//	srcScenarioNode, err = s.ScenarioNodeRepo.Get(srcId)
//
//	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.ScenarioNodeRepo.UpdateOrder(pos, targetId)
//	err = s.ScenarioNodeRepo.UpdateOrdAndParent(srcScenarioNode)
//
//	return
//}
//
//func (s *ScenarioNodeService) deleteScenarioNodeAndChildren(projectId, interfId uint) (err error) {
//	err = s.ScenarioNodeRepo.Delete(interfId)
//	if err == nil {
//		children, _ := s.ScenarioNodeRepo.GetChildren(projectId, interfId)
//		for _, child := range children {
//			s.deleteScenarioNodeAndChildren(child.ProjectId, child.ID)
//		}
//	}
//
//	return
//}
