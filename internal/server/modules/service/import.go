package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"strings"
)

type InterfaceService struct {
	InterfaceRepo   *repo2.InterfaceRepo   `inject:""`
	EnvironmentRepo *repo2.EnvironmentRepo `inject:""`
	ExtractorRepo   *repo2.ExtractorRepo   `inject:""`

	VariableService *VariableService `inject:""`
}

func (s *InterfaceService) Test(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	if req.Method == consts.GET {
		ret, err = httpHelper.Get(req)
	} else if req.Method == consts.POST {
		ret, err = httpHelper.Post(req)
	} else if req.Method == consts.PUT {
		ret, err = httpHelper.Put(req)
	} else if req.Method == consts.DELETE {
		ret, err = httpHelper.Delete(req)
	} else if req.Method == consts.PATCH {
		ret, err = httpHelper.Patch(req)
	} else if req.Method == consts.HEAD {
		ret, err = httpHelper.Head(req)
	} else if req.Method == consts.CONNECT {
		ret, err = httpHelper.Connect(req)
	} else if req.Method == consts.OPTIONS {
		ret, err = httpHelper.Options(req)
	} else if req.Method == consts.TRACE {
		ret, err = httpHelper.Trace(req)
	}

	s.GetContentProps(&ret)

	return
}

func (s *InterfaceService) GetContentProps(ret *v1.InvocationResponse) {
	ret.ContentLang = "plaintext"

	if ret.ContentLang == "" {
		return
	}

	arr := strings.Split(string(ret.ContentType), ";")

	arr1 := strings.Split(arr[0], "/")
	if len(arr1) == 1 {
		return
	}

	typeName := arr1[1]
	if typeName == "text" || typeName == "plain" {
		typeName = "plaintext"
	}
	ret.ContentLang = consts.HttpRespLangType(typeName)

	if len(arr) > 1 {
		arr2 := strings.Split(arr[1], "=")
		if len(arr2) > 1 {
			ret.ContentCharset = consts.HttpRespCharset(arr2[1])
		}
	}

	//ret.Content = mockHelper.FormatXml(ret.Content)

	return
}

func (s *InterfaceService) GetTree(projectId int) (root *model.Interface, err error) {
	root, err = s.InterfaceRepo.GetInterfaceTree(projectId)
	return
}

func (s *InterfaceService) Save(interf *model.Interface) (err error) {
	err = s.InterfaceRepo.Save(interf)

	return
}
func (s *InterfaceService) Create(req v1.InterfaceReq) (interf *model.Interface, err error) {
	interf = &model.Interface{
		Name:      req.Name,
		ProjectId: uint(req.ProjectId),
		IsDir:     req.Type == serverConsts.Dir,
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
