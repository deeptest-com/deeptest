package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
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
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("interface entity")
	CurrDebugInterfaceId = processor.EntityId

	execStartTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                  int(entity.ProcessorID),
		Name:                processor.Name,
		ProcessorCategory:   entity.ProcessorCategory,
		ProcessorType:       entity.ProcessorType,
		StartTime:           &execStartTime,
		ParentId:            int(entity.ParentID),
		EndpointInterfaceId: processor.EndpointInterfaceId,
		DebugInterfaceId:    processor.EntityId,
		ProcessorId:         processor.ID,
		ScenarioId:          processor.ScenarioId,
		LogId:               uuid.NewV4(),
		ParentLogId:         processor.Parent.Result.LogId,
		Round:               processor.Round,
	}

	detail := map[string]interface{}{}

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
	requestStartTime := time.Now()
	entity.Response, err = Invoke(&baseRequest)
	requestEndTime := time.Now()

	processor.Result.Cost = requestEndTime.UnixMilli() - requestStartTime.UnixMilli()
	reqContent, _ := json.Marshal(baseRequest)
	processor.Result.ReqContent = string(reqContent)
	respContent, _ := json.Marshal(entity.Response)
	processor.Result.RespContent = string(respContent)
	processor.Result.ResultStatus = consts.Pass
	if err != nil {
		processor.Result.ResultStatus = consts.Fail
		processor.Result.Summary = err.Error()
		detail["result"] = entity.Response.Content
		processor.Result.Detail = commonUtils.JsonEncode(detail)
		execUtils.SendErrorMsg(*processor.Result, session.WsMsg)
		processor.AddResultToParent()
		return
	}

	// exec post-condition
	entity.ExecPostConditions(processor, detail)
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	for _, c := range entity.Response.Cookies {
		SetCookie(processor.ParentId, c.Name, c.Value, c.Domain, c.ExpireTime)
	}

	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	stat := CountStat(processor.Result)
	execUtils.SendStatMsg(stat, session.WsMsg)
	processor.AddResultToParent()

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

			interfaceExecCondition := domain.InterfaceExecCondition{
				Type: condition.Type,
			}
			interfaceExecCondition.Raw, _ = json.Marshal(scriptBase)
			processor.Result.PreConditions = append(processor.Result.PreConditions, interfaceExecCondition)
		}
	}

	return
}
func (entity *ProcessorInterface) ExecPostConditions(processor *Processor, detail map[string]interface{}) (err error) {

	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)

			if extractorBase.Disabled || extractorBase.Variable == "" {
				continue
			}

			resp := entity.Response

			err = ExecExtract(&extractorBase, resp)
			extractorHelper.GenResultMsg(&extractorBase)

			if extractorBase.ResultStatus == consts.Pass {
				SetVariable(processor.ParentId, extractorBase.Variable, extractorBase.Result, consts.Public)
			}

			interfaceExecCondition := domain.InterfaceExecCondition{
				Type: condition.Type,
			}
			interfaceExecCondition.Raw, _ = json.Marshal(extractorBase)
			processor.Result.PostConditions = append(processor.Result.PostConditions, interfaceExecCondition)

		} else if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			if scriptBase.Disabled {
				continue
			}

			err = ExecScript(&scriptBase)
			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			interfaceExecCondition := domain.InterfaceExecCondition{
				Type: condition.Type,
			}
			interfaceExecCondition.Raw, _ = json.Marshal(scriptBase)
			processor.Result.PostConditions = append(processor.Result.PostConditions, interfaceExecCondition)
		} else if condition.Type == consts.ConditionTypeCheckpoint {

			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)
			if checkpointBase.Disabled {
				continue
			}

			resp := entity.Response
			err = ExecCheckPoint(&checkpointBase, resp, 0)
			checkpointHelper.GenResultMsg(&checkpointBase)

			interfaceExecCondition := domain.InterfaceExecCondition{
				Type: condition.Type,
			}

			interfaceExecCondition.Raw, _ = json.Marshal(checkpointBase)
			processor.Result.PostConditions = append(processor.Result.PostConditions, interfaceExecCondition)

			if _, ok := detail["checkpoint"]; !ok {
				detail["checkpoint"] = []map[string]interface{}{}
			}
			detail["checkpoint"] = append(detail["checkpoint"].([]map[string]interface{}), map[string]interface{}{
				"resultStatus": checkpointBase.ResultStatus, "resultMsg": checkpointBase.ResultMsg,
			})
		} else if condition.Type == consts.ConditionTypeResponseDefine {
			var responseDefineBase domain.ResponseDefineBase
			json.Unmarshal(condition.Raw, &responseDefineBase)
			if responseDefineBase.Disabled {
				continue
			}

			resp := entity.Response

			err = ExecResponseDefine(&responseDefineBase, resp)

			interfaceExecCondition := domain.InterfaceExecCondition{
				Type: condition.Type,
			}

			interfaceExecCondition.Raw, _ = json.Marshal(responseDefineBase)
			processor.Result.PostConditions = append(processor.Result.PostConditions, interfaceExecCondition)

			detail["responseDefine"] = map[string]interface{}{"resultStatus": responseDefineBase.ResultStatus, "resultMsg": responseDefineBase.ResultMsg}
		}
	}

	return
}
