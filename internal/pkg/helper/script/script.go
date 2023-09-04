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

	} else if name == ScriptGlobal {
		if GlobalScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "global.txt"))
			GlobalScript = string(bytes)
		}
		return GlobalScript

	} else if name == ScriptModule {
		if ModuleScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "module.txt"))
			ModuleScript = string(bytes)
		}
		return ModuleScript

	} else if name == ScriptDatapoolGet {
		if ModuleScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "datapool_get.txt"))
			ModuleScript = string(bytes)
		}
		return ModuleScript

	} else if name == ScriptVariablesGet {
		if VariablesGet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_get.txt"))
			VariablesGet = string(bytes)
		}
		return VariablesGet

	} else if name == ScriptVariablesSet {
		if VariablesSet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_set.txt"))
			VariablesSet = string(bytes)
		}
		return VariablesSet

	} else if name == ScriptVariablesClear {
		if VariablesClear == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_clear.txt"))
			VariablesClear = string(bytes)
		}
		return VariablesClear

	}

	//else if name == ScriptEnvironmentGet {
	//	if EnvironmentGet == "" {
	//		bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "environment_get.txt"))
	//		EnvironmentGet = string(bytes)
	//	}
	//	return EnvironmentGet
	//
	//} else if name == ScriptEnvironmentSet {
	//	if EnvironmentSet == "" {
	//		bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "environment_set.txt"))
	//		EnvironmentSet = string(bytes)
	//	}
	//	return EnvironmentSet
	//
	//} else if name == ScriptEnvironmentClear {
	//	if EnvironmentClear == "" {
	//		bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "environment_clear.txt"))
	//		EnvironmentClear = string(bytes)
	//	}
	//	return EnvironmentClear
	//
	//}

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

	po.ResultMsg = fmt.Sprintf("%s%s%s，输出\"%s\"。", name,
		_i118Utils.Sprintf("exec"), _i118Utils.Sprintf(po.ResultStatus.String()),
		po.Output)

	return
}

var (
	DpScript = ""

	GlobalScript = ""
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
	ScriptGlobal   = "global"
	ScriptModule   = "module"

	ScriptDatapoolGet = "datapool_get"

	//ScriptEnvironmentGet   = "environment_get"
	//ScriptEnvironmentSet   = "environment_set"
	//ScriptEnvironmentClear = "environment_clear"

	ScriptVariablesGet   = "variables_get"
	ScriptVariablesSet   = "variables_set"
	ScriptVariablesClear = "variables_clear"
)

func (e ScriptType) String() string {
	return string(e)
}
