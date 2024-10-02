package deeptest

import (
	"embed"
	commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	"os"
	"path/filepath"
)

//go:embed res
var resFileSys embed.FS

func ReadResData(pth string) (ret []byte, err error) {
	if commonUtils.IsRelease() {
		ret, err = resFileSys.ReadFile(pth)
	} else {
		ret, err = os.ReadFile(pth)
	}

	return
}

//go:embed internal/agent/_prompt_templ
var promptFileSys embed.FS

func ReadPromptTempl(pth string) (ret []byte, err error) {
	if commonUtils.IsRelease() {
		ret, err = promptFileSys.ReadFile(pth)
	} else {
		ret, err = os.ReadFile(filepath.Join("internal", "agent", "_prompt_templ", pth))
	}

	return
}
