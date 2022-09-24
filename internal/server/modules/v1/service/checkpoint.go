package service

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	execHelper "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"strings"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepo `inject:""`
	InterfaceRepo  *repo.InterfaceRepo  `inject:""`
	ProjectRepo    *repo.ProjectRepo    `inject:""`
	ExecCache      *business.ExecCache  `inject:""`
}

func (s *CheckpointService) List(interfaceId int) (checkpoints []model.InterfaceCheckpoint, err error) {
	checkpoints, err = s.CheckpointRepo.List(uint(interfaceId))

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

func (s *CheckpointService) CheckInterface(interf model.Interface, resp serverDomain.InvocationResponse,
	interfaceExecLog *model.Log) (logCheckpoints []domain.ExecInterfaceCheckpoint, err error) {
	checkpoints, _ := s.CheckpointRepo.List(interf.ID)

	for _, checkpoint := range checkpoints {
		logCheckpoint, err := s.Check(checkpoint, resp, interf.ProjectId, interfaceExecLog)
		if err == nil {
			interfaceCheckpoint := domain.ExecInterfaceCheckpoint{}
			copier.CopyWithOption(&interfaceCheckpoint, logCheckpoint, copier.Option{DeepCopy: true})
			logCheckpoints = append(logCheckpoints, interfaceCheckpoint)
		}
	}

	return
}

func (s *CheckpointService) Check(checkpoint model.InterfaceCheckpoint, resp serverDomain.InvocationResponse,
	projectId uint, interfaceExecLog *model.Log) (logCheckpoint model.LogCheckpoint, err error) {
	if checkpoint.Disabled {
		checkpoint.ResultStatus = ""

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	checkpoint.ResultStatus = consts.Fail

	// Response ResultStatus
	if checkpoint.Type == consts.ResponseStatus {
		expectCode := stringUtils.ParseInt(checkpoint.Value)

		checkpoint.ActualResult = fmt.Sprintf("%d", resp.StatusCode.Int())

		if checkpoint.Operator == consts.Equal && resp.StatusCode.Int() == expectCode {
			checkpoint.ResultStatus = consts.Pass
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint)
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
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint)
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
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	// Extractor
	if checkpoint.Type == consts.Extractor {
		extractorValue := s.ExecCache.Get(checkpoint.ExtractorVariable)
		checkpoint.ActualResult = extractorValue

		checkpoint.ResultStatus = execHelper.Compare(checkpoint.Operator, extractorValue, checkpoint.Value)

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	return
}
