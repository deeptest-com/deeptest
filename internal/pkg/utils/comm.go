package commUtils

import (
	_consts "github.com/aaronchen2k/deeptest/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"os"
	"path/filepath"
	"strings"
)

func GetExecDir() (dir string) { // where ztf exe file in
	exeDir, _ := os.Executable()

	if commonUtils.IsRelease() { // release
		dir = filepath.Dir(exeDir)
	} else { // debug mode
		if strings.Index(strings.ToLower(exeDir), "goland") > -1 { // run with ide
			dir = os.Getenv("ZTF_CODE_DIR")
		} else {
			dir = GetWorkDir()
		}
	}

	dir, _ = filepath.Abs(dir)
	dir = AddFilePathSepIfNeeded(dir)

	return
}

func GetWorkDir() string { // where we run file in
	dir, _ := os.Getwd()

	dir, _ = filepath.Abs(dir)
	dir = AddFilePathSepIfNeeded(dir)

	return dir
}

func AddFilePathSepIfNeeded(pth string) string {
	sep := _consts.FilePthSep

	if strings.LastIndex(pth, sep) < len(pth)-1 {
		pth += sep
	}
	return pth
}
