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
	ExtractorRepo         *repo.ExtractorRepo         `inject:""`
	CheckpointRepo        *repo.CheckpointRepo        `inject:""`

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

func (s *TestInterfaceService) ImportInterfaces(req serverDomain.TestInterfaceImportReq) (ret model.TestInterface, err error) {
	parent, _ := s.TestInterfaceRepo.Get(req.TargetId)

	if parent.Type != serverConsts.TestInterfaceTypeDir {
		parent, _ = s.TestInterfaceRepo.Get(parent.ParentId)
	}

	for _, interfaceId := range req.InterfaceIds {
		ret, err = s.addInterface(interfaceId, req.CreateBy, parent)
	}

	return
}

func (s *TestInterfaceService) addInterface(endpointInterfaceId int, createBy uint, parent model.TestInterface) (
	ret model.TestInterface, err error) {

	endpointInterface, err := s.EndpointInterfaceRepo.Get(uint(endpointInterfaceId))
	if err != nil {
		return
	}

	// convert or clone a debug interface obj
	debugData, err := s.DebugInterfaceService.GetDebugInterfaceByEndpointInterface(uint(endpointInterfaceId))
	debugData.DebugInterfaceId = 0 // force to clone the old one
	debugData.DebugInterfaceId = 0
	debugData.EndpointInterfaceId = uint(endpointInterfaceId)
	debugData.ServeId = parent.ServeId
	debugData.UsedBy = consts.TestDebug
	debugInterface, err := s.DebugInterfaceService.Save(debugData)

	// clone extractors and checkpoints if needed
	if endpointInterface.DebugInterfaceId <= 0 {
		s.ExtractorRepo.CloneFromEndpointInterfaceToDebugInterface(uint(endpointInterfaceId), debugInterface.ID, consts.TestDebug)
		s.CheckpointRepo.CloneFromEndpointInterfaceToDebugInterface(uint(endpointInterfaceId), debugInterface.ID, consts.TestDebug)
	}

	// save test interface
	testInterface := model.TestInterface{
		Title: endpointInterface.Name + "-" + string(endpointInterface.Method),
		Type:  serverConsts.TestInterfaceTypeInterface,
		Ordr:  s.TestInterfaceRepo.GetMaxOrder(parent.ID),

		DebugInterfaceId: debugInterface.ID,
		ParentId:         parent.ID,
		ServeId:          parent.ServeId,
		ProjectId:        parent.ProjectId,
		CreatedBy:        createBy,
	}
	s.TestInterfaceRepo.Save(&testInterface)

	// update test_interface_id
	values := map[string]interface{}{
		"test_interface_id": testInterface.ID,
	}
	s.DebugInterfaceRepo.UpdateDebugInfo(debugInterface.ID, values)

	ret = testInterface

	return
}
