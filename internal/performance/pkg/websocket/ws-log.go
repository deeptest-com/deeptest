package ptwebsocket

import (
	"encoding/json"
	ptconsts "github.com/aaronchen2k/deeptest/internal/performance/pkg/consts"
	ptdomain "github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

var (
	wsConnLog *neffos.Conn
)

func SendExecLogToClient(data interface{}, resultType ptconsts.MsgResultTypeToWsClient, execUUid string, wsMsg *websocket.Message) {
	resp := ptdomain.WsResp{
		Uuid:       execUUid,
		Category:   ptconsts.MsgCategoryLog,
		ResultType: resultType,
		Data:       data,
	}
	if data != nil {
		resp.Data = data
	}
	bytes, _ := json.Marshal(resp)

	if wsMsg != nil {
		mqData := _domain.MqMsg{
			Namespace: wsMsg.Namespace,
			Room:      wsMsg.Room,
			Event:     wsMsg.Event,
			Content:   string(bytes),
		}
		logUtils.Infof(_i118Utils.Sprintf("ws_send_exec_msg", wsMsg.Room, ptconsts.MsgCategoryLog))

		PubLogMsg(mqData)

	} else {
		logUtils.Infof(string(bytes))
	}
}

func BroadcastLog(namespace, room, event string, content string) {
	if wsConnLog == nil {
		return
	}

	wsConnLog.Server().Broadcast(nil, websocket.Message{
		Namespace: namespace,
		Room:      room,
		Event:     event,
		Body:      []byte(content),
	})
}

func SetLogConn(conn *neffos.Conn) {
	wsConnLog = conn
}
