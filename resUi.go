package deeptest

import (
	"embed"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"io/fs"
	"os"
)

var uiFileSys embed.FS

func GetUiFileSys() (ret fs.FS, err error) {
	//if commonUtils.IsRelease() {
	//	ret, err = fs.Sub(uiFileSys, "ui/dist")
	//} else {
	fileUtils.MkDirIfNeeded("ui/dist")
	ret = os.DirFS("ui/dist")
	//}

	return
}
