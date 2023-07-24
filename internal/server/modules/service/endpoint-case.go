package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type EndpointCaseService struct {
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`

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

func (s *EndpointCaseService) Save(req serverDomain.EndpointCaseSaveReq) (po model.EndpointCase, err error) {
	s.CopyValueFromRequest(&po, req)

	endpoint, err := s.EndpointRepo.Get(req.EndpointId)
	server, _ := s.ServeServerRepo.GetDefaultByServe(endpoint.ServeId)

	// create new DebugInterface
	url := req.DebugData.Url
	if url == "" {
		url = endpoint.Path
	}

	debugInterface := model.DebugInterface{
		InterfaceBase: model.InterfaceBase{
			Name: req.Name,

			InterfaceConfigBase: model.InterfaceConfigBase{
				Method: consts.GET,
				Url:    url,
			},
		},
		ServeId:  endpoint.ServeId,
		ServerId: server.ID,
		BaseUrl:  server.Url,
	}

	err = s.DebugInterfaceRepo.Save(&debugInterface)

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
	debugData := req.DebugData

	// save debug data
	req.DebugData.UsedBy = consts.CaseDebug
	debugInterface, err := s.DebugInterfaceService.SaveAs(debugData)

	// save case
	s.CopyValueFromRequest(&po, req)

	if po.EndpointId == 0 {
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(uint(req.EndpointInterfaceId))
		po.EndpointId = endpointInterface.EndpointId
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
