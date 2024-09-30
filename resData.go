package deeptest

import (
	"embed"
	commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
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
