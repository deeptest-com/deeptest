package service

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec"
	execUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	ws "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	WebsocketContextStore sync.Map
)

func RunWebsocket(act consts.ExecType, req *agentExec.WebsocketExecReq, localVarsCache iris.Map, wsMsg *websocket.Message) (err error) {
	logUtils.Infof("run websocket test")

	dpRoom := req.Room
	execData := req.Data

	if req.Data.ExtMode {
		if act == consts.ExecWebsocketConnect {
			err = connectToWebsocketExt(dpRoom, execData, wsMsg)
		} else if act == consts.ExecWebsocketDisconnect {
			err = disconnectToWebsocketExt(dpRoom, execData, wsMsg)
		} else if act == consts.ExecWebsocketSendMsg {
			err = sendMessageExt(dpRoom, execData, wsMsg)
		}
	} else {
		if act == consts.ExecWebsocketConnect {
			err = connectToWebsocketSimple(dpRoom, execData, wsMsg)
		} else if act == consts.ExecWebsocketDisconnect {
			err = disconnectToWebsocketSimple(dpRoom, execData, wsMsg)
		} else if act == consts.ExecWebsocketSendMsg {
			err = sendMessageSimple(dpRoom, execData, wsMsg)
		}
	}

	return
}

// Simple Websocket mode
func connectToWebsocketSimple(dpRoom string, execData domain.WebsocketDebugData, wsMsg *websocket.Message) (err error) {
	address := genUrl(execData.Address, execData.Params)
	header := genHeaders(execData.Headers)

	wsConn, _, err := ws.DefaultDialer.Dial(address, header)
	if err != nil {
		msg := fmt.Sprintf("dial websocket error: %s", err.Error())
		logUtils.Infof(msg)

		execUtils.SendExecMsg(msg, consts.Exception, wsMsg)
		return
	}

	setWsConn(dpRoom, wsConn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	setExecCtx(dpRoom, &ctx)

	// listen websocket messages
	go func() {
		defer cancel()

		for {
			message := ""

			_, msg, err := wsConn.ReadMessage()
			if err != nil {
				//message = "ERR: " + err.Error()
				//execUtils.SendExecMsg(message, consts.ProgressResult, wsMsg)
				goto Label_END_WEBSOCKET_TEST
			}

			message = string(msg)
			execUtils.SendExecMsg(message, consts.ProgressResult, wsMsg)

			select {
			case <-ctx.Done():
				logUtils.Debug("<<<<<<< stop websocket test")
				goto Label_END_WEBSOCKET_TEST

			default:
			}
		}

	Label_END_WEBSOCKET_TEST:
	}()

	return
}
func disconnectToWebsocketSimple(dpRoom string, execData domain.WebsocketDebugData, wsMsg *websocket.Message) (err error) {
	ctx := getExecCtx(dpRoom)
	if ctx != nil {
		(*ctx).Done()
	}

	setExecCtx(dpRoom, nil)

	wsConn := getWsConn(dpRoom)
	if wsConn == nil {
		return
	}

	err = wsConn.Close()
	if err != nil {
		logUtils.Infof("disconnect ws websocket error: %s", err.Error())
		return
	}

	return
}
func sendMessageSimple(dpRoom string, execData domain.WebsocketDebugData, wsMsg *websocket.Message) (err error) {
	wsConn := getWsConn(dpRoom)
	if wsConn == nil {
		return
	}

	err = wsConn.WriteMessage(ws.TextMessage, []byte(execData.Message))
	if err != nil {
		log.Println("write ws websocket msg error: %s", err.Error())
		return
	}

	return
}

// Extend Websocket mode
func connectToWebsocketExt(dpRoom string, execData domain.WebsocketDebugData, wsMsg *websocket.Message) (err error) {
	address := genUrl(execData.Address, execData.Params)
	//header := genHeaders(execData.Headers)

	events := neffos.WithTimeout{
		ReadTimeout:  0,
		WriteTimeout: 0,
		Namespaces: neffos.Namespaces{
			execData.Namespace: neffos.Events{
				neffos.OnAnyEvent: func(c *neffos.NSConn, msg neffos.Message) error {
					m := iris.Map{
						"namespace": msg.Namespace,
						"room":      msg.Room,
						"body":      string(msg.Body),
					}
					execUtils.SendExecMsg(m, consts.ProgressResult, wsMsg)

					return nil
				},
			},
		},
	}

	client, err := neffos.Dial(context.TODO(), gorilla.DefaultDialer, address, events)
	nsConn, err := client.Connect(context.TODO(), execData.Namespace)
	nsRoom, err := nsConn.JoinRoom(context.TODO(), execData.Room)

	setNsRoom(dpRoom, nsRoom)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	setExecCtx(dpRoom, &ctx)

	// just wait, message events caught in events
	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				logUtils.Debug("<<<<<<< stop websocket test")
				goto Label_END_WEBSOCKET_TEST

			default:
			}
		}

	Label_END_WEBSOCKET_TEST:
	}()

	return
}
func disconnectToWebsocketExt(dpRoom string, execData domain.WebsocketDebugData, wsMsg *websocket.Message) (err error) {
	ctx := getExecCtx(dpRoom)
	if ctx != nil {
		(*ctx).Done()
	}

	setExecCtx(dpRoom, nil)

	nsRoom := getNsRoom(dpRoom)
	if nsRoom == nil {
		return
	}

	err = nsRoom.NSConn.LeaveAll(context.Background())
	if err != nil {
		logUtils.Infof("disconnect ns websocket error: %s", err.Error())
		return
	}

	return
}
func sendMessageExt(dpRoom string, execData domain.WebsocketDebugData, wsMsg *websocket.Message) (err error) {
	nsRoom := getNsRoom(dpRoom)
	if nsRoom == nil {
		return
	}

	ok := nsRoom.Emit(execData.Event, []byte(execData.Message))
	if !ok {
		log.Println("write ns websocket msg failed")
		return
	}

	return
}

