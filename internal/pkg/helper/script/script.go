package scriptHelper

import (
	"fmt"
	"github.com/deeptest-com/deeptest"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	_i118Utils "github.com/deeptest-com/deeptest/pkg/lib/i118"
	"path"
)

func GetScript(name ScriptType) string {
	if name == ScriptDeepTest {
		if DeepTestScript == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "deeptest.js"))
			DeepTestScript = string(bytes)
		}
		return DeepTestScript

	} else if name == ScriptDeepTestSimple {
		if DeepTestScriptSimple == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "deeptest-simple.js"))
			DeepTestScriptSimple = string(bytes)
		}
		return DeepTestScriptSimple

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

	} else if name == ScriptCustom {
		if CustomScript == "" {
			bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "export", "custom.js"))
			CustomScript = string(bytes)
		}
		return CustomScript

	}

	return ""
}

func GetModule(name ScriptType) (ret string) {
	bytes, _ := deeptest.ReadResData(path.Join("res", "goja", "module", name.String()))
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
	DeepTestScriptSimple       = ""
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
	CustomScript      = ""
)

type ScriptType string

const (
	ScriptDeepTest                    ScriptType = "deeptest"
	ScriptDeepTestSimple              ScriptType = "deeptest-simple"
	DeclareDeepTest                   ScriptType = "deeptest.d"
	DeclareDeepTestPost               ScriptType = "deeptest-post.d"
	DeclareDeepTestScenarioCustomCode ScriptType = "deeptest-scenario-custom-code.d"
	DeclareChai                       ScriptType = "chai.d"

	ModuleMockJs ScriptType = "mockjs.js"

	ScriptMock  ScriptType = "mock"
	DeclareMock ScriptType = "mock.d"

	DeclareJslibs ScriptType = "jslibs"

	ScriptCustom ScriptType = "custom"

	SnippetDatapoolGet    ScriptType = "datapool_get"
	SnippetVariablesGet   ScriptType = "variables_get"
	SnippetVariablesSet   ScriptType = "variables_set"
	SnippetVariablesClear ScriptType = "variables_clear"
)

func (e ScriptType) String() string {
	return string(e)
}
