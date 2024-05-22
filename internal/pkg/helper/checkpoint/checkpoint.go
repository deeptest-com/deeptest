package checkpointHelpper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
)

func GenDesc(typ consts.CheckpointType, operator consts.ComparisonOperator, actualResult, expression,
	extractorVariable string, extractorType consts.ExtractorType, extractorExpression string) (ret string) {
	nameDesc := ""

	opt := fmt.Sprintf("%v", operator)
	optName := _i118Utils.Sprintf(opt)
	if typ == consts.ResponseStatus {
		nameDesc = _i118Utils.Sprintf("usage")
		nameDesc = fmt.Sprintf("状态码%s\"%s\"", optName, actualResult)
	} else if typ == consts.ResponseHeader {
		nameDesc = fmt.Sprintf("响应头%s%s\"%s\"", expression, optName, actualResult)
	} else if typ == consts.ResponseBody {
		nameDesc = fmt.Sprintf("响应体%s\"%s\"", optName, actualResult)
	} else if typ == consts.ExtractorVari {
		nameDesc = fmt.Sprintf("提取变量%s%s\"%s\"", extractorVariable, optName, actualResult)
	} else if typ == consts.Extractor {
		extractorDesc := extractorHelper.GenDescForCheckpoint(extractorType, extractorExpression)
		nameDesc = fmt.Sprintf("提取%s%s\"%s\"", extractorDesc, optName, actualResult)
	} else if typ == consts.Judgement {
		nameDesc = fmt.Sprintf("表达式\"%s\"", expression)
	}

	ret = nameDesc

	return
}

func GenResultMsg(po *domain.CheckpointBase) {
	desc := GenDesc(po.Type, po.Operator, po.ActualResult, po.Expression, po.ExtractorVariable, po.ExtractorType, po.ExtractorExpression)

	po.ResultMsg = fmt.Sprintf("%s", desc)

	if po.ResultStatus != consts.Pass {
		po.ResultMsg += fmt.Sprintf("，实际结果\"%s\"", po.ActualResult)
	}

	return
}
