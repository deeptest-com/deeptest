package deeptest

import (
	"embed"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"io/fs"
	"os"
)

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
