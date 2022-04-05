package service

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/comm/helper/http"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"strings"
)

type TestInterfaceService struct {
	InterfaceRepo *repo.TestInterfaceRepo `inject:""`
}

func NewTestInterfaceService() *TestInterfaceService {
	return &TestInterfaceService{}
}

func (s *TestInterfaceService) Test(req serverDomain.TestRequest) (ret serverDomain.TestResponse, err error) {
	if req.Method == consts.GET {
		ret, _ = httpHelper.Get(req.Url, req.Params)
	} else if req.Method == consts.POST {
		ret, _ = httpHelper.Post(req.Url, req.Params, req.Body, req.BodyType)
	} else if req.Method == consts.PUT {
		ret, _ = httpHelper.Put(req.Url, req.Params, req.Body, req.BodyType)
	} else if req.Method == consts.DELETE {
		ret, _ = httpHelper.Delete(req.Url, req.Params, req.Body, req.BodyType)
	}

	ret.ContentLang, ret.ContentCharset = s.GetContentProps(ret.ContentType)

	return
}

func (s *TestInterfaceService) GetContentProps(contentType consts.HttpContentType) (
	contentLang consts.HttpRespLangType, charset consts.HttpRespCharset) {

	contentLang = "plaintext"

	if contentType == "" {
		return
	}

	arr := strings.Split(string(contentType), ";")

	arr1 := strings.Split(arr[0], "/")
	if len(arr1) == 1 {
		return
	}

	contentLang = consts.HttpRespLangType(arr1[1])

	arr2 := strings.Split(arr[1], "=")
	if len(arr2) == 1 {
		return
	}

	charset = consts.HttpRespCharset(arr2[1])

	return
}

func (s *TestInterfaceService) GetTree(projectId int) (root *model.TestInterface, err error) {
	root, err = s.InterfaceRepo.GetInterfaceTree(projectId)
	return
}

func (s *TestInterfaceService) Get(interfId int) (interf model.TestInterface, err error) {
	if interfId > 0 {
		interf, err = s.InterfaceRepo.Get(uint(interfId))
	}

	if interf.Params == nil {
		interf.Params = []domain.Param{{Name: "", Value: ""}}
	}
	if interf.Headers == nil {
		interf.Headers = []domain.Header{{Name: "", Value: ""}}
	}

	return
}

func (s *TestInterfaceService) Save(interf *model.TestInterface) (err error) {
	err = s.InterfaceRepo.Save(interf)

	return
}
func (s *TestInterfaceService) Create(req serverDomain.TestInterfaceReq) (interf *model.TestInterface, err error) {
	interf = &model.TestInterface{Name: req.Name, ProjectId: uint(req.ProjectId),
		IsDir: req.Type == serverConsts.Dir}

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

func (s *TestInterfaceService) Delete(projectId, id uint) (err error) {
	err = s.deleteInterfaceAndChildren(projectId, id)

	return
}

func (s *TestInterfaceService) Move(srcId, targetId uint, pos serverConsts.DropPos, projectId uint) (
	srcInterface model.TestInterface, err error) {
	srcInterface, err = s.InterfaceRepo.Get(srcId)

	srcInterface.ParentId, srcInterface.Ordr = s.InterfaceRepo.UpdateOrder(pos, targetId)
	err = s.InterfaceRepo.UpdateOrdAndParent(srcInterface)

	return
}

func (s *TestInterfaceService) deleteInterfaceAndChildren(projectId, interfId uint) (err error) {
	err = s.InterfaceRepo.Delete(interfId)
	if err == nil {
		children, _ := s.InterfaceRepo.GetChildren(projectId, interfId)
		for _, child := range children {
			s.deleteInterfaceAndChildren(child.ProjectId, child.ID)
		}
	}

	return
}

func (s *TestInterfaceService) Update(id int, req serverDomain.TestInterfaceReq) (err error) {

	return
}
