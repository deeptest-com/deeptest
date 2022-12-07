package deeptest

import (
	"embed"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"io/fs"
	"os"
)

//go:embed res
var resFileSys embed.FS

func ReadResData(path string) (ret []byte, err error) {
	if commonUtils.IsRelease() {
		ret, err = resFileSys.ReadFile(path)
	} else {
		ret, err = os.ReadFile(path)
	}

	return
}

func GetResFileSys() (ret fs.FS, err error) {
	if commonUtils.IsRelease() {
		ret, err = fs.Sub(uiFileSys, "res")
	} else {
		ret = os.DirFS("res")
	}

	return
}

//go:embed ui/dist
var uiFileSys embed.FS

func GetUiFileSys() (ret fs.FS, err error) {
	if commonUtils.IsRelease() {
		ret, err = fs.Sub(uiFileSys, "ui/dist")
	} else {
		ret = os.DirFS("ui/dist")
	}

	return
}
