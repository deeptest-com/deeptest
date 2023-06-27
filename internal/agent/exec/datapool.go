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

func getDatapoolValue(placeholder string) (ret string) {
	// _dp(name, col, 1 | seq | rand >)

	regex := regexp.MustCompile(fmt.Sprintf("(?Ui)%s\\((.+),(.+),(.+)\\)", consts.PlaceholderPrefixDatapool))
	arrs := regex.FindAllStringSubmatch(placeholder, -1)

	if !(len(arrs) == 1 && len(arrs[0]) == 4) {
		return
	}

	dpName := strings.TrimSpace(arrs[0][1])
	dpCol := strings.TrimSpace(arrs[0][2])
	dpSeq := strings.TrimSpace(arrs[0][3])

	dp := ExecScene.Datapools[dpName]
	if dp == nil {
		ret = fmt.Sprintf("${%s}", placeholder)
		return
	}

	rowIndex := getDatapoolRow(dpName, dpSeq, ExecScene.Datapools)

	val := ExecScene.Datapools[dpName][rowIndex][dpCol]
	if val == nil {
		val = "NOT_FOUND"
	}

	ret = fmt.Sprintf("%v", val)

	return
}

func getDatapoolRow(dpName, seq string, datapools domain.Datapools) (ret int) {
	dp := datapools[dpName]
	if dp == nil {
		return
	}

	total := len(dp)

	if seq == "seq" {
		ret = DatapoolCursor[dpName] % total
		DatapoolCursor[dpName]++

	} else if seq == "rand" {
		rand.Seed(time.Now().Unix())
		ret = rand.Intn(total)

	} else {
		seqInt, _ := strconv.Atoi(seq)
		ret = seqInt % total
	}

	return
}
