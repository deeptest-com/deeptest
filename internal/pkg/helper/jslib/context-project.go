package jslibHelper

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"sync"
	"time"
)

var (
	ProjectContextStore sync.Map
)

func InitProjectContext(projectId uint) {
	val := ProjectContext{
		GojaRuntime:     nil,
		GojaRequire:     nil,
		AgentLoadedLibs: nil,
	}

	ProjectContextStore.Store(projectId, &val)
}

func ClearProjectContext(projectId uint) {
	ProjectContextStore.Store(projectId, nil)
}

func GetProjectContext(projectId uint) (val *ProjectContext) {
	inf, ok := ProjectContextStore.Load(projectId)
	if !ok {
		InitProjectContext(projectId)
	}

	inf, _ = ProjectContextStore.Load(projectId)
	val = inf.(*ProjectContext)

	return
}

func InitGojaRuntime(projectId uint) (execRuntime *goja.Runtime, execRequire *require.RequireModule) {
	projectContext := GetProjectContext(projectId)

	projectContext.GojaRuntime = goja.New()
	projectContext.GojaRuntime.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	registry := new(require.Registry) // registry 能夠被多个goja.Runtime共用
	execRequire = registry.Enable(projectContext.GojaRuntime)

	projectContext.GojaRequire = execRequire

	execRuntime = projectContext.GojaRuntime

	return
}

func GetAgentLoadedLibs(projectId uint) (agentLoadedLibs *map[uint]time.Time) {
	projectContext := GetProjectContext(projectId)
	agentLoadedLibs = projectContext.AgentLoadedLibs

	if agentLoadedLibs == nil {
		agentLoadedLibs = &map[uint]time.Time{}
		SetAgentLoadedLibs(projectId, agentLoadedLibs)
	}

	return
}
func SetAgentLoadedLibs(projectId uint, agentLoadedLibs *map[uint]time.Time) {
	projectContext := GetProjectContext(projectId)
	projectContext.AgentLoadedLibs = agentLoadedLibs

	return
}

type ProjectContext struct {
	// for goja js engine
	GojaRuntime     *goja.Runtime
	GojaRequire     *require.RequireModule
	AgentLoadedLibs *map[uint]time.Time
}
