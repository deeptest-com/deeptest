import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class SessionService {
  constructor(private _reqService:RequestService) {
  }

  _api_url = 'session/';

  get(id:number) {
    return this._reqService.post(this._api_url + 'get', {eventId: id});
  }

  save(model:any) {
    return this._reqService.post(this._api_url + 'save', model);
  }

  remove(id: number, type: string) {
    return this._reqService.post(this._api_url + 'remove', {id: id, type: type});
  }
}

