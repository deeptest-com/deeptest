package scriptHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	"path"
)

func GetScript(name ScriptType) string {
	if name == ScriptDeepTest {
		if DeepTestScript == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "deeptest.js"))
			DeepTestScript = string(bytes)
		}
		return DeepTestScript

	} else if name == DeclareDeepTest {
		if DeepTestDeclare == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "deeptest.d.ts"))
			DeepTestDeclare = string(bytes)
		}
		return DeepTestDeclare

	} else if name == DeclareDeepTestPost {
		if DeepTestDeclarePost == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "deeptest-post.d.ts"))
			DeepTestDeclarePost = string(bytes)
		}
		return DeepTestDeclarePost

	} else if name == DeclareDeepTestScenarioCustomCode {
		if DeepTestScenarioCustomCode == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "deeptest-scenario-custom-code.d.ts"))
			DeepTestScenarioCustomCode = string(bytes)
		}
		return DeepTestScenarioCustomCode

	} else if name == DeclareChai {
		if DeepTestDeclareChai == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "chai.d.ts"))
			DeepTestDeclareChai = string(bytes)
		}
		return DeepTestDeclareChai

	} else if name == ScriptMock {
		if MockScript == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "mock.js"))
			MockScript = string(bytes)
		}
		return MockScript

	} else if name == DeclareMock {
		if MockDeclare == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "mock.d.ts"))
			MockDeclare = string(bytes)
		}
		return MockDeclare

	} else if name == SnippetDatapoolGet {
		if DatapoolGetScript == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "snippet", "datapool_get.txt"))
			DatapoolGetScript = string(bytes)
		}
		return DatapoolGetScript

	} else if name == SnippetVariablesGet {
		if VariablesGet == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "snippet", "variables_get.txt"))
			VariablesGet = string(bytes)
		}
		return VariablesGet

	} else if name == SnippetVariablesSet {
		if VariablesSet == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "snippet", "variables_set.txt"))
			VariablesSet = string(bytes)
		}
		return VariablesSet

	} else if name == SnippetVariablesClear {
		if VariablesClear == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "snippet", "variables_clear.txt"))
			VariablesClear = string(bytes)
		}
		return VariablesClear

	} else if name == ScriptFuncs {
		if FuncsScript == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "funcs.js"))
			DeepTestScript = string(bytes)
		}
		return DeepTestScript

	}

	return ""
}

func GetModule(name string) (ret string) {
	bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "module", name))
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

	statusText := ""
	if po.ResultStatus == consts.Pass {
		statusText = "成功"
	} else if po.ResultStatus == consts.Fail {
		statusText = "失败"
	}

	po.ResultMsg = fmt.Sprintf("%s%s%s", name, _i118Utils.Sprintf("exec"), statusText)

	po.ResultMsg += " JSON~" + po.Output + "~JSON"

	return
}

var (
	DeepTestScript             = ""
	DeepTestDeclare            = ""
	DeepTestDeclarePost        = ""
	DeepTestScenarioCustomCode = ""
	DeepTestDeclareChai        = ""

	MockScript  = ""
	MockDeclare = ""

	JslibsDeclares = ""

	DatapoolGetScript = ""
	VariablesGet      = ""
	VariablesSet      = ""
	VariablesClear    = ""
	FuncsScript       = ""
)

type ScriptType string

const (
	ScriptDeepTest                    = "deeptest"
	DeclareDeepTest                   = "deeptest.d"
	DeclareDeepTestPost               = "deeptest-post.d"
	DeclareDeepTestScenarioCustomCode = "deeptest-scenario-custom-code.d"
	DeclareChai                       = "chai.d"

	ModuleMockJs = "mockjs.js"

	ScriptMock  = "mock"
	DeclareMock = "mock.d"

	DeclareJslibs = "jslibs"

	ScriptFuncs = "funcs"

	SnippetDatapoolGet    = "datapool_get"
	SnippetVariablesGet   = "variables_get"
	SnippetVariablesSet   = "variables_set"
	SnippetVariablesClear = "variables_clear"
)

func (e ScriptType) String() string {
	return string(e)
}
