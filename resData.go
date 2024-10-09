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

func ReadPromptTempl(pth string) (ret string, err error) {
	var bytes []byte

	if commonUtils.IsRelease() {
		bytes, err = promptFileSys.ReadFile(pth)
	} else {
		bytes, err = os.ReadFile(filepath.Join("internal", "agent", "_prompt_templ", pth))
	}

	ret = string(bytes)

	return
}
