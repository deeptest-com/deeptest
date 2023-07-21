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
		nameDesc = fmt.Sprintf("状态码检查点 %s \"%s\"", optName, value)
	} else if typ == consts.ResponseHeader {
		nameDesc = fmt.Sprintf("响应头检查点 %s \"%s\"", optName, expression)
	} else if typ == consts.ResponseBody {
		nameDesc = fmt.Sprintf("响应体检查点 %s \"%s\"", optName, value)
	} else if typ == consts.Extractor {
		nameDesc = fmt.Sprintf("提取器检查点 %s %s \"%s\"", extractorVariable, optName, value)
	} else if typ == consts.Judgement {
		nameDesc = fmt.Sprintf("表达式检查点 \"%s\"", expression)
	}

	ret = nameDesc

	return
}

func GenResultMsg(po *domain.CheckpointBase) (ret string) {
	desc := GenDesc(po.Type, po.Operator, po.Value, po.Expression, po.ExtractorVariable)

	ret = fmt.Sprintf("%s， 执行%s。", desc, _i118Utils.Sprintf(po.ResultStatus.String()))

	return
}
