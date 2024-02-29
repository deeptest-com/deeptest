package conductorExec

import (
	"context"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/websocket"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/nxadm/tail"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	logInst *PerformanceLogService
)

func CreatePerformanceLogService() *PerformanceLogService {
	logInst = &PerformanceLogService{
		uuid: _stringUtils.Uuid(),
	}

	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: logInst},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	return logInst
}

type PerformanceLogService struct {
	uuid string

	LogReq *ptdomain.PerformanceLogReq

	logCtx    context.Context
	logCancel context.CancelFunc
}

func (s *PerformanceLogService) StartSendLog(req ptdomain.PerformanceLogReq, wsMsg *websocket.Message) (err error) {
	if s.logCtx != nil {
		return
	}

	s.logCtx, s.logCancel = context.WithCancel(context.Background())

	room := req.Room
	logPath := ptlog.GetLogPath(room)
	var t *tail.Tail

	go func() {
		t, err = tail.TailFile(logPath, tail.Config{
			Follow: true,
			ReOpen: true,
		})
		if err != nil {
			s.logCancel()
			s.logCtx = nil
			return
		}
		//defer t.Cleanup()

		var buffer []string
		timeBefore := time.Now().UnixMilli()

		for line := range t.Lines {
			buffer = append(buffer, line.Text)

			if len(buffer) > 20 || time.Now().UnixMilli()-timeBefore > 1000 {
				data := iris.Map{
					"log": line.Text,
				}
				ptwebsocket.SendExecLogToClient(data, ptconsts.MsgResultRecord, req.Room, wsMsg)

				buffer = make([]string, 0)
				timeBefore = time.Now().UnixMilli()
			}
		}

		s.logCancel()
		s.logCtx = nil
	}()

	go func() {
		for true {
			if s.logCtx == nil || IsWsMsgSuspend() {
				if t != nil {
					t.Stop()
				}
				break
			}

			select {
			case <-s.logCtx.Done():
				_logUtils.Debug("<<<<<<< stop sendLog job by logCtx.Done")

				if t != nil {
					t.Stop()
				}
				break

			default:
			}

			time.Sleep(3 * time.Second)
		}

		s.logCancel()
		s.logCtx = nil
	}()

	return
}

func (s *PerformanceLogService) StopSendLog() (err error) {
	if s.logCancel != nil {
		s.logCancel()
	}
	s.logCtx = nil

	return
}