// helper methods
func setWsConn(dpRoom string, conn *ws.Conn) {
	WebsocketContextStore.Store("ws_conn_"+dpRoom, conn)
}
func getWsConn(dpRoom string) (ret *ws.Conn) {
	obj, ok := WebsocketContextStore.Load("ws_conn_" + dpRoom)

	if ok {
		ret = obj.(*ws.Conn)
	}

	return
}

func setNsRoom(dpRoom string, obj *neffos.Room) {
	WebsocketContextStore.Store("ns_room_"+dpRoom, obj)
}
func getNsRoom(dpRoom string) (ret *neffos.Room) {
	obj, ok := WebsocketContextStore.Load("ns_room_" + dpRoom)

	if ok {
		ret = obj.(*neffos.Room)
	}

	return
}

func setExecCtx(room string, ctx *context.Context) {
	WebsocketContextStore.Store("ctx_"+room, ctx)
}
func getExecCtx(room string) (ret *context.Context) {
	obj, ok := WebsocketContextStore.Load("ctx_" + room)

	if ok {
		ret = obj.(*context.Context)
	}

	return
}

func genUrl(address string, params *[]domain.Param) (ret string) {
	parsedUrl, _ := url.Parse(address)

	baseUrl := &url.URL{
		Scheme: parsedUrl.Scheme,
		Host:   parsedUrl.Host,
		Path:   parsedUrl.Path,
	}

	queryParams := parsedUrl.Query()
	for _, p := range *params {
		if p.Name == "" || p.Value == "" {
			continue
		}

		queryParams.Add(p.Name, p.Value)
	}

	baseUrl.RawQuery = queryParams.Encode()

	ret = baseUrl.String()

	return
}

func genHeaders(headers *[]domain.Header) (ret http.Header) {
	ret = http.Header{}

	for _, h := range *headers {
		if h.Name == "" || h.Value == "" {
			continue
		}

		if value, ok := ret[h.Name]; !ok || value == nil {
			ret[h.Name] = []string{h.Value}
		} else {
			ret[h.Name] = append(ret[h.Name], h.Value)
		}
	}

	return
}
