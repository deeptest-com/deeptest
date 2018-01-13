import { Injectable } from '@angular/core';
import { Subject }    from 'rxjs/Subject';

@Injectable()
export class GlobalState {

  private _data = new Subject<Object>();
  private _dataStream$ = this._data.asObservable();

  private _subscriptions: Map<string, Map<string, Function>> = new Map<string, Map<string, Function>>();

  constructor() {
    this._dataStream$.subscribe((data) => this._onEvent(data));
  }

  notifyDataChanged(event, value) {
    let current = this._data[event];
    if (current !== value) {

      this._data[event] = value;

      this._data.next({
        event: event,
        data: this._data[event]
      });

      // console.log('===notifyDataChanged', this._data);
    }
  }

  subscribe(event: string, code: string, callback: Function) {
    let subscribers: Map<string, Function> = this._subscriptions.get(event)
      || new Map<string, Function>();

    subscribers.set(code, callback);
    this._subscriptions.set(event, subscribers);

    // console.log('===subscribe', this._subscriptions);
  }
  unsubscribe(event: string, code: string) {
    let subscribers: Map<string, Function> = this._subscriptions.get(event)
      || new Map<string, Function>();

    subscribers.delete(code);
  }

  _onEvent(data: any) {
    // console.log('***_onEvent', data['event'], data['data']);

    let subscribers: Map<string, Function> = this._subscriptions.get(data['event'])
      || new Map<string, Function>();
    // console.log('===_onEvent', subscribers.keys());

    subscribers.forEach((value: Function, key: string) => {
      // console.log('---_onEvent', key);

      value.call(null, data['data']);
    });
  }
}
