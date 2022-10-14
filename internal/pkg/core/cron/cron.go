package cron

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/lib/cron"
	"github.com/aaronchen2k/deeptest/pkg/lib/date"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"sync"
	"time"
)

type ServerCron struct {
	syncMap sync.Map
}

func NewServerCron() *ServerCron {
	inst := &ServerCron{}
	return inst
}

func (s *ServerCron) Init() {
	s.syncMap.Store("isRunning", false)
	s.syncMap.Store("lastCompletedTime", int64(0))

	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.WebCheckInterval),
		func() {
			isRunning, _ := s.syncMap.Load("isRunning")
			lastCompletedTime, _ := s.syncMap.Load("lastCompletedTime")

			if isRunning.(bool) || time.Now().Unix()-lastCompletedTime.(int64) < consts.WebCheckInterval {
				_logUtils.Infof("skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
				return
			}

			s.syncMap.Store("isRunning", true)

			// do somethings

			s.syncMap.Store("isRunning", false)
			s.syncMap.Store("lastCompletedTime", time.Now().Unix())
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
