package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type WebsocketInterfaceService struct {
	WebsocketInterfaceRepo *repo.WebsocketInterfaceRepo `inject:""`
	DiagnoseInterfaceRepo  *repo.DiagnoseInterfaceRepo  `inject:""`
}

func (s *WebsocketInterfaceService) GetDebugData(tenantId consts.TenantId, diagnoseInterfaceId int) (
	ret domain.WebsocketDebugData, err error) {

	diagnose, err := s.DiagnoseInterfaceRepo.Get(tenantId, uint(diagnoseInterfaceId))
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, diagnose, copier.Option{
		DeepCopy: true,
	})

	po, err := s.WebsocketInterfaceRepo.Get(tenantId, diagnose.DebugInterfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, po, copier.Option{
		DeepCopy: true,
	})

	ret.Params, _ = s.ListParams(tenantId, po.ID)
	ret.Headers, _ = s.ListHeaders(tenantId, po.ID)

	if ret.Params == nil {
		ret.Params = &[]domain.Param{}
	}
	*ret.Params = append(*ret.Params, domain.Param{Name: "", Value: ""})

	if ret.Headers == nil {
		ret.Headers = &[]domain.Header{}
	}
	*ret.Headers = append(*ret.Headers, domain.Header{Name: "", Value: ""})

	return
}

func (s *WebsocketInterfaceService) SaveDebugData(data domain.WebsocketDebugData, tenantId consts.TenantId) (err error) {
	po := model.WebsocketInterface{}
	copier.CopyWithOption(&po, data, copier.Option{DeepCopy: true})
	err = s.WebsocketInterfaceRepo.SaveDebugData(po, tenantId)

	err = s.WebsocketInterfaceRepo.SaveParams(data.Params, po.ID, tenantId)
	err = s.WebsocketInterfaceRepo.SaveHeaders(data.Headers, po.ID, tenantId)

	return
}

func (s *WebsocketInterfaceService) ListParams(tenantId consts.TenantId, id uint) (
	ret *[]domain.Param, err error) {

	ret = &[]domain.Param{}

	pos, _ := s.WebsocketInterfaceRepo.ListParams(tenantId, id)

	for _, po := range pos {
		param := domain.Param{
			Name:     po.Name,
			Value:    po.Value,
			Disabled: po.Disabled,
			Type:     po.Type,
			IsGlobal: false,
		}

		(*ret) = append((*ret), param)
	}

	return
}

func (s *WebsocketInterfaceService) ListHeaders(tenantId consts.TenantId, id uint) (
	ret *[]domain.Header, err error) {

	ret = &[]domain.Header{}

	pos, _ := s.WebsocketInterfaceRepo.ListHeaders(tenantId, id)

	for _, po := range pos {
		param := domain.Header{
			Name:     po.Name,
			Value:    po.Value,
			Disabled: po.Disabled,
			Type:     po.Type,
		}

		(*ret) = append((*ret), param)
	}

	return
}
