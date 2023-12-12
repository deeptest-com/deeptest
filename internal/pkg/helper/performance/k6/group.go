package k6Converter

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	k6Comm "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/comm"
	"os"
)

func GroupToCode(group agentExec.Processor) (script string, jslibs []string, err error) {
	childrenContent, err := genChildrenContent(group.Children)

	data := map[string]string{
		"name": group.Name,
		"slot": childrenContent,
	}

	content := k6Comm.GetTmpl("group")
	script = os.Expand(content, func(k string) string { return data[k] })

	return
}
