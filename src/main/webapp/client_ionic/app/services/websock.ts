import {Injectable} from "@angular/core";
import "rxjs/add/operator/map";
import {CONSTANT} from "../utils/constant";

declare var ReconnectingWebSocket: any;
declare var SockJS: any;

@Injectable()
export class WebsockService {
  static ENDPOINT:string = 'ws/sockjs';
  static recInterval:any;
  static recTime:number = 0;
  static conn:any;
  
  constructor() {
  }

  static connect() {
    let me = this;
    
//    let url = CONSTANT.SERVICE_URL.replace('http', 'ws') + WebsockService.ENDPOINT + "?t=" + new Date().getTime();
    let url = CONSTANT.SERVICE_URL + WebsockService.ENDPOINT + "?t=" + new Date().getTime();
    console.log(url);

//    WebsockService.conn = new ReconnectingWebSocket(url, null, {maxReconnectAttempts: 3, reconnectDecay: 2});
    WebsockService.conn = new SockJS(url);
    console.log(WebsockService.conn);
    

    WebsockService.conn.onopen = function () {
      console.log('webscoket open');

      let msg = {'act': 'test', 'trans': 'trans'};
      WebsockService.conn.send(JSON.stringify(msg));
      // WebsockService.conn.close();
    };
    WebsockService.conn.onclose = function () {
      console.log('webscoket close');
    }

    WebsockService.conn.onerror = function () {
      console.log('webscoket error');
    };
    WebsockService.conn.onmessage = function(e) {
      console.log('message', e.data);
    };
  };
}
