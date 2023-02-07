package scriptHelper

import (
	"github.com/aaronchen2k/deeptest"
	"path/filepath"
)

func GetScript(name ScriptType) string {
	if name == ScriptDp {
		if DpScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "dp.js"))
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
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "datapool_get.txt"))
			ModuleScript = string(bytes)
		}
		return ModuleScript

	} else if name == ScriptEnvironmentGet {
		if EnvironmentGet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "environment_get"))
			EnvironmentGet = string(bytes)
		}
		return EnvironmentGet

	} else if name == ScriptEnvironmentSet {
		if EnvironmentSet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "environment_set"))
			EnvironmentSet = string(bytes)
		}
		return EnvironmentSet

	} else if name == ScriptEnvironmentClear {
		if EnvironmentClear == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "environment_clear"))
			EnvironmentClear = string(bytes)
		}
		return EnvironmentClear

	} else if name == ScriptVariablesGet {
		if VariablesGet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_get"))
			VariablesGet = string(bytes)
		}
		return VariablesGet

	} else if name == ScriptVariablesSet {
		if VariablesSet == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_set"))
			VariablesSet = string(bytes)
		}
		return VariablesSet

	} else if name == ScriptVariablesClear {
		if VariablesClear == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "variables_clear"))
			VariablesClear = string(bytes)
		}
		return VariablesClear

	}

	return ""
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
	ScriptDp     = "dp"
	ScriptGlobal = "global"
	ScriptModule = "module"

	ScriptDatapoolGet = "datapool_get"

	ScriptEnvironmentGet   = "environment_get"
	ScriptEnvironmentSet   = "environment_set"
	ScriptEnvironmentClear = "environment_clear"

	ScriptVariablesGet   = "variables_get"
	ScriptVariablesSet   = "variables_set"
	ScriptVariablesClear = "variables_clear"
)

func (e ScriptType) String() string {
	return string(e)
}
