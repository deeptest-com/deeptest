package service

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_cacheUtils "github.com/aaronchen2k/deeptest/pkg/lib/cache"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepo `inject:""`
	InterfaceRepo  *repo.InterfaceRepo  `inject:""`
	ProjectRepo    *repo.ProjectRepo    `inject:""`
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
	interfaceExecLog *model.Log) (logCheckpoints []domain.InterfaceCheckpoint, err error) {
	checkpoints, _ := s.CheckpointRepo.List(interf.ID)

	for _, checkpoint := range checkpoints {
		logCheckpoint, err := s.Check(checkpoint, resp, interf.ProjectId, interfaceExecLog)
		if err == nil {
			interfaceCheckpoint := domain.InterfaceCheckpoint{}
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
		extractorValue := _cacheUtils.GetCache(strconv.Itoa(int(projectId)), checkpoint.ExtractorVariable)
		logUtils.Infof("%s = %v", checkpoint.ExtractorVariable, extractorValue)
		checkpoint.ActualResult = extractorValue

		if checkpoint.Operator == consts.Equal {
			if extractorValue == checkpoint.Value {
				checkpoint.ResultStatus = consts.Pass
			}
		} else if checkpoint.Operator == consts.NotEqual {
			if extractorValue != checkpoint.Value {
				checkpoint.ResultStatus = consts.Pass
			}
		} else if checkpoint.Operator == consts.Contain {
			if strings.Contains(extractorValue, checkpoint.Value) {
				checkpoint.ResultStatus = consts.Pass
			}
		} else {
			checkpoint.ResultStatus = s.Compare(checkpoint.Operator, extractorValue, checkpoint.Value)
		}

		if interfaceExecLog == nil { // run by interface
			s.CheckpointRepo.UpdateResult(checkpoint)
		} else { // run by processor
			logCheckpoint, err = s.CheckpointRepo.UpdateResultToExecLog(checkpoint, interfaceExecLog)
		}

		return
	}

	return
}

func (s *CheckpointService) Compare(operator consts.ComparisonOperator, actual, expect string) (
	result consts.ResultStatus) {

	result = consts.Fail

	actualFloat, err1 := strconv.ParseFloat(actual, 64)
	expectFloat, err2 := strconv.ParseFloat(expect, 64)

	if err1 != nil || err2 != nil { // not a number
		return
	}

	switch operator.String() {
	case consts.GreaterThan.String():
		result = GetResult(actualFloat > expectFloat)

	case consts.LessThan.String():
		result = GetResult(actualFloat < expectFloat)

	case consts.GreaterThanOrEqual.String():
		result = GetResult(actualFloat >= expectFloat)

	case consts.LessThanOrEqual.String():
		result = GetResult(actualFloat <= expectFloat)

	default:

	}

	return
}

func GetResult(b bool) (
	result consts.ResultStatus) {

	if b {
		result = consts.Pass
	} else {
		result = consts.Fail
	}

	return

}
