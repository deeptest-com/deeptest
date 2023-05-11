package service

import (
	"encoding/json"
	"fmt"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	agentUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
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

func (s *CheckpointService) List(interfaceId uint) (checkpoints []model.DebugInterfaceCheckpoint, err error) {
	checkpoints, err = s.CheckpointRepo.List(interfaceId)

	return
}

func (s *CheckpointService) Get(id uint) (checkpoint model.DebugInterfaceCheckpoint, err error) {
	checkpoint, err = s.CheckpointRepo.Get(id)

	return
}

func (s *CheckpointService) Create(checkpoint *model.DebugInterfaceCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(checkpoint)

	return
}

func (s *CheckpointService) Update(checkpoint *model.DebugInterfaceCheckpoint) (err error) {
	err = s.CheckpointRepo.Save(checkpoint)

	return
}

func (s *CheckpointService) Delete(reqId uint) (err error) {
	err = s.CheckpointRepo.Delete(reqId)

	return
}

func (s *CheckpointService) CheckInterface(endpointInterfaceId, scenarioProcessorId uint, resp domain.DebugResponse, usedBy consts.UsedBy) (
	logCheckpoints []domain.ExecInterfaceCheckpoint, status consts.ResultStatus, err error) {

	checkpoints, _ := s.CheckpointRepo.List(endpointInterfaceId)

	status = consts.Pass
	for _, checkpoint := range checkpoints {
		logCheckpoint, err := s.Check(checkpoint, scenarioProcessorId, resp, usedBy)

		if err != nil || logCheckpoint.ResultStatus == consts.Fail {
			status = consts.Fail
		}
	}

	return
}

func (s *CheckpointService) Check(checkpoint model.DebugInterfaceCheckpoint, scenarioProcessorId uint, resp domain.DebugResponse,
	usedBy consts.UsedBy) (logCheckpoint model.ExecLogCheckpoint, err error) {

	if checkpoint.Disabled {
		checkpoint.ResultStatus = ""

		s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

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

		s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

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

		s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	checkpoint.ActualResult = ""

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

		s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

		return
	}

	// Judgement
	if checkpoint.Type == consts.Judgement {
		variableMap, datapools, _ := s.VariableService.GetCombinedVarsForCheckpoint(checkpoint.EndpointInterfaceId, scenarioProcessorId)

		result, _ := agentExec.EvaluateGovaluateExpressionWithVariables(checkpoint.Expression, variableMap, datapools)

		checkpoint.ActualResult = fmt.Sprintf("%v", result)

		ret, ok := result.(bool)
		if ok && ret {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

		return
	}

	// Extractor
	if checkpoint.Type == consts.Extractor {
		// get extractor variable value saved by previous extract opt
		extractorPo, _ := s.ExtractorRepo.GetByInterfaceVariable(checkpoint.ExtractorVariable, 0, checkpoint.EndpointInterfaceId)
		checkpoint.ActualResult = extractorPo.Result

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		s.CheckpointRepo.UpdateResult(checkpoint, usedBy)

		return
	}

	return
}
