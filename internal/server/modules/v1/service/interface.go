package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/comm/helper/http"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	requestHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/request"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"strings"
)

type InterfaceService struct {
	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
}

func (s *InterfaceService) Test(req serverDomain.InvocationRequest) (ret serverDomain.InvocationResponse, err error) {
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

	arr2 := strings.Split(arr[1], "=")
	if len(arr2) == 1 {
		return
	}

	ret.ContentCharset = consts.HttpRespCharset(arr2[1])

	return
}

func (s *InterfaceService) GetTree(projectId int) (root *model.Interface, err error) {
	root, err = s.InterfaceRepo.GetInterfaceTree(projectId)
	return
}

func (s *InterfaceService) Get(interfId int) (interf model.Interface, err error) {
	if interfId > 0 {
		interf, err = s.InterfaceRepo.Get(uint(interfId))

		interf.Params, _ = s.InterfaceRepo.ListParams(uint(interfId))
		interf.Headers, _ = s.InterfaceRepo.ListHeaders(uint(interfId))
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
	interf = &model.Interface{Name: req.Name, ProjectId: uint(req.ProjectId),
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

	copier.Copy(interf, req)

	return
}

func (s *InterfaceService) ReplaceVariables(req *serverDomain.InvocationRequest, projectId int) (err error) {
	interfaceId := req.Id

	environment, _ := s.EnvironmentRepo.GetByInterface(interfaceId)
	environmentVariables, _ := s.EnvironmentRepo.GetVars(environment.ID)
	extractorVariables, _ := s.ExtractorRepo.ListExtractorVariable(interfaceId)

	requestHelper.ReplaceVariables(req, environmentVariables, extractorVariables)

	return
}

func replaceUrl(req *serverDomain.InvocationRequest, variableArr [][]string) {
	req.Url = ReplaceValue(req.Url, variableArr, 0)
}
func replaceParams(req *serverDomain.InvocationRequest, variableArr [][]string) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceValue(param.Value, variableArr, 0)
	}
}
func replaceHeaders(req *serverDomain.InvocationRequest, variableArr [][]string) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceValue(header.Value, variableArr, 0)
	}
}
func replaceBody(req *serverDomain.InvocationRequest, variableArr [][]string) {
	req.Body = ReplaceValue(req.Body, variableArr, 0)
}
func replaceAuthor(req *serverDomain.InvocationRequest, variableArr [][]string) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceValue(req.BasicAuth.Username, variableArr, 0)
		req.BasicAuth.Password = ReplaceValue(req.BasicAuth.Password, variableArr, 0)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Username = ReplaceValue(req.BearerToken.Username, variableArr, 0)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Key = ReplaceValue(req.OAuth20.Key, variableArr, 0)
		req.OAuth20.OidcDiscoveryURL = ReplaceValue(req.OAuth20.OidcDiscoveryURL, variableArr, 0)
		req.OAuth20.AuthURL = ReplaceValue(req.OAuth20.AuthURL, variableArr, 0)
		req.OAuth20.AccessTokenURL = ReplaceValue(req.OAuth20.AccessTokenURL, variableArr, 0)
		req.OAuth20.ClientID = ReplaceValue(req.OAuth20.ClientID, variableArr, 0)
		req.OAuth20.Scope = ReplaceValue(req.OAuth20.Scope, variableArr, 0)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Username = ReplaceValue(req.ApiKey.Username, variableArr, 0)
		req.ApiKey.Value = ReplaceValue(req.ApiKey.Value, variableArr, 0)
		req.ApiKey.TransferMode = ReplaceValue(req.ApiKey.TransferMode, variableArr, 0)
	}
}

func genVariableArr(environmentVariables []model.EnvironmentVar, extractorVariables []serverDomain.Variable) (
	ret [][]string) {

	variableMap := iris.Map{}
	for _, item := range environmentVariables {
		variableMap[item.Name] = item.Value
	}
	for _, item := range extractorVariables {
		variableMap[item.Name] = item.Value
	}

	for key, val := range variableMap {
		ret = append(ret, []string{fmt.Sprintf("${%s}", key), val.(string)})
	}

	return
}

func ReplaceValue(value string, variableArr [][]string, index int) (ret string) {
	if len(variableArr) == 0 || !strings.Contains(value, "${") {
		return
	}

	old := variableArr[index][0]
	new := variableArr[index][1]
	ret = strings.ReplaceAll(value, old, new)

	if len(variableArr) > index+1 {
		ret = ReplaceValue(ret, variableArr, index+1)
	}

	return
}
