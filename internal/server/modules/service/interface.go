package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type InterfaceService struct {
	InterfaceRepo         *repo.InterfaceRepo          `inject:""`
	ScenarioInterfaceRepo *repo.ProcessorInterfaceRepo `inject:""`
	EnvironmentRepo       *repo.EnvironmentRepo        `inject:""`
	ExtractorRepo         *repo.ExtractorRepo          `inject:""`

	VariableService *VariableService `inject:""`
}

func (s *InterfaceService) GetTree(projectId int) (root *model.Interface, err error) {
	root, err = s.InterfaceRepo.GetInterfaceTree(projectId)
	return
}

func (s *InterfaceService) Get(interfId uint) (interf model.Interface, err error) {
	if interfId > 0 {
		interf, err = s.InterfaceRepo.Get(interfId)

		interf.Headers, _ = s.InterfaceRepo.ListHeaders(interfId)
		interf.Params, _ = s.InterfaceRepo.ListParams(interfId)

		interf.BodyFormData, _ = s.InterfaceRepo.ListBodyFormData(interfId)
		interf.BodyFormUrlencoded, _ = s.InterfaceRepo.ListBodyFormUrlencoded(interfId)

		interf.BasicAuth, _ = s.InterfaceRepo.GetBasicAuth(interfId)
		interf.BearerToken, _ = s.InterfaceRepo.GetBearerToken(interfId)
		interf.OAuth20, _ = s.InterfaceRepo.GetOAuth20(interfId)
		interf.ApiKey, _ = s.InterfaceRepo.GetApiKey(interfId)
	}

	interf.Headers = append(interf.Headers, model.InterfaceHeader{
		InterfaceHeaderBase: model.InterfaceHeaderBase{Name: "", Value: ""}})
	interf.Params = append(interf.Params, model.InterfaceParam{
		InterfaceParamBase: model.InterfaceParamBase{Name: "", Value: ""}})

	interf.BodyFormData = append(interf.BodyFormData, model.InterfaceBodyFormDataItem{
		InterfaceBodyFormDataItemBase: model.InterfaceBodyFormDataItemBase{Name: "", Value: "", Type: consts.FormDataTypeText}})
	interf.BodyFormUrlencoded = append(interf.BodyFormUrlencoded, model.InterfaceBodyFormUrlEncodedItem{
		InterfaceBodyFormUrlEncodedItemBase: model.InterfaceBodyFormUrlEncodedItemBase{Name: "", Value: ""},
	})

	return
}

func (s *InterfaceService) Save(interf *model.Interface) (err error) {
	err = s.InterfaceRepo.Save(interf)

	return
}
func (s *InterfaceService) Create(req v1.InterfaceReq) (interf *model.Interface, err error) {
	interf = &model.Interface{
		InterfaceBase: model.InterfaceBase{Name: req.Name,
			ProjectId: uint(req.ProjectId),
			IsDir:     req.Type == serverConsts.Dir,
		},
	}

	var dropPos serverConsts.DropPos
	if req.Mode == serverConsts.Child {
		dropPos = serverConsts.Inner
	} else {
		dropPos = serverConsts.After
	}

	interf.ParentId, interf.Ordr = s.InterfaceRepo.UpdateOrder(dropPos, uint(req.Target))
	err = s.InterfaceRepo.Save(interf)

	return
}

func (s *InterfaceService) UpdateName(req v1.InterfaceReq) (err error) {
	err = s.InterfaceRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *InterfaceService) Delete(projectId, id uint) (err error) {
	err = s.deleteInterfaceAndChildren(projectId, id)

	return
}

func (s *InterfaceService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcInterface model.Interface, err error) {
	srcInterface, err = s.InterfaceRepo.Get(srcId)

	srcInterface.ParentId, srcInterface.Ordr = s.InterfaceRepo.UpdateOrder(pos, targetId)
	err = s.InterfaceRepo.UpdateOrdAndParent(srcInterface)

	return
}

func (s *InterfaceService) deleteInterfaceAndChildren(projectId, interfId uint) (err error) {
	err = s.InterfaceRepo.Delete(interfId)
	if err == nil {
		children, _ := s.InterfaceRepo.GetChildren(projectId, interfId)
		for _, child := range children {
			s.deleteInterfaceAndChildren(child.ProjectId, child.ID)
		}
	}

	return
}

func (s *InterfaceService) Update(id int, req v1.InterfaceReq) (err error) {

	return
}

func (s *InterfaceService) UpdateByConfig(req v1.InvocationRequest) (err error) {
	interf := model.Interface{}
	s.CopyValueFromRequest(&interf, req)

	err = s.InterfaceRepo.Update(interf)

	return
}
func (s *InterfaceService) UpdateByInvocation(req v1.InvocationRequest) (err error) {
	interf := model.Interface{}
	s.CopyValueFromRequest(&interf, req)

	err = s.InterfaceRepo.Update(interf)

	return
}

func (s *InterfaceService) CopyValueFromRequest(interf *model.Interface, req v1.InvocationRequest) (err error) {
	interf.ID = req.Id

	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}

func (s *InterfaceService) GetDetail(id uint) (interf model.Interface, err error) {
	return s.InterfaceRepo.GetDetail(id)
}
