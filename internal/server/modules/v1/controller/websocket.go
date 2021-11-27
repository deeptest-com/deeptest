package controller

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

type WsCtrl struct {
	Namespace         string
	*websocket.NSConn `stateless:"true"`

	WebSocketService *service.WebSocketService `inject:""`
}

func NewWsCtrl() *WsCtrl {
	return &WsCtrl{Namespace: serverConsts.WsDefaultNameSpace}
}

func (c *WsCtrl) OnNamespaceConnected(msg websocket.Message) error {
	c.WebSocketService.SetConn(c.Conn)

	_logUtils.Infof("websocket client connected %s", c.Conn.ID())

	data := map[string]string{"msg": "from server: connected to websocket"}
	c.WebSocketService.Broadcast(msg.Namespace, "", "OnVisit", data)
	return nil
}

// OnNamespaceDisconnect This will call the "OnVisit" event on all clients, except the current one,
// it can't because it's left but for any case use this type of design
func (c *WsCtrl) OnNamespaceDisconnect(msg websocket.Message) error {
	_logUtils.Infof("%s disconnected", c.Conn.ID())

	data := map[string]string{"msg": "from server: disconnected to websocket"}
	c.WebSocketService.Broadcast(msg.Namespace, "", "OnVisit", data)
	return nil
}

// OnChat This will call the "OnVisit" event on all clients, including the current one, with the 'newCount' variable.
func (c *WsCtrl) OnChat(msg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)

	str := ctx.RemoteAddr()
	_logUtils.Info(str + ", " + string(msg.Body))

	data := map[string]string{"data": fmt.Sprintf("from server: response %s", "abc")}
	c.WebSocketService.Broadcast(msg.Namespace, msg.Room, msg.Event, data)

	return
}

func (c *WsCtrl) TestWs(ctx iris.Context) {
	data := map[string]interface{}{"action": "taskUpdate", "taskId": 1, "msg": ""}
	c.WebSocketService.Broadcast(serverConsts.WsDefaultNameSpace, serverConsts.WsDefaultRoom, serverConsts.WsEvent, data)

	ctx.JSON(domain.Response{Code: domain.NoErr.Code, Data: nil, Msg: domain.NoErr.Msg})
}
