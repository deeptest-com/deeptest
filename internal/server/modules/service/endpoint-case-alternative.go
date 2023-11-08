package service

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/kataras/iris/v12"
	"github.com/tidwall/sjson"
	"strings"
)

type EndpointCaseAlternativeService struct {
	EndpointCaseAlternativeRepo *repo.EndpointCaseAlternativeRepo `inject:""`
	EndpointCaseRepo            *repo.EndpointCaseRepo            `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	PreConditionRepo      *repo.PreConditionRepo      `inject:""`
	PostConditionRepo     *repo.PostConditionRepo     `inject:""`
	CategoryRepo          *repo.CategoryRepo          `inject:""`

	EndpointCaseService      *EndpointCaseService      `inject:""`
	EndpointService          *EndpointService          `inject:""`
	DebugInterfaceService    *DebugInterfaceService    `inject:""`
	EndpointMockParamService *EndpointMockParamService `inject:""`

	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	SceneService    *SceneService         `inject:""`
}

func (s *EndpointCaseAlternativeService) LoadAlternative(endpointId uint, method consts.HttpMethod) (
	root casesHelper.AlternativeCase, err error) {

	root.Title = "备选用例"
	root.Category = consts.AlternativeCaseRoot
	root.Key = _stringUtils.Uuid()
	root.Slots = iris.Map{"icon": "icon"}
	root.IsDir = true

	//casePo, _ := s.EndpointCaseRepo.Get(baseId)

	_, endpointInterfaceId := s.EndpointInterfaceRepo.GetByMethod(endpointId, method)
	if endpointInterfaceId == 0 {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	endpoint, err := s.EndpointRepo.GetWithInterface(endpointInterface.EndpointId, "v0.1.0")
	if err != nil {
		return
	}

	// get spec
	doc3 := s.EndpointService.Yaml(endpoint)

	// TEST:
	//pth := "/Users/aaron/rd/project/gudi/deeptest/xdoc/openapi/openapi3/test2.yaml"
	//loader := &openapi3.Loader{Context: context.Background(), IsExternalRefsAllowed: true}
	//doc3, err := loader.LoadFromFile(pth)
	//if err != nil || doc3 == nil {
	//	return
	//}

	apiPathItem, _ := casesHelper.GetApiPathItem(doc3)

	apiOperation, err := casesHelper.GetApiOperation(method, apiPathItem)
	if err != nil || apiOperation == nil {
		return
	}

	root.Children = append(root.Children, casesHelper.LoadForPathParams(apiPathItem.Parameters))
	root.Children = append(root.Children, casesHelper.LoadForQueryParams(apiOperation.Parameters))
	root.Children = append(root.Children, casesHelper.LoadForHeaders(apiOperation.Parameters))
	root.Children = append(root.Children, casesHelper.LoadForBody(apiOperation.RequestBody, doc3))

	return
}

func (s *EndpointCaseAlternativeService) LoadFactor(caseId uint) (ret map[string]model.EndpointCaseAlternativeFactor, err error) {
	pos, err := s.EndpointCaseAlternativeRepo.LoadFactor(caseId)

	for _, po := range pos {
		ret[po.Path] = po
	}

	return
}

func (s *EndpointCaseAlternativeService) CreateBenchmarkCase(req serverDomain.EndpointCaseBenchmarkCreateReq) (
	po model.EndpointCase, err error) {

	if req.BaseCaseId > 0 {
		// clone
		//po, _ = s.EndpointCaseService.Copy(req.BaseCaseId, "alter-", req.CreateUserId, req.CreateUserName)
		po, err = s.EndpointCaseService.Get(uint(req.BaseCaseId))
	} else if req.EndpointInterfaceId > 0 {
		// convert from endpoint interface define
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)
		debugData, _ := s.DebugInterfaceService.GetDebugInterfaceByEndpointInterface(req.EndpointInterfaceId)

		saveReq := serverDomain.EndpointCaseSaveReq{
			Name:           req.Name,
			Method:         debugData.Method,
			DebugData:      debugData,
			EndpointId:     endpointInterface.EndpointId,
			CreateUserId:   req.CreateUserId,
			CreateUserName: req.CreateUserName,
		}

		po, err = s.EndpointCaseService.SaveFromDebugInterface(saveReq)
	}
	if err != nil {
		return
	}

	po.CaseType = consts.CaseBenchmark
	po.BaseCase = uint(req.BaseCaseId)

	s.EndpointCaseRepo.UpdateInfo(po.ID, map[string]interface{}{
		"case_type": po.CaseType,
		"base_case": po.BaseCase,
	})

	if req.BaseCaseId > 0 {
		s.PreConditionRepo.CloneAll(po.DebugInterfaceId, 0, po.DebugInterfaceId, consts.CaseDebug, consts.CaseDebug)
		s.PostConditionRepo.CloneAll(po.DebugInterfaceId, 0, po.DebugInterfaceId, consts.CaseDebug, consts.CaseDebug)
	}

	return
}

func (s *EndpointCaseAlternativeService) SaveFactor(req serverDomain.EndpointCaseFactorSaveReq) (err error) {
	err = s.EndpointCaseAlternativeRepo.SaveFactor(req)

	return
}

func (s *EndpointCaseAlternativeService) SaveCase(req serverDomain.EndpointCaseAlternativeSaveReq) (count int, err error) {
	typ := req.Type
	if typ == "multi" {
		count, err = s.GenMultiCases(req)
	} else if typ == "single" {
		count, err = s.GenSingleCase(req)
	}

	return
}

func (s *EndpointCaseAlternativeService) GenMultiCases(req serverDomain.EndpointCaseAlternativeSaveReq) (count int, err error) {
	for _, val := range req.Values {
		if val.Category != consts.AlternativeCaseCase {
			continue
		}

		newEndpointCase, err1 := s.EndpointCaseService.Copy(req.BaseId, "extend-", req.CreateUserId, req.CreateUserName)
		if err1 != nil {
			err = err1
			return
		}

		s.EndpointCaseRepo.UpdateInfo(newEndpointCase.ID, map[string]interface{}{
			"case_type": consts.CaseAlternative,
		})

		newDebugData, err1 := s.DebugInterfaceService.GetDebugDataFromDebugInterface(newEndpointCase.DebugInterfaceId)
		if err1 != nil {
			err = err1
			return
		}

		fieldIn, fieldNameOrPath := s.getFieldProps(val.Path)
		if fieldIn == "" {
			logUtils.Info("failed to getFieldProps")
			continue
		}
		s.changeFieldProps(&newDebugData, fieldIn, fieldNameOrPath, val.Sample, val.FieldType)

		_, err = s.DebugInterfaceService.Update(newDebugData, newDebugData.DebugInterfaceId)

		count += 1
	}

	return
}

func (s *EndpointCaseAlternativeService) GenSingleCase(req serverDomain.EndpointCaseAlternativeSaveReq) (count int, err error) {
	// copy new case
	newEndpointCase, err := s.EndpointCaseService.Copy(req.BaseId, "extend-",
		req.CreateUserId, req.CreateUserName)

	s.EndpointCaseRepo.UpdateInfo(newEndpointCase.ID, map[string]interface{}{
		"case_type": consts.CaseAlternative,
	})

	// get new case's debug data
	newDebugData, err := s.DebugInterfaceService.GetDebugDataFromDebugInterface(newEndpointCase.DebugInterfaceId)
	if err != nil {
		return
	}

	// change field value by path if exist
	for _, val := range req.Values {
		if val.Category != consts.AlternativeCaseCase {
			continue
		}

		fieldIn, fieldNameOrPath := s.getFieldProps(val.Path)
		if fieldIn == "" {
			logUtils.Error("failed to getFieldProps")
			continue
		}
		s.changeFieldProps(&newDebugData, fieldIn, fieldNameOrPath, val.Sample, val.FieldType)
	}

	// update to db
	_, err = s.DebugInterfaceService.Update(newDebugData, newDebugData.DebugInterfaceId)

	return
}

func (s *EndpointCaseAlternativeService) changeFieldProps(debugData *domain.DebugData,
	fieldIn, fieldNameOrPath string, sample interface{}, fieldType casesHelper.OasFieldType) {

	if fieldIn == "[query]" {
		s.changeParams(&debugData.QueryParams, fieldNameOrPath, sample)
	} else if fieldIn == "[path]" {
		s.changeParams(&debugData.PathParams, fieldNameOrPath, sample)
	} else if fieldIn == "[header]" {
		s.changeHeaders(&debugData.Headers, fieldNameOrPath, sample)
	} else if fieldIn == "[body]/[application-json]" {
		s.changeBody(debugData, fieldNameOrPath, sample, fieldType)
	} else if fieldIn == "[body]/[multipart-form-data]" {
		s.changeForm(debugData, fieldNameOrPath, sample, fieldType)
	} else if fieldIn == "[body]/[application-x-www-form-urlencoded]" {
		s.changeFormUrlencoded(debugData, fieldNameOrPath, sample, fieldType)
	}

	return
}

func (s *EndpointCaseAlternativeService) changeParams(params *[]domain.Param,
	fieldNameOrPath string, sample interface{}) {

	found := false
	for index, param := range *params {
		if param.Name == fieldNameOrPath {
			(*params)[index].Value = fmt.Sprintf("%v", sample)

			found = true
			break
		}
	}

	if !found {
		*params = append(*params, domain.Param{
			Name:  fieldNameOrPath,
			Value: fmt.Sprintf("%v", sample),
		})
	}
}
func (s *EndpointCaseAlternativeService) changeHeaders(headers *[]domain.Header,
	fieldNameOrPath string, sample interface{}) {

	found := false
	for index, header := range *headers {
		if header.Name == fieldNameOrPath {
			(*headers)[index].Value = fmt.Sprintf("%v", sample)

			found = true
			break
		}
	}

	if !found {
		*headers = append(*headers, domain.Header{
			Name:  fieldNameOrPath,
			Value: fmt.Sprintf("%v", sample),
		})
	}
}
func (s *EndpointCaseAlternativeService) changeForm(debugData *domain.DebugData,
	fieldNameOrPath string, sample interface{}, fieldType casesHelper.OasFieldType) {
	// form_item2/[format]

	fieldPath := s.getFieldPath(fieldNameOrPath)

	found := false
	for index, item := range debugData.BodyFormData {
		if item.Name == fieldPath {
			debugData.BodyFormData[index].Value = valueUtils.InterfaceToStr(sample)
			found = true
		}
	}

	if !found {
		debugData.BodyFormData = append(debugData.BodyFormData, domain.BodyFormDataItem{
			Name:  fieldPath,
			Value: valueUtils.InterfaceToStr(sample),
		})
	}
}
func (s *EndpointCaseAlternativeService) changeFormUrlencoded(debugData *domain.DebugData,
	fieldNameOrPath string, sample interface{}, fieldType casesHelper.OasFieldType) {

	fieldPath := s.getFieldPath(fieldNameOrPath)

	found := false
	for index, item := range debugData.BodyFormUrlencoded {
		if item.Name == fieldPath {
			debugData.BodyFormUrlencoded[index].Value = valueUtils.InterfaceToStr(sample)
			found = true
		}
	}

	if !found {
		debugData.BodyFormUrlencoded = append(debugData.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
			Name:  fieldPath,
			Value: valueUtils.InterfaceToStr(sample),
		})
	}
}
func (s *EndpointCaseAlternativeService) changeBody(debugData *domain.DebugData,
	fieldNameOrPath string, sample interface{}, fieldType casesHelper.OasFieldType) {
	// id/[required]
	// id/[rule]/[min]
	// pet/age/[rule]/[min]
	// sons/[arr]/email/[format]

	fieldPath := s.getFieldPath(fieldNameOrPath)

	debugData.Body, _ = sjson.Set(debugData.Body, fieldPath, sample)
}

func (s *EndpointCaseAlternativeService) getFieldPath(pth string) (ret string) {
	retArr := []string{}

	arr := strings.Split(pth, "/")

	for _, item := range arr {
		if item == "[arr]" {
			retArr = append(retArr, "0") // change the first item
			continue
		}

		if strings.Index(item, "[") == 0 {
			break
		}

		retArr = append(retArr, item)
	}

	ret = strings.Join(retArr, ".")

	return
}

func (s *EndpointCaseAlternativeService) getFieldProps(pth string) (fieldIn string, fieldNameOrPath string) {
	arr := strings.Split(pth, "/")

	if len(arr) < 3 {
		return
	}

	if arr[0] == "[path]" {
		// [path]/path1/[format]
		fieldIn = arr[0]
		fieldNameOrPath = arr[1]

	} else if arr[0] == "[query]" {
		// [query/count/[required]
		// [query/count/[rule/[min]
		fieldIn = arr[0]
		fieldNameOrPath = arr[1]

	} else if arr[0] == "[header]" {
		// [header/header1/[format
		fieldIn = arr[0]
		fieldNameOrPath = arr[1]

	} else if arr[0] == "[body]" {
		// [body]/[application-json]/id/[required]
		// [body]/[application-json]/id/[rule]/[min]
		// [body]/[application-json]/pet/age/[rule]/[min]
		// [body]/[application-json]/sons/[arr]/email/[format]
		// [body]/[multipart-form-data]
		// [body]/[application-x-www-form-urlencoded]

		if !_stringUtils.StrInArr(arr[1],
			[]string{"[application-json]", "[multipart-form-data]", "[application-x-www-form-urlencoded]"}) {
			return
		} // ignore no-json and no-form request body

		fieldIn = strings.Join(arr[:2], "/")
		fieldNameOrPath = strings.Join(arr[2:], "/")
	}

	return
}

func (s *EndpointCaseAlternativeService) LoadCaseForExec(req agentExec.CasesExecObj) (
	ret agentExec.InterfaceExecObj, err error) {

	ret.DebugData, _ = s.LoadDebugDataForExec(req)

	// load default environment for user
	env, _ := s.EnvironmentRepo.GetByUserAndProject(req.UserId, req.ProjectId)
	if env.ID > 0 {
		ret.DebugData.ServerId = env.ID
	}

	ret.PreConditions, _ = s.PreConditionRepo.ListTo(
		ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId, req.UsedBy)
	ret.PostConditions, _ = s.PostConditionRepo.ListTo(
		ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId, req.UsedBy)

	ret.ExecScene.ShareVars = ret.DebugData.ShareVars // for execution
	ret.DebugData.ShareVars = nil                     // for display on debug page only

	// get environment and settings on project level
	s.SceneService.LoadEnvVars(&ret.ExecScene, ret.DebugData)
	s.SceneService.LoadProjectSettings(&ret.ExecScene, ret.DebugData.ProjectId)

	return
}

func (s *EndpointCaseAlternativeService) LoadDebugDataForExec(req agentExec.CasesExecObj) (
	ret domain.DebugData, err error) {

	endpointCase, err := s.EndpointCaseService.Get(req.BaseCaseId)
	if err != nil {
		return
	}

	ret, err1 := s.DebugInterfaceService.GetDebugDataFromDebugInterface(endpointCase.DebugInterfaceId)
	if err1 != nil {
		err = err1
		return
	}

	fieldIn, fieldNameOrPath := s.getFieldProps(req.Path)
	if fieldIn == "" {
		logUtils.Info("failed to getFieldProps")
		return
	}

	s.changeFieldProps(&ret, fieldIn, fieldNameOrPath, req.Sample, req.FieldType)

	return
}
