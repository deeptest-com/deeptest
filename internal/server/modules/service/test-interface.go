package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type TestInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	TestInterfaceRepo     *repo.TestInterfaceRepo     `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
}

func (s *TestInterfaceService) Load(projectId, serveId int) (root *serverDomain.TestInterface, err error) {
	root, err = s.TestInterfaceRepo.GetTree(uint(projectId), uint(serveId))

	return
}

func (s *TestInterfaceService) GetTest(endpointInterfaceId uint) (ret serverDomain.TestInterfaceLoadReq, err error) {
	//debugInterfaceId, _ := s.TestInterfaceRepo.HasTestInterfaceRecord(endpointInterfaceId)
	//
	//if debugInterfaceId > 0 {
	//	ret, err = s.GetTestDataFromTestInterface(debugInterfaceId)
	//} else {
	//	ret, err = s.ConvertTestDataFromEndpointInterface(endpointInterfaceId)
	//}

	return
}

func (s *TestInterfaceService) Save(req serverDomain.TestInterfaceCreateReq) (debug model.TestInterface, err error) {
	//s.CopyValueFromRequest(&debug, req)
	//
	//endpointInterface, _ := s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)
	//debug.EndpointId = endpointInterface.EndpointId
	//
	//debugInterfaceId, _ := s.TestInterfaceRepo.HasTestInterfaceRecord(debug.EndpointInterfaceId)
	//if debugInterfaceId > 0 {
	//	debug.ID = debugInterfaceId
	//}
	//
	//err = s.TestInterfaceRepo.Save(&debug)

	return
}
