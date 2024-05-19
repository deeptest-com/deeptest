package jslibHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"sync"
	"time"
)

var (
	ProjectContextStore sync.Map
)

func InitProjectContext(tenantId consts.TenantId, projectId uint) {
	val := ProjectContext{
		GojaRuntime:     nil,
		GojaRequire:     nil,
		AgentLoadedLibs: nil,
	}
	key := getKey(tenantId, projectId)
	ProjectContextStore.Store(key, &val)
}

func ClearProjectContext(tenantId consts.TenantId, projectId uint) {
	key := getKey(tenantId, projectId)
	ProjectContextStore.Store(key, nil)
}

func GetProjectContext(tenantId consts.TenantId, projectId uint) (val *ProjectContext) {
	key := getKey(tenantId, projectId)
	inf, ok := ProjectContextStore.Load(key)
	if !ok {
		InitProjectContext(tenantId, projectId)
	}

	inf, _ = ProjectContextStore.Load(key)
	val = inf.(*ProjectContext)

	return
}

func InitProjectGojaRuntime(tenantId consts.TenantId, projectId uint) (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	projectContext := GetProjectContext(tenantId, projectId)

	projectContext.GojaRuntime = goja.New()
	projectContext.GojaRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	execRequire = registry.Enable(projectContext.GojaRuntime)

	projectContext.GojaRequire = execRequire

	execRuntime = projectContext.GojaRuntime

	return
}

func GetProjectGojaRuntime(tenantId consts.TenantId, projectId uint) (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	projectContext := GetProjectContext(tenantId, projectId)
	execRuntime = projectContext.GojaRuntime
	execRequire = projectContext.GojaRequire

	return
}

func GetAgentLoadedLibs(tenantId consts.TenantId, projectId uint) (agentLoadedLibs *map[uint]time.Time) {
	projectContext := GetProjectContext(tenantId, projectId)
	agentLoadedLibs = projectContext.AgentLoadedLibs

	if agentLoadedLibs == nil {
		agentLoadedLibs = &map[uint]time.Time{}
		SetAgentLoadedLibs(tenantId, projectId, agentLoadedLibs)
	}

	return
}
func SetAgentLoadedLibs(tenantId consts.TenantId, projectId uint, agentLoadedLibs *map[uint]time.Time) {
	projectContext := GetProjectContext(tenantId, projectId)
	projectContext.AgentLoadedLibs = agentLoadedLibs

	return
}

type ProjectContext struct {
	// for goja js engine
	GojaRuntime     *goja.Runtime
	GojaRequire     *require.RequireModule
	AgentLoadedLibs *map[uint]time.Time
}

func getKey(tenantId consts.TenantId, projectId uint) string {
	return fmt.Sprintf("%v-%d", tenantId, projectId)
}
