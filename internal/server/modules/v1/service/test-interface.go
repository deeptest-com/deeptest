package service

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/comm/helper/http"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
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
	} else if req.Method == consts.PATCH {
		ret, _ = httpHelper.Patch(req.Url, req.Params, req.Body, req.BodyType)
	} else if req.Method == consts.HEAD {
		ret, _ = httpHelper.Head(req.Url, req.Params)
	} else if req.Method == consts.CONNECT {
		ret, _ = httpHelper.Connect(req.Url, req.Params)
	} else if req.Method == consts.OPTIONS {
		ret, _ = httpHelper.Options(req.Url, req.Params)
	} else if req.Method == consts.TRACE {
		ret, _ = httpHelper.Trace(req.Url, req.Params)
	}

	s.GetContentProps(&ret)

	return
}

func (s *TestInterfaceService) GetContentProps(ret *serverDomain.TestResponse) {
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

	arr2 := strings.Split(arr[1], "=")
	if len(arr2) == 1 {
		return
	}

	ret.ContentCharset = consts.HttpRespCharset(arr2[1])

	return
}

func (s *TestInterfaceService) GetTree(projectId int) (root *model.TestInterface, err error) {
	root, err = s.InterfaceRepo.GetInterfaceTree(projectId)
	return
}

func (s *TestInterfaceService) Get(interfId int) (interf model.TestInterface, err error) {
	if interfId > 0 {
		interf, err = s.InterfaceRepo.Get(uint(interfId))

		interf.Params, _ = s.InterfaceRepo.ListParams(uint(interfId))
		interf.Headers, _ = s.InterfaceRepo.ListHeaders(uint(interfId))
	}

	interf.Params = append(interf.Params, model.TestInterfaceParam{Name: "", Value: ""})

	interf.Headers = append(interf.Headers, model.TestInterfaceHeader{Name: "", Value: ""})

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

func (s *TestInterfaceService) UpdateByConfig(req serverDomain.TestRequest) (err error) {
	interf := model.TestInterface{}
	s.CopyValueFromRequest(&interf, req)

	err = s.InterfaceRepo.Update(interf)

	return
}
func (s *TestInterfaceService) UpdateByRequest(req serverDomain.TestRequest) (err error) {
	interf := model.TestInterface{}
	s.CopyValueFromRequest(&interf, req)

	err = s.InterfaceRepo.Update(interf)

	return
}

func (s *TestInterfaceService) CopyValueFromRequest(interf *model.TestInterface, req serverDomain.TestRequest) (err error) {
	interf.ID = req.Id

	copier.Copy(interf, req)

	return
}
