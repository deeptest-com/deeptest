package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"strings"
)

type InterfaceService struct {
	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
}

func (s *InterfaceService) Test(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
	if req.Method == consts.GET {
		ret, _ = httpHelper.Get(req)
	} else if req.Method == consts.POST {
		ret, _ = httpHelper.Post(req)
	} else if req.Method == consts.PUT {
		ret, _ = httpHelper.Put(req)
	} else if req.Method == consts.DELETE {
		ret, _ = httpHelper.Delete(req)
	} else if req.Method == consts.PATCH {
		ret, _ = httpHelper.Patch(req)
	} else if req.Method == consts.HEAD {
		ret, _ = httpHelper.Head(req)
	} else if req.Method == consts.CONNECT {
		ret, _ = httpHelper.Connect(req)
	} else if req.Method == consts.OPTIONS {
		ret, _ = httpHelper.Options(req)
	} else if req.Method == consts.TRACE {
		ret, _ = httpHelper.Trace(req)
	}

	s.GetContentProps(&ret)

	return
}

func (s *InterfaceService) GetContentProps(ret *serverDomain.InvocationResponse) {
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

func (s *InterfaceService) Get(interfId uint) (interf model.Interface, err error) {
	if interfId > 0 {
		interf, err = s.InterfaceRepo.Get(interfId)

		interf.Params, _ = s.InterfaceRepo.ListParams(uint(interfId))
		interf.Headers, _ = s.InterfaceRepo.ListHeaders(uint(interfId))

		interf.BasicAuth, _ = s.InterfaceRepo.GetBasicAuth(uint(interfId))
		interf.BearerToken, _ = s.InterfaceRepo.GetBearerToken(uint(interfId))
		interf.OAuth20, _ = s.InterfaceRepo.GetOAuth20(uint(interfId))
		interf.ApiKey, _ = s.InterfaceRepo.GetApiKey(uint(interfId))
	}

	interf.Params = append(interf.Params, model.InterfaceParam{Name: "", Value: ""})
	interf.Headers = append(interf.Headers, model.InterfaceHeader{Name: "", Value: ""})

	return
}

func (s *InterfaceService) Save(interf *model.Interface) (err error) {
	err = s.InterfaceRepo.Save(interf)

	return
}
func (s *InterfaceService) Create(req serverDomain.InterfaceReq) (interf *model.Interface, err error) {
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

func (s *InterfaceService) UpdateName(req serverDomain.InterfaceReq) (err error) {
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

func (s *InterfaceService) Update(id int, req serverDomain.InterfaceReq) (err error) {

	return
}

func (s *InterfaceService) UpdateByConfig(req serverDomain.InvocationRequest) (err error) {
	interf := model.Interface{}
	s.CopyValueFromRequest(&interf, req)

	err = s.InterfaceRepo.Update(interf)

	return
}
func (s *InterfaceService) UpdateByInvocation(req serverDomain.InvocationRequest) (err error) {
	interf := model.Interface{}
	s.CopyValueFromRequest(&interf, req)

	err = s.InterfaceRepo.Update(interf)

	return
}

func (s *InterfaceService) CopyValueFromRequest(interf *model.Interface, req serverDomain.InvocationRequest) (err error) {
	interf.ID = req.Id

	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}

func (s *InterfaceService) ReplaceEnvironmentExtractorAndExecVariables(req serverDomain.InvocationRequest) (
	ret serverDomain.InvocationRequest, err error) {
	interfaceId := req.Id

	environmentVariables, _ := s.EnvironmentRepo.ListByInterface(interfaceId)
	interfaceExtractorVariables, _ := s.ExtractorRepo.ListExtractorVariableByProject(req.ProjectId)

	variableArr := requestHelper.MergeVariables(environmentVariables, interfaceExtractorVariables, nil)
	requestHelper.ReplaceAll(&req, variableArr)

	ret = req

	return
}
