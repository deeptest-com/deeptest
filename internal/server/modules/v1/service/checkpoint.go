package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	_cacheUtils "github.com/aaronchen2k/deeptest/pkg/lib/cache"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
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

func (s *CheckpointService) CheckByInterface(interfaceId uint, resp serverDomain.InvocationResponse, projectId int) (err error) {
	checkpoints, _ := s.CheckpointRepo.List(interfaceId)

	for _, checkpoint := range checkpoints {
		s.Check(checkpoint, resp, projectId)
	}

	return
}

func (s *CheckpointService) Check(checkpoint model.InterfaceCheckpoint, resp serverDomain.InvocationResponse, projectId int) (err error) {
	if checkpoint.Disabled {
		checkpoint.Result = ""
		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	checkpoint.Result = consts.Fail

	// Response Status
	if checkpoint.Type == consts.ResponseStatus {
		expectCode := stringUtils.ParseInt(checkpoint.Value)

		if checkpoint.Operator == consts.Equal && resp.StatusCode.Int() == expectCode {
			checkpoint.Result = consts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
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

		if checkpoint.Operator == consts.Equal && headerValue == checkpoint.Value {
			checkpoint.Result = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && headerValue != checkpoint.Value {
			checkpoint.Result = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(headerValue, checkpoint.Value) {
			checkpoint.Result = consts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	// Response Body
	if checkpoint.Type == consts.ResponseBody {
		if checkpoint.Operator == consts.Equal && resp.Content == checkpoint.Value {
			checkpoint.Result = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && resp.Content != checkpoint.Value {
			checkpoint.Result = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(resp.Content, checkpoint.Value) {
			checkpoint.Result = consts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	// Extractor
	if checkpoint.Type == consts.Extractor {
		extractorValue := _cacheUtils.GetCache(strconv.Itoa(projectId), checkpoint.ExtractorVariable)
		logUtils.Infof("%s = %v", checkpoint.ExtractorVariable, extractorValue)

		if checkpoint.Operator == consts.Equal {
			if extractorValue == checkpoint.Value {
				checkpoint.Result = consts.Pass
			}
		} else if checkpoint.Operator == consts.NotEqual {
			if extractorValue != checkpoint.Value {
				checkpoint.Result = consts.Pass
			}
		} else if checkpoint.Operator == consts.Contain {
			if strings.Contains(extractorValue, checkpoint.Value) {
				checkpoint.Result = consts.Pass
			}
		} else {
			checkpoint.Result = s.Compare(checkpoint.Operator, extractorValue, checkpoint.Value)
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	return
}

func (s *CheckpointService) Compare(operator consts.CheckpointOperator, actual, expect string) (
	result consts.CheckpointResult) {

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
	result consts.CheckpointResult) {

	if b {
		result = consts.Pass
	} else {
		result = consts.Fail
	}

	return

}
