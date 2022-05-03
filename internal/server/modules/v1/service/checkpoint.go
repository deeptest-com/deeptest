package service

import (
	"encoding/json"
	_cacheUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/cache"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
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

	if checkpoint.Type == serverConsts.ResponseStatus {
		checkpoint.Result = serverConsts.Fail

		if checkpoint.Value == checkpoint.Value {
			checkpoint.Result = serverConsts.Pass
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	if checkpoint.Type == serverConsts.ResponseHeader {
		checkpoint.Result = serverConsts.Fail
		for _, h := range resp.Headers {
			if h.Name == checkpoint.Expression {
				if checkpoint.Operator == serverConsts.Contain && strings.Contains(h.Value, checkpoint.Value) {
					checkpoint.Result = serverConsts.Pass
					break

				} else if checkpoint.Operator == serverConsts.Equal && h.Value == checkpoint.Value {
					checkpoint.Result = serverConsts.Pass
					break

				}
			}
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	if checkpoint.Type == serverConsts.ResponseBody {
		if checkpoint.Operator == serverConsts.Contain {
			if strings.Index(resp.Content, checkpoint.Value) > -1 {
				checkpoint.Result = serverConsts.Pass
			} else {
				checkpoint.Result = serverConsts.Fail
			}
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	if checkpoint.Type == serverConsts.Extractor {
		extractorValue := _cacheUtils.GetCache(strconv.Itoa(projectId), checkpoint.ExtractorVariable)
		logUtils.Infof("%s = %v", checkpoint.ExtractorVariable, extractorValue)

		if checkpoint.Operator == serverConsts.Equal {
			if checkpoint.Value == extractorValue {
				checkpoint.Result = serverConsts.Pass
			} else {
				checkpoint.Result = serverConsts.Fail
			}
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	return
}
