package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"strings"
)

type CheckpointService struct {
	CheckpointRepo  *repo.CheckpointRepo  `inject:""`
	InterfaceRepo   *repo.InterfaceRepo   `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ExtractorRepo   *repo.ExtractorRepo   `inject:""`
	VariableService *VariableService      `inject:""`
}

func (s *CheckpointService) List(interfaceId uint, usedBy consts.UsedBy) (checkpoints []model.InterfaceCheckpoint, err error) {
	checkpoints, err = s.CheckpointRepo.List(interfaceId, usedBy)

	return
}

func (s *CheckpointService) Get(id uint) (checkpoint model.InterfaceCheckpoint, err error) {
	checkpoint, err = s.CheckpointRepo.Get(id)

	return
}

func (s *CheckpointService) Create(checkpoint *model.InterfaceCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(checkpoint)

	return
}

func (s *CheckpointService) Update(checkpoint *model.InterfaceCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(checkpoint)

	return
}

func (s *CheckpointService) Delete(reqId uint) (err error) {
	err = s.CheckpointRepo.Delete(reqId)

	return
}

func (s *CheckpointService) CheckInterface(interfaceId uint, resp v1.InvocationResponse,
	interfaceExecLog *model.ExecLogProcessor, usedBy consts.UsedBy) (
	logCheckpoints []domain.ExecInterfaceCheckpoint, status consts.ResultStatus, err error) {

	checkpoints, _ := s.CheckpointRepo.List(interfaceId, usedBy)

	status = consts.Pass
	for _, checkpoint := range checkpoints {
		logCheckpoint, err := s.Check(checkpoint, resp, interfaceExecLog, usedBy)

		if logCheckpoint.ResultStatus == consts.Fail {
			status = consts.Fail
		}

		if err == nil && interfaceExecLog != nil { // gen report for processor
			interfaceCheckpoint := domain.ExecInterfaceCheckpoint{}
			copier.CopyWithOption(&interfaceCheckpoint, logCheckpoint, copier.Option{DeepCopy: true})

			logCheckpoints = append(logCheckpoints, interfaceCheckpoint)
		}
	}

	return
}

func (s *CheckpointService) Check(checkpoint model.InterfaceCheckpoint, resp v1.InvocationResponse,
	interfaceExecLog *model.ExecLogProcessor, usedBy consts.UsedBy) (logCheckpoint model.ExecLogCheckpoint, err error) {
	if checkpoint.Disabled {
		checkpoint.ResultStatus = ""

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)

		}

		return
	}

	checkpoint.ResultStatus = consts.Pass

	// Response ResultStatus
	if checkpoint.Type == consts.ResponseStatus {
		expectCode := stringUtils.ParseInt(checkpoint.Value)

		checkpoint.ActualResult = fmt.Sprintf("%d", resp.StatusCode.Int())

		if checkpoint.Operator == consts.Equal && resp.StatusCode.Int() == expectCode {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint, usedBy)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	// Response Header
	if checkpoint.Type == consts.ResponseHeader {
		headerValue := ""
		for _, h := range resp.Headers {
			if h.Name == checkpoint.Expression {
				headerValue = h.Value
				break
			}
		}

		checkpoint.ActualResult = headerValue

		if checkpoint.Operator == consts.Equal && headerValue == checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && headerValue != checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(headerValue, checkpoint.Value) {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint, usedBy)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	checkpoint.ActualResult = "<RESPONSE_BODY>"

	// Response Body
	if checkpoint.Type == consts.ResponseBody {
		if checkpoint.Operator == consts.Equal && resp.Content == checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && resp.Content != checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(resp.Content, checkpoint.Value) {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint, usedBy)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	// Judgement
	if checkpoint.Type == consts.Judgement {
		var result interface{}
		if interfaceExecLog != nil { // run by processor
			result, _ = agentExec.EvaluateGovaluateExpressionByScope(checkpoint.Expression, interfaceExecLog.ProcessorId)
		} else {
			variables, _ := s.VariableService.GetVariablesByInterface(checkpoint.InterfaceId)
			result, _ = agentExec.EvaluateGovaluateExpressionWithVariables(checkpoint.Expression, variables)
		}

		ret, ok := result.(bool)
		if ok && ret {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}
		checkpoint.ActualResult = fmt.Sprintf("%v", ret)

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint, usedBy)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	// Extractor
	if checkpoint.Type == consts.Extractor {
		// get extractor variable value saved by previous extract opt
		extractorPo, _ := s.ExtractorRepo.GetByInterfaceVariable(checkpoint.ExtractorVariable, 0, checkpoint.InterfaceId)
		checkpoint.ActualResult = extractorPo.Result

		checkpoint.ResultStatus = utils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint, usedBy)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	return
}
