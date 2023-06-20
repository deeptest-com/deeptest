package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`

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
	// its debug data will load in webpage

	return
}

func (s *TestInterfaceService) Save(req serverDomain.TestInterfaceSaveReq) (testInterface model.TestInterface, err error) {
	s.CopyValueFromRequest(&testInterface, req)

	// create new DebugInterface
	debugInterface := model.DebugInterface{
		InterfaceBase: model.InterfaceBase{
			Name: req.Title,
			InterfaceConfigBase: model.InterfaceConfigBase{
				Method: consts.GET,
			},
		},
	}
	err = s.DebugInterfaceRepo.Save(&debugInterface)
	testInterface.DebugInterfaceId = debugInterface.ID

	err = s.TestInterfaceRepo.Save(&testInterface)

	values := map[string]interface{}{
		"test_interface_id": testInterface.ID,
	}
	err = s.DebugInterfaceRepo.UpdateDebugInfo(debugInterface.ID, values)

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

func (s *TestInterfaceService) SaveDebugData(req domain.DebugData) (debugInterface model.DebugInterface, err error) {
	s.DebugInterfaceService.Save(req)

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
