package third_party

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type LecangCronService struct {
	CronConfigLecangRepo *repo.CronConfigLecangRepo `inject:""`
}

func (s *LecangCronService) Run(options map[string]interface{}) (f func() error) {
	f = func() error {
		return nil
	}
	return
}

func (s *LecangCronService) CallBack(options map[string]interface{}, err error) func() {
	return func() {

	}
}

func (s *LecangCronService) Get(id uint) (ret model.CronConfigLecang, err error) {
	return s.CronConfigLecangRepo.GetById(id)
}

func (s *LecangCronService) Create(req v1.LecangCronReq) (id uint, err error) {
	cronConfigLecangReq := model.CronConfigLecang{}
	copier.CopyWithOption(&cronConfigLecangReq, &req, copier.Option{DeepCopy: true})

	id, err = s.CronConfigLecangRepo.Create(cronConfigLecangReq)

	return
}

func (s *LecangCronService) Update(id uint, req v1.LecangCronReq) (err error) {
	cronConfigLecangReq := model.CronConfigLecang{}
	copier.CopyWithOption(&cronConfigLecangReq, &req, copier.Option{DeepCopy: true})
	cronConfigLecangReq.ID = id

	err = s.CronConfigLecangRepo.Update(cronConfigLecangReq)

	return
}

func (s *LecangCronService) Save(id uint, req v1.LecangCronReq) (ret uint, err error) {
	cronConfigLecangReq := model.CronConfigLecang{}
	copier.CopyWithOption(&cronConfigLecangReq, &req, copier.Option{DeepCopy: true})
	cronConfigLecangReq.ID = id

	ret, err = s.CronConfigLecangRepo.Save(cronConfigLecangReq)

	return
}
