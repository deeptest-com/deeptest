import * as neffos from 'neffos.js';
import {NSConn} from "neffos.js";

import bus from "@/utils/eventBus";
import {getToken} from "@/utils/localToken";
import settings from "@/config/settings";
import {getAgentUrl} from "@/utils/agentEnv";

export type WsEvent = {
  room: string;
  code: string;
  data: any;
};

export const WsDefaultNameSpace = 'default'
export const WsDefaultRoom = 'default'


export class WebSocket {
  static conn: NSConn

  static async init(reConn): Promise<any> {
    const url = await getWebSocketApi()
    console.log(`init websocket, connect to ` + url)

    if (reConn || !WebSocket.conn) {
      try {
        const conn = await neffos.dial(url, {
          default: {
            _OnNamespaceConnected: (nsConn, msg) => {
              if (nsConn.conn.wasReconnected()) {
                console.log('re-connected after ' + nsConn.conn.reconnectTries.toString() + ' trie(s)')
              }

              console.log('connected to namespace: ' + msg.Namespace)
              WebSocket.conn = nsConn
              bus.emit(settings.eventWebSocketConnStatus, {msg: '{"conn": "success"}'});
            },

            _OnNamespaceDisconnect: (_nsConn, msg) => {
              console.log('disconnected from namespace: ' + msg.Namespace)
            },

            OnVisit: (_nsConn, msg) => {
              console.log('OnVisit', msg)
            },

            // implement in webpage
            OnChat: (_nsConn, json) => {
              console.log('OnChat in util cls', json)
              bus.emit(settings.eventWebSocketMsg, {room: json.Room, msg: json.Body});
            }
          }
        })

        await conn.connect(WsDefaultNameSpace)

      } catch (err) {
        console.log('failed connect to websocket', err)
        bus.emit(settings.eventWebSocketConnStatus, {msg: '{"conn": "fail"}'});
      }
    }
    return WebSocket
  }

  static joinRoomAndSend(roomName: string, msg: string): void {
    if (WebSocket.conn && WebSocket.conn.room(roomName)) {
      WebSocket.conn.room(roomName).emit('OnChat', msg)
      return
    } else {
      WebSocket.init(true).then(() => {
        WebSocket.conn.joinRoom(roomName).then((room) => {
          console.log(`success to join room "${roomName}"`)
          WebSocket.conn.room(roomName).emit('OnChat', msg)

        }).catch(err => {
          console.log(`fail to join room ${roomName}`, err)
          bus.emit(settings.eventWebSocketConnStatus, {msg: '{"conn": "fail"}'});
        })
      })
    }
  }

  static async sentMsg(roomName: string, msg: string): Promise<void> {
    console.log(`send msg to room "${roomName}"`)
    roomName = await getToken() || roomName;
    if (!WebSocket.conn) return

    WebSocket.conn.leaveAll().then(() =>
        this.joinRoomAndSend(roomName, msg)
    )
  }
}

export async function getWebSocketApi() {
  // const isProd = process.env.NODE_ENV === 'production'
  // const loc = window.location
  // console.log(`${isProd}, ${loc.toString()}`)

  // const apiHost = process.env.VUE_APP_API_AGENT ? process.env.VUE_APP_API_AGENT : ''
  const apiHost = await getAgentUrl();
  const url = apiHost.replace('http', 'ws') + '/ws'
  console.log(`websocket url = ${url}`)

  return url
}
