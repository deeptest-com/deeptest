package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type TestInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	TestInterfaceRepo     *repo.TestInterfaceRepo     `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	DebugInterfaceService *DebugInterfaceService `inject:""`
}

func (s *TestInterfaceService) Load(projectId, serveId int) (ret []*serverDomain.TestInterface, err error) {
	root, err := s.TestInterfaceRepo.GetTree(uint(projectId), uint(serveId))

	if root != nil {
		ret = root.Children
	}

	return
}

func (s *TestInterfaceService) Get(id int) (ret model.TestInterface, err error) {
	ret, err = s.TestInterfaceRepo.Get(uint(id))

	if ret.DebugInterfaceId > 0 {
		debugInterface, err := s.DebugInterfaceService.GetDetail(ret.DebugInterfaceId)
		if err == nil {
			ret.DebugInterface = &debugInterface
		}
	} else {
		ret.DebugInterface, err = s.DebugInterfaceService.GenSample(ret.ProjectId, ret.ServeId)
	}

	return
}

func (s *TestInterfaceService) Save(req serverDomain.TestInterfaceSaveReq) (debug model.TestInterface, err error) {
	s.CopyValueFromRequest(&debug, req)
	err = s.TestInterfaceRepo.Save(&debug)

	return
}

func (s *TestInterfaceService) Remove(id int, typ serverConsts.TestInterfaceType) (err error) {
	err = s.TestInterfaceRepo.Remove(uint(id), typ)
	return
}

func (s *TestInterfaceService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcScenarioNode model.TestInterface, err error) {
	srcScenarioNode, err = s.TestInterfaceRepo.Get(srcId)

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.TestInterfaceRepo.UpdateOrder(pos, targetId, projectId)
	err = s.TestInterfaceRepo.UpdateOrdAndParent(srcScenarioNode)

	return
}

func (s *TestInterfaceService) SaveDebugData(req domain.DebugData) (debug model.TestInterface, err error) {
	s.CopyDebugDataValueFromRequest(&debug, req)

	//endpointInterface, _ := s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)
	//debug.EndpointId = endpointInterface.EndpointId
	//
	//scenarioInterfaceId, _ := s.TestInterfaceRepo.HasScenarioInterfaceRecord(debug.EndpointInterfaceId)
	//if scenarioInterfaceId > 0 {
	//	debug.ID = scenarioInterfaceId
	//}

	err = s.TestInterfaceRepo.SaveDebugData(&debug)

	return
}

func (s *TestInterfaceService) CopyValueFromRequest(interf *model.TestInterface, req serverDomain.TestInterfaceSaveReq) {
	copier.CopyWithOption(interf, req, copier.Option{
		DeepCopy: true,
	})
}

func (s *TestInterfaceService) CopyDebugDataValueFromRequest(interf *model.TestInterface, req domain.DebugData) (err error) {
	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}
