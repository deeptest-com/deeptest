package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getDatapoolValue(placeholder string, execUuid string) (ret string) {
	execScene := GetExecScene(execUuid)
	// _dp(name, col, 1 | seq | rand >)

	regex := regexp.MustCompile(fmt.Sprintf("(?Ui)%s\\((.+),(.+),(.+)\\)", consts.PlaceholderPrefixDatapool))
	arrs := regex.FindAllStringSubmatch(placeholder, -1)

	if !(len(arrs) == 1 && len(arrs[0]) == 4) {
		return
	}

	dpName := strings.TrimSpace(arrs[0][1])
	dpCol := strings.TrimSpace(arrs[0][2])
	dpSeq := strings.TrimSpace(arrs[0][3])

	dp := execScene.Datapools[dpName]
	if dp == nil {
		ret = fmt.Sprintf("${%s}", placeholder)
		return
	}

	rowIndex := getDatapoolRow(dpName, dpSeq, execScene.Datapools, execUuid)

	if rowIndex > len(execScene.Datapools[dpName])-1 {
		ret = "OUT_OF_RANGE"
		return
	}

	val := execScene.Datapools[dpName][rowIndex][dpCol]
	if val == nil {
		val = "NOT_FOUND"
	}

	ret = fmt.Sprintf("%v", val)

	return
}

func getDatapoolRow(dpName, seq string, datapools domain.Datapools, execUuid string) (ret int) {
	datapoolCursor := GetDatapoolCursor(execUuid)

	dp := datapools[dpName]
	if dp == nil {
		return
	}

	total := len(dp)

	if seq == "seq" {
		ret = datapoolCursor[dpName] % total
		datapoolCursor[dpName]++

	} else if seq == "rand" {
		rand.Seed(time.Now().Unix())
		ret = rand.Intn(total)

	} else {
		seqInt, _ := strconv.Atoi(seq)
		ret = seqInt - 1
		ret = ret % total

		if ret < 0 {
			ret = 0
		}
	}

	return
}
