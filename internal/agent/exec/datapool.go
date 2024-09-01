package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"math/rand"
	"strconv"
	"time"
)

func getDatapoolValue(dpName, dpCol string, dpSeq interface{}, session *ExecSession) (ret string) {
	execScene := session.ExecScene
	// _dp(name, col, 1 | seq | rand >)

	dp := execScene.Datapools[dpName]
	if dp == nil {
		ret = fmt.Sprintf("datapoll ${%s} no found", dpName)
		return
	}

	rowIndex := getDatapoolRow(dpName, dpSeq, execScene.Datapools, session.ScenarioDebug.DatapoolCursor)

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

func getDatapoolRow(dpName string, seq interface{}, datapools domain.Datapools, datapoolCursor map[string]int) (ret int) {
	dp := datapools[dpName]
	if dp == nil {
		return
	}

	total := len(dp)

	seqStr := _stringUtils.InterfToStr(seq)

	if seq == "seq" {
		ret = datapoolCursor[dpName] % total
		datapoolCursor[dpName]++

	} else if seq == "rand" {
		rand.Seed(time.Now().Unix())
		ret = rand.Intn(total)

	} else {
		seqInt, _ := strconv.Atoi(seqStr)
		ret = seqInt - 1
		ret = ret % total

		if ret < 0 {
			ret = 0
		}
	}

	return
}
