package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	cookieHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cookie"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorInterface struct {
	ID uint `json:"id"`
	ProcessorEntityBase

	domain.BaseRequest
	Response domain.DebugResponse `json:"response"`

	BaseUrl string `json:"baseUrl"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`
}

func (entity ProcessorInterface) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("interface entity")
	CurrDebugInterfaceId = processor.EntityId

	startTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                  int(entity.ProcessorID),
		Name:                entity.Name,
		ProcessorCategory:   entity.ProcessorCategory,
		ProcessorType:       entity.ProcessorType,
		StartTime:           &startTime,
		ParentId:            int(entity.ParentID),
		EndpointInterfaceId: processor.EndpointInterfaceId,
		DebugInterfaceId:    processor.EntityId,
		ProcessorId:         processor.ID,
		ScenarioId:          processor.ScenarioId,
		LogId:               uuid.NewV4(),
		ParentLogId:         processor.Parent.Result.LogId,
	}

	//在循环过程中，processor 被执行多次，变量替换会受到影响，第一次跌替换之后，就不能根据实际情况替换了
	var baseRequest domain.BaseRequest
	copier.CopyWithOption(&baseRequest, &entity.BaseRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// exec pre-condition
	entity.ExecPreConditions(processor, session)

	// dealwith variables
	ReplaceVariables(&baseRequest, consts.ScenarioDebug)

	// add cookies
	DealwithCookies(&baseRequest, entity.ProcessorID)

	// gen request url
	GenRequestUrlWithBaseUrlAndPathParam(&baseRequest, processor.EntityId, entity.BaseUrl)

	// send request
	//startTime := time.UnixNano()
	entity.Response, err = Invoke(&baseRequest)

	processor.Result.Cost = time.Now().UnixMilli() - startTime.UnixMilli()
	reqContent, _ := json.Marshal(baseRequest)
	processor.Result.ReqContent = string(reqContent)
	respContent, _ := json.Marshal(entity.Response)
	processor.Result.RespContent = string(respContent)

	if err != nil {
		processor.Result.ResultStatus = consts.Fail
		processor.Result.Summary = err.Error()
		processor.AddResultToParent()
		execUtils.SendErrorMsg(*processor.Result, session.WsMsg)
		return
	}

	// exec post-condition
	entity.ExecPostConditions(processor, session)

	for _, c := range entity.Response.Cookies {
		SetCookie(processor.ParentId, c.Name, c.Value, c.Domain, c.ExpireTime)
	}

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

func (entity *ProcessorInterface) ExecPreConditions(processor *Processor, session *Session) (err error) {
	for _, condition := range entity.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)

			err = ExecScript(&scriptBase)
			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			processor.Result.ScriptsResult = append(processor.Result.ScriptsResult, scriptBase)
		}
	}

	return
}
func (entity *ProcessorInterface) ExecPostConditions(processor *Processor, session *Session) (err error) {
	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)

			if extractorBase.Disabled || extractorBase.Variable == "" {
				continue
			}

			brother, ok := getPreviousBrother(*processor)
			if !ok || brother.EntityType != consts.ProcessorInterfaceDefault {
				processor.Result.Summary = fmt.Sprintf("先前节点不是接口，无法应用提取器。")
				processor.AddResultToParent()
				execUtils.SendExecMsg(*processor.Result, session.WsMsg)
				return
			}

			resp := domain.DebugResponse{}
			json.Unmarshal([]byte(brother.Result.RespContent), &resp)

			err = ExecExtract(&extractorBase, resp)
			extractorHelper.GenResultMsg(&extractorBase)

			if extractorBase.ResultStatus == consts.Pass {
				SetVariable(processor.ID, extractorBase.Variable, extractorBase.Result, consts.Public)
			}

			processor.Result.ExtractorsResult = append(processor.Result.ExtractorsResult, extractorBase)

		} else if condition.Type == consts.ConditionTypeCookie {
			var cookieBase domain.CookieBase
			json.Unmarshal(condition.Raw, &cookieBase)

			if cookieBase.Disabled {
				continue
			}

			brother, ok := getPreviousBrother(*processor)
			if !ok || brother.EntityType != consts.ProcessorInterfaceDefault {
				processor.Result.Summary = fmt.Sprintf("先前节点不是接口，无法应用提取器。")
				processor.AddResultToParent()
				execUtils.SendExecMsg(*processor.Result, session.WsMsg)
				return
			}

			resp := domain.DebugResponse{}
			json.Unmarshal([]byte(brother.Result.RespContent), &resp)

			err = ExecCookie(&cookieBase, resp)
			cookieHelper.GenResultMsg(&cookieBase)

			if cookieBase.ResultStatus == consts.Pass {
				SetVariable(processor.ParentId, cookieBase.VariableName, cookieBase.Result, consts.Public)
			}

			processor.Result.CookiesResult = append(processor.Result.CookiesResult, cookieBase)

		} else if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			if scriptBase.Disabled {
				continue
			}

			err = ExecScript(&scriptBase)
			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			processor.Result.ScriptsResult = append(processor.Result.ScriptsResult, scriptBase)
		}
	}

	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeCheckpoint {
			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)
			if checkpointBase.Disabled {
				continue
			}

			brother, ok := getPreviousBrother(*processor)
			if !ok || brother.EntityType != consts.ProcessorInterfaceDefault {
				processor.Result.Summary = fmt.Sprintf("先前节点不是接口，无法应用提取器。")
				processor.AddResultToParent()
				execUtils.SendExecMsg(*processor.Result, session.WsMsg)
				return
			}

			resp := domain.DebugResponse{}
			json.Unmarshal([]byte(brother.Result.RespContent), &resp)

			err = ExecCheckPoint(&checkpointBase, resp, processor.ID)
			checkpointHelper.GenResultMsg(&checkpointBase)

			processor.Result.CheckpointsResult = append(processor.Result.CheckpointsResult, checkpointBase)

		}
	}

	return
}
