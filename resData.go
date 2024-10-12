package deeptest

import (
	"embed"
	commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	_logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
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
		currentPath, _ := os.Getwd()

		fullPath := filepath.Join(currentPath, "internal", "agent", "_prompt_templ", pth)
		bytes, err = os.ReadFile(fullPath)
	}

	if err != nil {
		_logUtils.Infof(err.Error())
	}

	ret = string(bytes)

	return
}
