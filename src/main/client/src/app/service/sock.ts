import * as _ from 'lodash';

import {Injectable} from "@angular/core";
import {CONSTANT} from '../utils/constant';

declare var SockJS;

@Injectable()
export class SockService {
  private uri: string = 'ws/sockjs';

  private url: string;
  private socket: any;
  private status: string;

  private resolveConPromise: (...args: any[]) => void;
  private timer: any;

  constructor() {
    this.url = CONSTANT.SERVICE_URL + this.uri;
  }

  public open(): Promise<{}> {
    this.socket = new SockJS(this.url);
    this.socket.onopen = this.onopen;
    this.socket.onmessage = this.onmessage;
    this.socket.onclose = this.onclose;
    this.socket.onerror = this.onerror;

    return new Promise((resolve, reject) => this.resolveConPromise = resolve);
  }
  public close() {
    if (this.socket) {
      this.socket.close();
    }
  };

  public send(json: any) {
    this.socket.send(JSON.stringify(json));
  }

  onopen = () => {
    this.resolveConPromise();

    this.status = 'CONNECTED';
    console.log('Connected to ' + this.url);
  }

  onmessage = (e) => {
    console.log('message', e.data);
  };
  onclose = () => {
    this.status = 'CLOSED';
    console.log('close');
  };
  onerror = () => {
    console.log('reconnecting...');
    this.timer = setTimeout(() => {
      this.open();
    }, 3000);

    this.status = 'ERROR';
    console.log('error');
  };
}

