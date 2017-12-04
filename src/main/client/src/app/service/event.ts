import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class EventService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'event/';

    list(pageSize: number, page: number, status: string) {
        return this._reqService.post(this._api_url + 'list', {pageSize: pageSize, page: page, status: status});
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: number) {
        return this._reqService.post(this._api_url + 'save', model);
    }
}

