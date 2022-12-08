package action

import (
	"fmt"
	"github.com/aaronchen2k/deeptest"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/fatih/color"
	"path"
)

var ()

func PrintUsage(lang string) {
	usageFile := path.Join("res", lang, "usage.txt")

	logUtils.Info("\n" + color.CyanString(_i118Utils.Sprintf("usage")))

	usageData, _ := deeptest.ReadResData(usageFile)
	exeFile := consts.App
	if commonUtils.IsWin() {
		exeFile += ".exe"
	}
	usage := fmt.Sprintf(string(usageData), exeFile)
	fmt.Printf("%s", usage)
}
