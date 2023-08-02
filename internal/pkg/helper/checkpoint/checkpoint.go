package checkpointHelpper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
)

func GenDesc(typ consts.CheckpointType, operator consts.ComparisonOperator, value, expression, extractorVariable string) (ret string) {
	nameDesc := ""

	opt := fmt.Sprintf("%v", operator)
	optName := _i118Utils.Sprintf(opt)
	if typ == consts.ResponseStatus {
		nameDesc = _i118Utils.Sprintf("usage")
		nameDesc = fmt.Sprintf("状态码%s \"%s\"", optName, value)
	} else if typ == consts.ResponseHeader {
		nameDesc = fmt.Sprintf("响应头%s \"%s\"", optName, expression)
	} else if typ == consts.ResponseBody {
		nameDesc = fmt.Sprintf("响应体%s \"%s\"", optName, value)
	} else if typ == consts.Extractor {
		nameDesc = fmt.Sprintf("提取器%s %s \"%s\"", extractorVariable, optName, value)
	} else if typ == consts.Judgement {
		nameDesc = fmt.Sprintf("表达式\"%s\"", expression)
	}

	ret = "断言 " + nameDesc

	return
}

func GenResultMsg(po *domain.CheckpointBase) {
	desc := GenDesc(po.Type, po.Operator, po.Value, po.Expression, po.ExtractorVariable)

	po.ResultMsg = fmt.Sprintf("%s", desc)

	return
}
