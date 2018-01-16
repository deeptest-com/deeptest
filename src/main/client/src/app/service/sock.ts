import * as _ from 'lodash';

import {Injectable} from "@angular/core";
import {CONSTANT} from '../utils/constant';
import {WS_CONSTANT} from '../utils/ws-constant';
import {GlobalState} from '../global.state';
declare var SockJS;

@Injectable()
export class SockService {
  private uri: string = 'ws/sockjs';

  private url: string;
  private sock: any;
  private handlers = {};

  constructor(private _state:GlobalState) {
    this.url = CONSTANT.SERVICE_URL + this.uri;
  }

  private _opened: boolean = false;

  public open(): void {

    this.close();
    console.log('wsConnect open', this._opened, this.url);

    // if (!this._opened) {
      this.sock = new SockJS(this.url);
      this.sock.onopen = (e) => {
        this.callHandlers('open', e);
      }
      this.sock.onmessage = (e) => {
        this.callHandlers('message', JSON.parse(e.data));
      }
      this.sock.onclose = (e) => {
        this.callHandlers('close', e);
      }
      this._opened = true;
    // }
  }

  public close(): void {
    if (this._opened) {
      this.sock.close();
      delete this.sock;
      this._opened = false;
    }
  }

  private callHandlers (type: string, ...params: any[]) {
    if (this.handlers[type]) {
      this.handlers[type].forEach(function(cb) {
        cb.apply(cb, params);
      });
    }
  }

  private addEvent (type: string, callback: Function) : void {
    if (!this.handlers[type]) this.handlers[type] = [];
    this.handlers[type].push(callback);
  }

  public onOpen (callback: (e: any) => any) : void {
    this.addEvent('open', callback);
  }
  public onMessage (callback: (data: any) => any) : void {
    this.addEvent('message', callback);
  }
  public onClose (callback: (e: any) => any) : void {
    this.addEvent('close', callback);
  }

  public send (data: any) {
    if (this._opened) {
      var msg = JSON.stringify(data);
      this.sock.send(msg);
    }
  }

  wsConnect() {
    this.handlers = {};

    this.onMessage((json) => {

      if (json.code != 1) {
        console.log('ws error: ', json.code);
        return;
      }

      this._state.notifyDataChanged(json.type, json);
    });
    this.onOpen((e) => {
      console.log('wsConnect onOpen');

      this.send({
        type: WS_CONSTANT.WS_OPEN
      });
    });

    this.open();
  }

}

