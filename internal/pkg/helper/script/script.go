package scriptHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	"path/filepath"
)

func GetScript(name ScriptType) string {
	if name == ScriptDeepTest {
		if DpScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "deeptest.js"))
			DpScript = string(bytes)
		}
		return DpScript

	} else if name == DeclareGlobal {
		if GlobalScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "global.d.ts"))
			GlobalScript = string(bytes)
		}
		return GlobalScript

	} else if name == ScriptMock {
		if MockScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "mock.js"))
			MockScript = string(bytes)
		}
		return MockScript

	} else if name == DeclareMock {
		if MockScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "mock.d.ts"))
			MockScript = string(bytes)
		}
		return MockScript

	} else if name == DeclareModule {
		if ModuleScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "module.d.ts"))
			ModuleScript = string(bytes)
		}
		return ModuleScript

	} else if name == SnippetDatapoolGet {
		if ModuleScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "datapool_get.txt"))
			ModuleScript = string(bytes)
		}
		return ModuleScript

	} else if name == SnippetVariablesGet {
		if VariablesGet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_get.txt"))
			VariablesGet = string(bytes)
		}
		return VariablesGet

	} else if name == SnippetVariablesSet {
		if VariablesSet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_set.txt"))
			VariablesSet = string(bytes)
		}
		return VariablesSet

	} else if name == SnippetVariablesClear {
		if VariablesClear == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_clear.txt"))
			VariablesClear = string(bytes)
		}
		return VariablesClear

	}

	return ""
}

func GetModule(name string) (ret string) {
	bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "module", name))
	ret = string(bytes)

	return
}

func GenResultMsg(po *domain.ScriptBase) {
	name := "脚本"
	if po.ConditionSrc == consts.ConditionSrcPre {
		name = "预请求" + name
	} else if po.ConditionSrc == consts.ConditionSrcPost {
		name = "后处理" + name
	}

	po.ResultMsg = fmt.Sprintf("%s%s%s，输出%s。", name,
		_i118Utils.Sprintf("exec"), _i118Utils.Sprintf(po.ResultStatus.String()),
		po.Output)

	return
}

var (
	DpScript = ""

	GlobalScript = ""
	MockScript   = ""
	ModuleScript = ""

	EnvironmentGet   = ""
	EnvironmentSet   = ""
	EnvironmentClear = ""

	VariablesGet   = ""
	VariablesSet   = ""
	VariablesClear = ""
)

type ScriptType string

const (
	ScriptDeepTest = "deeptest"
	DeclareGlobal  = "global"
	DeclareModule  = "modules"

	ScriptMock  = "mock"
	DeclareMock = "mock"

	SnippetDatapoolGet    = "datapool_get"
	SnippetVariablesGet   = "variables_get"
	SnippetVariablesSet   = "variables_set"
	SnippetVariablesClear = "variables_clear"
)

func (e ScriptType) String() string {
	return string(e)
}
