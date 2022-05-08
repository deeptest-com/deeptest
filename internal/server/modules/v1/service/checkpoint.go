package service

import (
	"encoding/json"
	_cacheUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/cache"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	stringUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/string"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
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

	checkpoint.Result = serverConsts.Fail

	// Response Status
	if checkpoint.Type == serverConsts.ResponseStatus {
		expectCode := stringUtils.ParseInt(checkpoint.Value)

		if checkpoint.Operator == serverConsts.Equal && resp.StatusCode.Int() == expectCode {
			checkpoint.Result = serverConsts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	// Response Header
	if checkpoint.Type == serverConsts.ResponseHeader {
		headerValue := ""
		for _, h := range resp.Headers {
			if h.Name == checkpoint.Expression {
				headerValue = h.Value
				break
			}
		}

		if checkpoint.Operator == serverConsts.Equal && headerValue == checkpoint.Value {
			checkpoint.Result = serverConsts.Pass
		} else if checkpoint.Operator == serverConsts.NotEqual && headerValue != checkpoint.Value {
			checkpoint.Result = serverConsts.Pass
		} else if checkpoint.Operator == serverConsts.Contain && strings.Contains(headerValue, checkpoint.Value) {
			checkpoint.Result = serverConsts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	// Response Body
	if checkpoint.Type == serverConsts.ResponseBody {
		if checkpoint.Operator == serverConsts.Equal && resp.Content == checkpoint.Value {
			checkpoint.Result = serverConsts.Pass
		} else if checkpoint.Operator == serverConsts.NotEqual && resp.Content != checkpoint.Value {
			checkpoint.Result = serverConsts.Pass
		} else if checkpoint.Operator == serverConsts.Contain && strings.Contains(resp.Content, checkpoint.Value) {
			checkpoint.Result = serverConsts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	// Extractor
	if checkpoint.Type == serverConsts.Extractor {
		extractorValue := _cacheUtils.GetCache(strconv.Itoa(projectId), checkpoint.ExtractorVariable)
		logUtils.Infof("%s = %v", checkpoint.ExtractorVariable, extractorValue)

		if checkpoint.Operator == serverConsts.Equal {
			if extractorValue == checkpoint.Value {
				checkpoint.Result = serverConsts.Pass
			}
		} else if checkpoint.Operator == serverConsts.NotEqual {
			if extractorValue != checkpoint.Value {
				checkpoint.Result = serverConsts.Pass
			}
		} else if checkpoint.Operator == serverConsts.Contain {
			if strings.Contains(extractorValue, checkpoint.Value) {
				checkpoint.Result = serverConsts.Pass
			}
		} else {
			checkpoint.Result = s.Compare(checkpoint.Operator, extractorValue, checkpoint.Value)
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	return
}

func (s *CheckpointService) Compare(operator serverConsts.CheckpointOperator, actual, expect string) (
	result serverConsts.CheckpointResult) {

	result = serverConsts.Fail

	actualFloat, err1 := strconv.ParseFloat(actual, 64)
	expectFloat, err2 := strconv.ParseFloat(expect, 64)

	if err1 != nil || err2 != nil { // not a number
		return
	}

	switch operator.String() {
	case serverConsts.GreaterThan.String():
		result = GetResult(actualFloat > expectFloat)

	case serverConsts.LessThan.String():
		result = GetResult(actualFloat < expectFloat)

	case serverConsts.GreaterThanOrEqual.String():
		result = GetResult(actualFloat >= expectFloat)

	case serverConsts.LessThanOrEqual.String():
		result = GetResult(actualFloat <= expectFloat)

	default:

	}

	return
}

func GetResult(b bool) (
	result serverConsts.CheckpointResult) {

	if b {
		result = serverConsts.Pass
	} else {
		result = serverConsts.Fail
	}

	return

}
