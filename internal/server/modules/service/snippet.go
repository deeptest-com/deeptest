package service

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	mockHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/mock"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/snowlyg/helper/dir"
	"path/filepath"
	"time"
)

var (
	JslibsDeclares []string
)

type SnippetService struct {
	SnippetRepo           *repo.SnippetRepo      `inject:""`
	JslibRepo             *repo.JslibRepo        `inject:""`
	MockJsService         *MockJsService         `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`
}

func (s *SnippetService) ListJslibNames(tenantId consts.TenantId, projectId int) (names []string, err error) {
	libs, _ := s.JslibRepo.List(tenantId, "", projectId, true)

	for _, po := range libs {
		names = append(names, po.Name)
	}

	return
}

func (s *SnippetService) Get(name scriptHelper.ScriptType) (po jslibHelper.Jslib, err error) {
	script := scriptHelper.GetScript(name)

	po = jslibHelper.Jslib{
		Script: script,
	}
	return
}

func (s *SnippetService) GetJslibs(tenantId consts.TenantId, projectId int) (pos []jslibHelper.Jslib, err error) {
	//if JslibsDeclares == nil {

	JslibsDeclares = nil
	libs, _ := s.JslibRepo.List(tenantId, "", projectId, true)

	for _, lib := range libs {
		pth := filepath.Join(dir.GetCurrentAbPath(), lib.TypesFile)
		content := fileUtils.ReadFile(pth)

		JslibsDeclares = append(JslibsDeclares, content)
	}
	//}

	for _, item := range JslibsDeclares {
		po := jslibHelper.Jslib{
			Script: item,
		}
		pos = append(pos, po)
	}

	return
}

func (s *SnippetService) GetJslibsForAgent(tenantId consts.TenantId, loadedLibs map[uint]time.Time, projectId int) (tos []jslibHelper.Jslib, err error) {
	pos, _ := s.JslibRepo.List(tenantId, "", projectId, true)

	for _, po := range pos {
		pth := filepath.Join(consts.WorkDir, po.ScriptFile)
		content := fileUtils.ReadFile(pth)

		updateTime := po.UpdatedAt
		if updateTime == nil {
			updateTime = po.CreatedAt
		}

		loadTime, found := loadedLibs[po.ID]
		if !found || loadTime.Before(*updateTime) {
			to := jslibHelper.Jslib{
				Id:        po.ID,
				Name:      po.Name,
				Script:    content,
				UpdatedAt: *updateTime,
			}
			tos = append(tos, to)
		}
	}

	return
}

func (s *SnippetService) ListVar(tenantId consts.TenantId, req domain.DebugInfo) (res []serverDomain.SnippetRes) {
	data, err := s.DebugInterfaceService.Load(tenantId, req)
	if err != nil {
		return
	}

	for _, item := range data.EnvDataToView.EnvVars {
		expression := fmt.Sprintf("${%s}", item.Name)
		res = append(res, serverDomain.SnippetRes{Label: item.Name, Value: expression, Desc: "环境变量"})
	}

	for _, item := range data.EnvDataToView.ShareVars {
		expression := fmt.Sprintf("${%s}", item.Name)
		res = append(res, serverDomain.SnippetRes{Label: item.Name, Value: expression, Desc: "共享变量"})
	}

	for _, item := range data.EnvDataToView.GlobalVars {
		expression := fmt.Sprintf("${%s}", item.Name)
		res = append(res, serverDomain.SnippetRes{Label: item.Name, Value: expression, Desc: "全局变量"})
	}

	return
}

func (s *SnippetService) ListMock(tenantId consts.TenantId) (res []serverDomain.SnippetRes) {
	list, _ := s.MockJsService.ListExpressions(tenantId)
	for _, item := range list {
		expression := fmt.Sprintf("${_mock(\"@%s\")}", item.Expression)
		res = append(res, serverDomain.SnippetRes{Label: "@" + item.Expression, Value: expression, Desc: item.Name})
	}

	return
}

func (s *SnippetService) ListSysFunc() (res []serverDomain.SnippetRes) {

	for _, item := range agentExec.SysFuncList {
		expression := fmt.Sprintf("${%s}", item.Value)
		res = append(res, serverDomain.SnippetRes{Label: item.Label, Value: expression, Desc: item.Desc})
	}

	return
}

func (s *SnippetService) ListCustomFunc(tenantId consts.TenantId, projectId uint) (res []serverDomain.SnippetRes) {
	res = make([]serverDomain.SnippetRes, 0)
	mockHelper.InitJsRuntime(tenantId, projectId)

	libs, _ := s.JslibRepo.List(tenantId, "", int(projectId), true)
	for _, item := range libs {
		jslib := jslibHelper.GetJslibCache(tenantId, item.ID)
		for _, function := range jslib.Functions {
			functionName := fmt.Sprintf("%s(%s)", function.Name, function.Args)
			expression := fmt.Sprintf("${%s.%s}", item.Name, functionName)
			res = append(res, serverDomain.SnippetRes{Label: functionName, Value: expression, Desc: item.Name})
		}
	}

	return res
}
