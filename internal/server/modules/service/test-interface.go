package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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
}

func (s *TestInterfaceService) Load(projectId, serveId int) (ret []*serverDomain.TestInterface, err error) {
	root, err := s.TestInterfaceRepo.GetTree(uint(projectId), uint(serveId))

	if root != nil {
		ret = root.Children
	}

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

func (s *TestInterfaceService) Save(req serverDomain.TestInterfaceSaveReq) (debug model.TestInterface, err error) {
	s.CopyValueFromRequest(&debug, req)
	err = s.TestInterfaceRepo.Save(&debug)

	return
}

func (s *TestInterfaceService) Remove(id int, typ serverConsts.TestInterfaceType) (err error) {
	err = s.TestInterfaceRepo.Remove(uint(id), typ)
	return
}

func (s *TestInterfaceService) CopyValueFromRequest(interf *model.TestInterface, req serverDomain.TestInterfaceSaveReq) {
	copier.CopyWithOption(interf, req, copier.Option{
		DeepCopy: true,
	})
}
