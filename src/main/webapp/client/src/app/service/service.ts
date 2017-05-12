import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class ServiceService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'service/';

    list(eventId: number) {
        return this._reqService.post(this._api_url + 'list', {eventId: eventId});
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: number) {
        return this._reqService.post(this._api_url + 'save', model);
    }

    disable(id: number) {
      return this._reqService.post(this._api_url + 'disable', {id: id});
    }
}

