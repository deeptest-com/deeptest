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
		if DeepTestScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "deeptest.js"))
			DeepTestScript = string(bytes)
		}
		return DeepTestScript

	} else if name == DeclareDeepTest {
		if DeepTestDeclare == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "deeptest.d.ts"))
			DeepTestDeclare = string(bytes)
		}
		return DeepTestDeclare

	} else if name == ScriptMock {
		if MockScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "mock.js"))
			MockScript = string(bytes)
		}
		return MockScript

	} else if name == DeclareMock {
		if MockDeclare == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "export", "mock.d.ts"))
			MockDeclare = string(bytes)
		}
		return MockDeclare

	} else if name == SnippetDatapoolGet {
		if DatapoolGetScript == "" {
			bytes, _ := deeptest.ReadResData(filepath.Join("res", "goja", "snippet", "datapool_get.txt"))
			DatapoolGetScript = string(bytes)
		}
		return DatapoolGetScript

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
	DeepTestScript  = ""
	DeepTestDeclare = ""

	MockScript  = ""
	MockDeclare = ""

	JslibsDeclares = ""

	DatapoolGetScript = ""
	VariablesGet      = ""
	VariablesSet      = ""
	VariablesClear    = ""
)

type ScriptType string

const (
	ScriptDeepTest  = "deeptest"
	DeclareDeepTest = "deeptest.d"

	ScriptMock  = "mock"
	DeclareMock = "mock.d"

	DeclareJslibs = "jslibs"

	SnippetDatapoolGet    = "datapool_get"
	SnippetVariablesGet   = "variables_get"
	SnippetVariablesSet   = "variables_set"
	SnippetVariablesClear = "variables_clear"
)

func (e ScriptType) String() string {
	return string(e)
}
