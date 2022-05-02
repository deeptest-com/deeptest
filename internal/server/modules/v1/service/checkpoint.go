package service

import (
	"encoding/json"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
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

func (s *CheckpointService) CheckByInterface(interfaceId uint, resp serverDomain.InvocationResponse) (err error) {
	checkpoints, _ := s.CheckpointRepo.List(interfaceId)

	for _, checkpoint := range checkpoints {
		s.Check(checkpoint, resp)
	}

	return
}

func (s *CheckpointService) Check(checkpoint model.InterfaceCheckpoint, resp serverDomain.InvocationResponse) (err error) {
	if checkpoint.Disabled {
		checkpoint.Result = ""
		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	if checkpoint.Type == serverConsts.ResponseStatus {
		checkpoint.Result = "FAIL"

		if checkpoint.Value == checkpoint.Value {
			checkpoint.Result = "PASS"
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	if checkpoint.Type == serverConsts.ResponseHeader {
		checkpoint.Result = "FAIL"
		for _, h := range resp.Headers {
			if h.Name == checkpoint.Expression {
				if h.Value == checkpoint.Value {
					checkpoint.Result = "PASS"
				}
				break
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
				checkpoint.Result = "PASS"
			} else {
				checkpoint.Result = "FAIL"
			}
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	if checkpoint.Type == serverConsts.Extractor {
		extractorValue, _ := serverConsts.EnvVar.Load(checkpoint.ExtractorVariable)
		logUtils.Infof("%s = %v", checkpoint.ExtractorVariable, extractorValue)

		if checkpoint.Operator == serverConsts.Equal {
			if checkpoint.Value == extractorValue {
				checkpoint.Result = "PASS"
			} else {
				checkpoint.Result = "FAIL"
			}
		}

		s.CheckpointRepo.UpdateResult(checkpoint)
		return
	}

	return
}
